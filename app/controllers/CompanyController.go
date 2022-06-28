package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gowlfer/gowlfer/app/models"
	"strconv"
)

func GetCompanies(ctx *fiber.Ctx) error {
	var comp models.GowlferCompany
	companies, err := comp.GetCompanies()
	if err != nil {
		return err
	}
	return ctx.JSON(companies)
}

func GetCompany(ctx *fiber.Ctx) error {
	var comp models.GowlferCompany
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	company, err := comp.GetCompany(int64(id))
	if err != nil {
		return err
	}
	return ctx.JSON(company)
}

func CreateCompany(ctx *fiber.Ctx) error {
	data := struct {
		CompanyName   string `json:"company_name"`
		CompanyTypeID string `json:"company_type"`
		CompanyOwner  string `json:"company_owner"`
	}{}

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var comp models.GowlferCompany

	id, _ := strconv.ParseInt(data.CompanyTypeID, 10, 64)

	if err := comp.CreateCompany(data.CompanyName, data.CompanyOwner, id); err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Company Created",
	})
}
