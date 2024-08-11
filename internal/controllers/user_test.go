package controllers_test

import (
	"encoding/json"
	"fmt"
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/controllers"
	"golang-e-wallet-rest-api/internal/dtos"
	"golang-e-wallet-rest-api/internal/mocks"
	"golang-e-wallet-rest-api/internal/pkgs/utils"
	"golang-e-wallet-rest-api/internal/router/middlewares"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewUserController(t *testing.T) {
	t.Run("Should return not nil userController struct pointer", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockUserService := mocks.NewUserService(t)

		userController := controllers.NewUserController(mockUserService)

		assert.NotNil(t, userController)
	})
}

