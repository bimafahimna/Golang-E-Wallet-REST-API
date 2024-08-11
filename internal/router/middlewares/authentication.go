package middlewares

import (
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/pkgs/utils"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const jwtAuthMiddleware = "jwt authentication middleware"

func JWTAuth(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")

	if bearerToken == "" {
		err := apperrors.NewCustomError(nil, errormsg.ErrMsgFailedToAuthenticate, jwtAuthMiddleware)
		c.AbortWithStatusJSON(utils.StatusCode(err), utils.ResponseMsgBody(err, nil, nil))
		return
	}

	token := strings.Split(bearerToken, " ")
	if token[0] != "Bearer" {
		err := apperrors.NewCustomError(nil, errormsg.ErrMsgFailedToAuthenticate, jwtAuthMiddleware)
		c.AbortWithStatusJSON(utils.StatusCode(err), utils.ResponseMsgBody(err, nil, nil))
		return
	}

	jwtProvider := utils.NewJWTProviderHS256(os.Getenv("ISSUER"), os.Getenv("SECRET_KEY"))
	claim, err := jwtProvider.VerifyToken(token[1])
	if err != nil {
		err := apperrors.NewCustomError(nil, errormsg.ErrMsgFailedToAuthenticate, jwtAuthMiddleware)
		c.AbortWithStatusJSON(utils.StatusCode(err), utils.ResponseMsgBody(err, nil, nil))
		return
	}

	c.Set("user_id", claim.UserID)
	c.Next()
}
