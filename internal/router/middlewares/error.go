package middlewares

import (
	"errors"
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/pkgs/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorMiddleware(c *gin.Context) {
	utils.RegisterJSONTagName()
	c.Next()
	if len(c.Errors) > 0 {
		err := c.Errors[0].Err

		switch e := err.(type) {
		case validator.ValidationErrors:
			var ve validator.ValidationErrors
			errors.As(err, &ve)
			newErr := apperrors.NewCustomValidationErrors(e)
			c.AbortWithStatusJSON(utils.StatusCode(newErr), utils.ResponseMsgBody(newErr, nil, nil))
		case error:
			newErr, ok := err.(*apperrors.CustomError)
			if ok {
				c.AbortWithStatusJSON(utils.StatusCode(newErr), utils.ResponseMsgBody(newErr, nil, nil))
				return
			}
			c.AbortWithStatusJSON(utils.StatusCode(newErr), utils.ResponseMsgBody(err, nil, nil))
		default:
			c.AbortWithStatusJSON(utils.StatusCode(e), utils.ResponseMsgBody(err, nil, nil))
		}

	}
}
