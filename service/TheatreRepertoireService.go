package service

import (
	"time"

	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
	"gorm.io/gorm"
)

type TheatreRepertoireService struct {
	MovieRepo       *repo.MovieRepository
	ProjectionRepo  *repo.ProjectionRepository
	TheatreInfoRepo *repo.TheatreInfoRepository
	ContractRepo    *repo.DistributionContractRepository
	HallRepo        *repo.HallRepository
}

func NewTheatreRepertoireService(moviRepo *repo.MovieRepository, projectionRepo *repo.ProjectionRepository, theatreInfoRepo *repo.TheatreInfoRepository,
	contractRepo *repo.DistributionContractRepository, hallRepo *repo.HallRepository) *TheatreRepertoireService {
	return &TheatreRepertoireService{
		MovieRepo:       moviRepo,
		ProjectionRepo:  projectionRepo,
		TheatreInfoRepo: theatreInfoRepo,
		ContractRepo:    contractRepo,
		HallRepo:        hallRepo,
	}
}

func (rs *TheatreRepertoireService) AddProjection(movieId, hallId, days int, ticketPrice float64, timeslot model.Timeslot) ([]model.Projection, error) {
	if !timeslot.IsValid() {
		return nil, &CustomError{Code: ErrBadRequest, Message: "Invalid timeslot"}
	}
	_, err := rs.HallRepo.GetByID(uint(hallId))
	if err != nil {
		return nil, &CustomError{Code: ErrNotFound, Message: "Hall not found"}
	}
	projections := []model.Projection{}

	for i := 0; i < days; i++ {
		projectionTimeslot := timeslot.AddDays(i)
		projection := model.Projection{
			MovieID:   uint(movieId),
			HallId:    uint(hallId),
			Timeslot:  projectionTimeslot,
			Price:     ticketPrice,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		inContract, err := rs.ContractRepo.IsDateInContract(uint(movieId), projectionTimeslot.StartTime)
		if err != nil {
			return nil, err
		}

		if !inContract {
			return nil, &CustomError{Code: ErrBadRequest, Message: "No existing distribution contract for selected projection date"}
		}

		existingProjections, err := rs.ProjectionRepo.GetFilteredProjections(uint(hallId), 0, projectionTimeslot.StartTime)
		if err != nil {
			return nil, err
		}

		if len(rs.removeOverlappingTimeslots([]model.Timeslot{projection.Timeslot}, existingProjections)) < 1 {
			return nil, &CustomError{Code: ErrBadRequest, Message: "Projection overlaps with existing projections"}
		}
		projections = append(projections, projection)
	}

	createdProjections := []model.Projection{}
	for _, projection := range projections {
		newProjection, err := rs.ProjectionRepo.CreateProjection(projection)
		if err != nil {
			return nil, err
		}
		createdProjections = append(createdProjections, *newProjection)
	}

	return createdProjections, nil
}

func (rs *TheatreRepertoireService) GetProjections(hallId, movieId int, date time.Time) ([]model.Projection, error) {
	projections, err := rs.ProjectionRepo.GetFilteredProjections(uint(hallId), uint(movieId), date)
	if err != nil {
		return nil, err
	}

	return projections, nil
}

func (rs *TheatreRepertoireService) GetAvailableTimeslots(movieId, hallId, days int, startDate time.Time) ([]model.Timeslot, error) {
	availableTimeslots, err := rs.getInitialTimeslots(movieId, hallId, startDate)
	if err != nil {
		return nil, err
	}

	for i := 1; i < days; i++ {
		projections, err := rs.ProjectionRepo.GetFilteredProjections(uint(hallId), 0, startDate.AddDate(0, 0, i))
		if err != nil {
			return nil, err
		}
		availableTimeslots = rs.removeOverlappingTimeslots(availableTimeslots, projections)
	}

	for i := range availableTimeslots {
		availableTimeslots[i].ToCET()
	}
	return availableTimeslots, nil
}

func (rs *TheatreRepertoireService) getInitialTimeslots(movieId, hallId int, startDate time.Time) ([]model.Timeslot, error) {
	availableTimeslots := []model.Timeslot{}

	movie, err := rs.MovieRepo.GetByID(uint(movieId))
	if err != nil {
		return nil, &CustomError{Code: ErrNotFound, Message: "Movie not found"}
	}

	projections, err := rs.ProjectionRepo.GetFilteredProjections(uint(hallId), 0, startDate)
	if err != nil {
		return nil, err
	}

	theatreInfo, err := rs.TheatreInfoRepo.GetTheatreInfo()
	if err != nil {
		return nil, err
	}

	startHour := theatreInfo.OpeningHour
	endHour := theatreInfo.ClosingHour
	loc, err := time.LoadLocation("CET")
	if err != nil {
		return nil, err
	}

	startDateTime := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), startHour, 0, 0, 0, loc)
	endDateTime := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), endHour, 0, 0, 0, loc)

	currentTimeSlot := model.NewTimeslot(startDateTime, movie.Duration)
	lastTimeSlot := model.NewTimeslot(endDateTime, movie.Duration)

	for !currentTimeSlot.StartsAfter(lastTimeSlot) && currentTimeSlot.IsValid() {
		isAvailable := true
		for _, projection := range projections {
			if currentTimeSlot.Overlaps(projection.Timeslot, 15) {
				isAvailable = false
				currentTimeSlot = model.NewTimeslot(projection.Timeslot.EndTime, movie.Duration)
				break
			}
		}

		if isAvailable {
			availableTimeslots = append(availableTimeslots, currentTimeSlot)
		}

		currentTimeSlot = currentTimeSlot.AddMinutes(15)
	}
	return availableTimeslots, nil
}

