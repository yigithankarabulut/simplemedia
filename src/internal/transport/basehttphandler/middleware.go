package basehttphandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yigithankarabulut/simplemedia/src/pkg/constant"
	"github.com/yigithankarabulut/simplemedia/src/pkg/util"
	"os"
	"strings"
)

func AuthMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var utils util.Util
		tokenStr := c.Get("Authorization")
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		if tokenStr == "" {
			return c.JSON(utils.BasicError(fmt.Sprintf(constant.Authenticate, "empty token"), fiber.StatusUnauthorized))
		}
		token, err := jwt.ParseWithClaims(tokenStr, &util.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return c.JSON(utils.BasicError(fmt.Sprintf(constant.Authenticate, err.Error()), fiber.StatusUnauthorized))
		}
		if !token.Valid {
			return c.JSON(utils.BasicError(fmt.Sprintf(constant.Authenticate, "invalid token"), fiber.StatusUnauthorized))
		}
		claims, ok := token.Claims.(*util.JwtCustomClaims)
		if !ok {
			return c.JSON(utils.BasicError(fmt.Sprintf(constant.Authenticate, "invalid token"), fiber.StatusUnauthorized))
		}
		c.Locals("userID", claims.UserID)
		return c.Next()
	}
}
