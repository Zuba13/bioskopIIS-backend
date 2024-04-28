package model

type DistributionCompany struct {
	ID        uint                   `gorm:"primary_key" json:"id"`
	Name      string                 `gorm:"not null" json:"name"`
	Contracts []DistributionContract `gorm:"foreignKey:CompanyID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION" json:"-"`
}
