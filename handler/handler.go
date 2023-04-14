package handler

import (
	"call-billing/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func APIHandlerStart(conf *model.Config) {
	// ======================================================================================
	// Initialization Echo webservice
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())
	// ======================================================================================
	// Handler
	mobileHandler := newMobileHandler(*conf)
	// ======================================================================================
	// Route
	/*
		Mobile Handler
	*/
	mobileAPI := e.Group("/mobile")
	mobileAPI.PUT("/:username/call", mobileHandler.PutUserCall)
	mobileAPI.GET("/:username/billing", mobileHandler.GetUserBilling)
	// ======================================================================================
	// Run
	e.Logger.Fatal(e.Start(conf.App.Address))
}
