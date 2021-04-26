package middleware

import (
	"backend-services/handlers"
	"backend-services/services"
	"backend-services/services/wrapper"
	"strings"

	"github.com/gin-gonic/gin"
)

// import (
// 	"codebase/handlers"
// 	"codebase/services"
// 	"codebase/services/wrapper"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

func AdminAuthenticationRequired(handler handlers.Handler, requiredRoleLevels ...int) gin.HandlerFunc {
	const notAuthorizedError = "user not authorized, please login as admin"
	const notValidTokenError = "token is invalid or expired, please re-log"
	const insufficientRoleLevel = "user have an insufficient role level"

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(wrapper.StatusForbidden.New(notAuthorizedError, nil))
			c.Abort()
			return
		}

		authValues := strings.Split(authHeader, " ")
		if len(authValues) != 2 {
			c.JSON(wrapper.StatusForbidden.New(notAuthorizedError, nil))
			c.Abort()
			return
		}

		authType := strings.ToLower(authValues[0])
		token := authValues[1]
		if authType != "bearer" {
			c.JSON(wrapper.StatusForbidden.New(notAuthorizedError, nil))
			c.Abort()
			return
		}

		jwt := services.UserRoleJwt{}
		status := jwt.ParseToken(token, handler.Env.JwtSecretKey)
		if status == services.StatusTokenExpired {
			c.JSON(wrapper.StatusUnauthorized.New(notValidTokenError, nil))
			c.Abort()
			return
		} else if status != services.StatusTokenOK {
			c.JSON(wrapper.StatusForbidden.New(notValidTokenError, nil))
			c.Abort()
			return
		}

		// -- validate based on role for admin
		if jwt.RoleID != 1 {
			c.JSON(wrapper.StatusUnauthorized.New(notAuthorizedError, nil))
			c.Abort()
			return
		}

		c.Set("user_id", jwt.UserID)
		c.Set("role_id", jwt.RoleID)

		c.Next()
	}
}

func UserAuthenticationRequired(handler handlers.Handler, requiredRoleLevels ...int) gin.HandlerFunc {
	const notAuthorizedError = "user not authorized, please login as admin"
	const notValidTokenError = "token is invalid or expired, please re-log"
	const insufficientRoleLevel = "user have an insufficient role level"

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(wrapper.StatusForbidden.New(notAuthorizedError, nil))
			c.Abort()
			return
		}

		authValues := strings.Split(authHeader, " ")
		if len(authValues) != 2 {
			c.JSON(wrapper.StatusForbidden.New(notAuthorizedError, nil))
			c.Abort()
			return
		}

		authType := strings.ToLower(authValues[0])
		token := authValues[1]
		if authType != "bearer" {
			c.JSON(wrapper.StatusForbidden.New(notAuthorizedError, nil))
			c.Abort()
			return
		}

		jwt := services.UserRoleJwt{}
		status := jwt.ParseToken(token, handler.Env.JwtSecretKey)
		if status == services.StatusTokenExpired {
			c.JSON(wrapper.StatusUnauthorized.New(notValidTokenError, nil))
			c.Abort()
			return
		} else if status != services.StatusTokenOK {
			c.JSON(wrapper.StatusForbidden.New(notValidTokenError, nil))
			c.Abort()
			return
		}

		c.Set("user_id", jwt.UserID)
		c.Set("role_id", jwt.RoleID)

		c.Next()
	}
}
