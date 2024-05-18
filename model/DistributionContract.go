package model

import (
	"errors"
	"time"
)

type ContractModel string

const (
	Bidding    ContractModel = "Bidding"
	Percentage ContractModel = "Percentage"
)

type DistributionContract struct {
	ID          uint                `gorm:"primary_key" json:"id"`
	CompanyID   uint                `gorm:"not null" json:"-"`
	Company     DistributionCompany `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MovieID     uint                `gorm:"not null" json:"movieId"`
	ManagerID   uint                `gorm:"not null" json:"managerId"`
	Manager     User                `gorm:"foreignKey:ManagerID" json:"-"`
	StartDate   Date                `gorm:"type:date" json:"startDate"`
	EndDate     Date                `gorm:"type:date" json:"endDate"`
	Model       ContractModel       `gorm:"type:contractmodel;not null" json:"model"`
	AgreedSum   *float32            `gorm:"" json:"agreedSum"`
	WeeklyCosts *float32            `gorm:"" json:"weeklyCosts"`
	Percentage  *float32            `gorm:"" json:"percentage"`
}

func (dc *DistributionContract) Validate() error {
	if dc.Company.Name == "" {
		return errors.New("company cannot be null")
	}
	if dc.StartDate.After(dc.EndDate) {
		return errors.New("startDate must be before endDate")
	}
	if dc.Model == Bidding && dc.AgreedSum == nil {
		return errors.New("agreedSum cannot be null when contract model is bidding")
	}
	if dc.Model == Bidding && (dc.WeeklyCosts != nil || dc.Percentage != nil) {
		return errors.New("weeklyCosts and percentage must be null when contract model is bidding")
	}
	if dc.Model == Percentage {
		if dc.AgreedSum != nil {
			return errors.New("agreedSum must be null when contract model is percentage")
		}
		if dc.WeeklyCosts == nil {
			return errors.New("weeklyCosts cannot be null when contract model is percentage")
		}
		if dc.Percentage == nil {
			return errors.New("percentage cannot be null when contract model is percentage")
		}
	}
	return nil
}

func (dc *DistributionContract) IsExpired() bool {
	return NewDate(time.Now()).After(dc.EndDate)
}

func (dc *DistributionContract) OverlapsWith(other *DistributionContract) bool {
	return !(dc.StartDate.After(other.EndDate) || dc.EndDate.Before(other.StartDate))
}

func (dc *DistributionContract) IsActive() bool {
	currentDate := NewDate(time.Now())
	return !currentDate.Before(dc.StartDate) && !currentDate.After(dc.EndDate)
}
