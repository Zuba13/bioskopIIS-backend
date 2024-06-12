package main

import (
	"log"
	"net/http"

	"bioskop.com/projekat/bioskopIIS-backend/handler"
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "host=localhost user=postgres password=super dbname=cinema port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}

	// Auto-migrate models for the User and Movie modules
	//database.Exec("ALTER TABLE projections ADD COLUMN IF NOT EXISTS is_canceled boolean;")
	//database.Exec("UPDATE projections SET is_canceled = false WHERE is_canceled IS NULL;")
	//database.Exec("ALTER TABLE projections ALTER COLUMN is_canceled SET NOT NULL;")
	database.AutoMigrate(&model.TheatreInfo{})
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Movie{})
	database.AutoMigrate(&model.Projection{})
	database.AutoMigrate(&model.Hall{})
	database.AutoMigrate(&model.Ticket{})

	database.AutoMigrate(&model.Review{})
	database.AutoMigrate(&model.Actor{})
	database.AutoMigrate(&model.Director{})
	database.Exec("DO $$ BEGIN CREATE TYPE contractmodel AS ENUM ('Bidding', 'Percentage'); EXCEPTION WHEN duplicate_object THEN null; END $$;")
	database.AutoMigrate(&model.DistributionCompany{}, &model.DistributionContract{})
	database.AutoMigrate(&model.Contract{})
	database.AutoMigrate(&model.ContractItem{})
	database.AutoMigrate(&model.Supplier{})
	database.AutoMigrate(&model.ProjectionCanceledNotification{})

	return database
}

func startServer(userHandler *handler.UserHandler, movieHandler *handler.MovieHandler) {
	router := mux.NewRouter().StrictSlash(true)

	userHandler.RegisterUserHandler(router)
	movieHandler.RegisterMovieHandler(router)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	log.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8085", router))
}

func main() {
	database := initDB()
	if database == nil {
		log.Println("Failed to connect to database!")
		return
	}

	// Initialize User repository, service, and handler
	userRepo := &repo.UserRepository{DatabaseConnection: database}
	userService := &service.UserService{UserRepo: *userRepo}
	userHandler := &handler.UserHandler{UserService: *userService}

	// Initialize Movie repository, service, and handler
	movieRepo := &repo.MovieRepository{DatabaseConnection: database}
	movieService := &service.MovieService{MovieRepository: *movieRepo}
	movieHandler := &handler.MovieHandler{MovieService: *movieService}

	projectionRepo := &repo.ProjectionRepository{DatabaseConnection: database}
	projectionService := &service.ProjectionService{ProjectionRepo: *projectionRepo}
	projectionHandler := &handler.ProjectionHandler{ProjectionService: *projectionService}

	hallRepo := &repo.HallRepository{DatabaseConnection: database}
	hallService := &service.HallService{HallRepository: *hallRepo}
	hallHandler := &handler.HallHandler{HallService: *hallService}

	ticketRepo := &repo.TicketRepository{DB: database}
	ticketService := service.NewTicketService(ticketRepo, userRepo)
	ticketHandler := &handler.TicketHandler{TicketService: *ticketService, UserService: *userService, ProjectionService: *projectionService}


	reviewRepo := &repo.ReviewRepository{DatabaseConnection: database}
	reviewService := &service.ReviewService{ReviewRepo: *reviewRepo}
	reviewHandler := &handler.ReviewHandler{ReviewService: *reviewService}

	actorRepo := &repo.ActorRepository{DatabaseConnection: database}
	directorRepo := &repo.DirectorRepository{DatabaseConnection: database}
	distContractRepo := &repo.DistributionContractRepository{DatabaseConnection: database}
	actorService := &service.ActorService{ActorRepository: *actorRepo}
	directorService := &service.DirectorService{DirectorRepository: *directorRepo}
	catalogService := &service.CatalogService{MovieRepository: *movieRepo, ContractReposiotry: *distContractRepo}
	catalogHandler := &handler.CatalogHandler{CatalogService: *catalogService, ActorService: *actorService, DirectorService: *directorService}

	distCompanyRepo := &repo.DistributionCompanyRepository{DatabaseConnection: database}
	distContractService := &service.DistributionContractService{DistributionContractRepository: distContractRepo, DistributionCompanyRepository: distCompanyRepo}
	distContractHandler := &handler.DistributionContractHandler{DistributionContractService: distContractService}

	contractRepo := &repo.ContractRepository{DatabaseConnection: database}
	contractService := &service.ContractService{ContractRepository: *contractRepo}
	contractHandler := &handler.ContractHandler{ContractService: *contractService}

	contractItemRepo := &repo.ContractItemRepository{DatabaseConnection: database}
	contractItemService := &service.ContractItemService{ContractItemRepository: *contractItemRepo}
	contractItemHandler := &handler.ContractItemHandler{ContractItemService: *contractItemService}

	supplierRepo := &repo.SupplierRepository{DatabaseConnection: database}
	supplierService := &service.SupplierService{SupplierRepository: *supplierRepo}
	supplierHandler := &handler.SupplierHandler{SupplierService: *supplierService}

	notificationRepo := &repo.NotificationRepository{DatabaseConnection: database}
	notificationService := service.NewNotificationService(notificationRepo, ticketRepo)
	notificationHandler := &handler.NotificationHandler{NotificationService: *notificationService}

	theatreInfoRepo := &repo.TheatreInfoRepository{DatabaseConnection: database}
	repertoireHandler := &handler.TheatreRepertoireHandler{TheatreRepertoireService: *service.NewTheatreRepertoireService(movieRepo, projectionRepo, theatreInfoRepo,
		distContractRepo, hallRepo, notificationService, ticketService)}

	// Create a new router
	router := mux.NewRouter().StrictSlash(true)

	// Register your handlers with the router
	userHandler.RegisterUserHandler(router)
	movieHandler.RegisterMovieHandler(router)
	projectionHandler.RegisterProjectionHandler(router)
	hallHandler.RegisterHallHandler(router)
	ticketHandler.RegisterTicketHandler(router)

	reviewHandler.RegisterReviewHandler(router)

	catalogHandler.RegisterCatalogHandler(router)
	distContractHandler.RegisterDistributionContractHandler(router)
	contractHandler.RegisterContractHandler(router)
	contractItemHandler.RegisterContractItemHandler(router)
	supplierHandler.RegisterSupplierHandler(router)
	repertoireHandler.RegisterTheatreRepertoireHandler(router)
	notificationHandler.RegisterNotificationHandler(router)

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	// Create a new CORS middleware with desired options
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"}, // Adjust as needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Use the CORS middleware to wrap the router
	handler := corsMiddleware.Handler(router)

	// Start the server with the CORS-enabled handler
	log.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8085", handler))
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