func (rs *TheatreRepertoireService) removeOverlappingTimeslots(timeslots []model.Timeslot, projections []model.Projection) []model.Timeslot {
	for j := 0; j < len(timeslots); {
		timeslot := timeslots[j]
		overlap := false
		for _, projection := range projections {
			if timeslot.Overlaps(projection.Timeslot, 15) {
				overlap = true
				break
			}
		}
		if overlap {
			// Remove the overlapping timeslot from the slice
			timeslots = append(timeslots[:j], timeslots[j+1:]...)
		} else {
			j++
		}
	}
	return timeslots
}

func (trs *TheatreRepertoireService) CountConsecutiveDaysWithProjections(projectionId uint) (int, error) {
	count := 0
	projection, err := trs.ProjectionRepo.GetProjectionByID(projectionId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, &CustomError{Code: ErrBadRequest, Message: "Bad projection ID"}
		}
		return 0, err
	}

	startTime := projection.Timeslot.StartTime
	endTime := projection.Timeslot.EndTime

	for {
		startTime = startTime.AddDate(0, 0, 1) // add one day
		endTime = endTime.AddDate(0, 0, 1)     // add one day
		_, err := trs.ProjectionRepo.GetProjectionByDateAndIds(startTime, endTime, projection.MovieID, projection.HallId)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				break
			} else {
				return 0, err
			}
		}

		count++
	}

	return count, nil
}

func (trs *TheatreRepertoireService) CancelProjection(projectionId uint, cancelConsecutive bool) ([]model.Projection, error) {
	canceledProjections := []model.Projection{}
	projection, err := trs.ProjectionRepo.GetProjectionByID(projectionId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &CustomError{Code: ErrBadRequest, Message: "Bad projection ID"}
		}
		return nil, err
	}

	projection.IsCanceled = true
	projection.UpdatedAt = time.Now()

	tx := trs.ProjectionRepo.DatabaseConnection.Begin()

	projection, err = trs.ProjectionRepo.UpdateProjection(projection, tx)
	if err != nil {
		return nil, err
	}
	canceledProjections = append(canceledProjections, *projection)

	if !cancelConsecutive {
		if err := tx.Commit().Error; err != nil {
			return nil, err
		}
		return canceledProjections, nil
	}

	canceledProjections, err = trs.cancelConsecutiveProjections(projection, tx, canceledProjections)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return canceledProjections, nil
}

func (trs *TheatreRepertoireService) cancelConsecutiveProjections(projection *model.Projection, tx *gorm.DB, canceledProjections []model.Projection) ([]model.Projection, error) {
	startTime := projection.Timeslot.StartTime
	endTime := projection.Timeslot.EndTime

	for {
		startTime = startTime.AddDate(0, 0, 1)
		endTime = endTime.AddDate(0, 0, 1)
		nextProjection, err := trs.ProjectionRepo.GetProjectionByDateAndIds(startTime, endTime, projection.MovieID, projection.HallId)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				break
			} else {
				tx.Rollback()
				return nil, err
			}
		}
		nextProjection.IsCanceled = true
		nextProjection.UpdatedAt = time.Now()
		nextProjection, err = trs.ProjectionRepo.UpdateProjection(nextProjection, tx)
		if err != nil {
			return nil, err
		}
		canceledProjections = append(canceledProjections, *nextProjection)
	}
	return canceledProjections, nil
}
