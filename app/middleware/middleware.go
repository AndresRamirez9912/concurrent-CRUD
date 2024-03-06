package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/constants"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/services"
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/utils"
)

func LimitGoroutines() gin.HandlerFunc {
	// The maximum amount of goroutines is the chan size x endpoint
	channel := make(chan bool, 100)

	return func(c *gin.Context) {
		select {
		case channel <- true: // Check randoml
		default:
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		defer func() { <-channel }()

		c.Next()
	}
}

func ValidateUser(active bool, auth services.AuthInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Active or deactivate the auth validation
		if !active {
			c.Next()
			return
		}

		token, err := c.Cookie(constants.TOKEN)
		if err != nil {
			errorResponse := utils.CreateErrorResponse(http.StatusBadRequest, "Auth not found, please signUp")
			c.JSON(http.StatusBadRequest, errorResponse)
			c.Abort()
			return
		}

		err = auth.ValidateUser(constants.VALIDATE_URL, token)
		if err != nil {
			errorResponse := utils.CreateErrorResponse(http.StatusMethodNotAllowed, "Invalid auth token, please logIn")
			c.JSON(http.StatusMethodNotAllowed, errorResponse)
			c.Abort()
			return
		}

		c.Next()
	}
}
