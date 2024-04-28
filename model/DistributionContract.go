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
	Company     DistributionCompany `gorm:"foreignKey:CompanyID" json:"company"`
	ManagerID   uint                `gorm:"not null" json:"managerId"`
	Manager     User                `gorm:"foreignKey:ManagerID" json:"-"`
	StartDate   time.Time           `gorm:"type:date" json:"startDate"`
	EndDate     time.Time           `gorm:"type:date" json:"endDate"`
	Model       ContractModel       `gorm:"type:enum('Bidding','Percentage');not null" json:"type"`
	AgreedSum   *float32            `gorm:"" json:"agreedSum"`
	WeeklyCosts *float32            `gorm:"" json:"weeklyCosts"`
	Percentage  *float32            `gorm:"" json:"percentage"`
}

func (dc *DistributionContract) Validate() error {
	if dc.StartDate.After(dc.EndDate) {
		return errors.New("startDate must be before endDate")
	}
	if dc.Model == Bidding && dc.AgreedSum == nil {
		return errors.New("agreedSum cannot be null when contract type is bidding")
	}
	if dc.Model == Percentage {
		if dc.WeeklyCosts == nil {
			return errors.New("weeklyCosts cannot be null when contract type is percentage")
		}
		if dc.Percentage == nil {
			return errors.New("percentage cannot be null when contract type is percentage")
		}
	}
	return nil
}
