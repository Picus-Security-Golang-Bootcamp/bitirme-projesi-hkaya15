package mw

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/config"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/base/jwt"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/api/model"
	"go.uber.org/zap"
)
// AuthorizationMiddleware checks that user is admin
func AuthorizationMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims, err := VerifyTokenMiddleware(c.GetHeader("Authorization"), cfg)
			if err != nil {
				zap.L().Debug("jwt.verifytokenMiddleware: decodedClaims ", zap.Error(err))
			}
			if decodedClaims != nil && decodedClaims.Role==true{
				c.Next()
				c.Abort()
				return
			}
			
			c.JSON(http.StatusForbidden, APIResponse{Code: http.StatusForbidden,Message: "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, APIResponse{Code: http.StatusUnauthorized,Message: "You are not authorized!"})
		}
		c.Abort()
		return
	}
}
