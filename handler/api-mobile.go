package handler

import (
	"net/http"
	"strings"
	"time"

	"call-billing/adapter/repository"
	"call-billing/model"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type MobileHandler struct {
	repo   repository.Repository
	config model.Config
}

func newMobileHandler(conf model.Config) IMobile {
	handler := &MobileHandler{config: conf}
	if repo, err := repository.New(conf.Adapter.Mongo); err != nil {
		// Stop API
		panic(err)
	} else {
		handler.repo = repo
	}
	// Success
	return handler
}

func (inst *MobileHandler) PutUserCall(c echo.Context) error {
	var body model.RequestUserCall
	// Validate
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}
	if body.CallDuration <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid call duration"})
	}
	body.Username = strings.TrimSpace(body.Username)
	if body.Username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid username"})
	}
	// Process
	if err := inst.repo.Record().Insert(&model.Record{
		Created:  time.Now().UnixMilli(),
		Username: body.Username,
		Duration: body.CallDuration,
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error"})
	}
	// Success
	return c.JSON(http.StatusCreated, map[string]string{"message": "success"})
}

func (inst *MobileHandler) GetUserBilling(c echo.Context) error {
	var body model.RequestUser
	// Validate
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}
	body.Username = strings.TrimSpace(body.Username)
	if body.Username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid username"})
	}
	// Process
	documents, err := inst.repo.Record().FindAll(&bson.M{"username": body.Username}, []string{}, 0)
	if err != nil {
		if !strings.Contains(err.Error(), "no documents") {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error"})
		}
	}
	var block int64 = 0
	for _, document := range *documents {
		block += document.Duration / int64(30000)
		if document.Duration%int64(30000) > 0 {
			block += 1
		}
	}
	// Success
	return c.JSON(http.StatusOK, map[string]interface{}{"call_count": len(*documents), "block_count": block})
}
