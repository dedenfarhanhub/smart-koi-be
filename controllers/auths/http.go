package auths

import (
	"github.com/dedenfarhanhub/smart-koi-be/business/users"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/auths/request"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/auths/response"
	basereponse "github.com/dedenfarhanhub/smart-koi-be/helper/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	userUseCase users.UseCase
}

func NewAuthController(uc users.UseCase) *AuthController {
	return &AuthController{
		userUseCase: uc,
	}
}

// Login All godoc
// @Tags auths-controller
// @Summary Login
// @Description Put all mandatory parameter
// @Param LoginRequest body request.LoginRequest true "LoginRequest"
// @Accept json
// @Produce json
// @Success 200 {object} response.LoginResponse
// @Router /auth/login [post]
func (controller *AuthController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var userLogin request.LoginRequest
	if err := c.Bind(&userLogin); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, _, err := controller.userUseCase.Login(ctx, userLogin.Username, userLogin.Password, false)

	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result := response.LoginResponse{}
	result.Token = token

	return basereponse.NewSuccessResponse(c, result)
}