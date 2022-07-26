package middleware

import (
	"errors"
	baseresponse "github.com/dedenfarhanhub/smart-koi-be/helper/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RoleValidation(roles []string) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := GetUser(c)
			for _, role := range roles{
				if claims.Role == role {
					return hf(c)
				}
			}

			return baseresponse.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden roles"))
		}
	}
}
