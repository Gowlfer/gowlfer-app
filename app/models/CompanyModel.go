package models

import (
	"errors"
	"github.com/gowlfer/gowlfer/internal/utils/database"
	"gorm.io/gorm"
)

type GowlferCompanyType struct {
	gorm.Model
	CompanyType string `gorm:"not null;unique_index"`
}

type GowlferCompany struct {
	gorm.Model
	CompanyName   string `gorm:"not null"`
	CompanyTypeID int64  `gorm:"not null"`
	CompanyOwner  string `gorm:"not null"`
}

func (c *GowlferCompany) CreateCompany(companyName, companyOwner string, companyTypeID int64) error {

	c.CompanyName = companyName
	c.CompanyTypeID = companyTypeID
	c.CompanyOwner = companyOwner

	created := database.DB.Create(&c)

	if created.Error != nil {
		return errors.New("failed to create company")
	}

	return nil
}

func (c *GowlferCompany) GetCompanies() ([]GowlferCompany, error) {
	var companies []GowlferCompany
	found := database.DB.Find(&companies)

	if found.Error != nil {
		return nil, errors.New("failed to get companies")
	}

	return companies, nil
}

func (c *GowlferCompany) GetCompany(id int64) (GowlferCompany, error) {
	var company GowlferCompany
	found := database.DB.Where("id = ?", id).Find(&company)

	if found.Error != nil {
		return company, errors.New("failed to get company")
	}

	return company, nil
}
