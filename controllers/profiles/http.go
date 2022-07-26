package profiles

import (
	"github.com/dedenfarhanhub/smart-koi-be/app/middleware"
	"github.com/dedenfarhanhub/smart-koi-be/business/users"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/profiles/request"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/profiles/response"
	basereponse "github.com/dedenfarhanhub/smart-koi-be/helper/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProfileController struct {
	userUseCase users.UseCase
}

func NewProfileController(uc users.UseCase) *ProfileController {
	return &ProfileController{
		userUseCase: uc,
	}
}

// FindByToken All godoc
// @Tags profile-controller
// @Summary Profile
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Success 200 {object} response.Profile
// @Router /profile [get]
func (controller *ProfileController) FindByToken(c echo.Context) error {
	ctx := c.Request().Context()
	id := middleware.GetUserId(c)
	user, err := controller.userUseCase.GetByID(ctx, id)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return basereponse.NewSuccessResponse(c, response.FromDomain(user))
}

// Update All godoc
// @Tags profile-controller
// @Summary Update
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param UpdateProfile body request.UpdateProfile true "UpdateProfile"
// @Accept json
// @Produce json
// @Success 200 {object} response.Profile
// @Router /profile [put]
func (controller *ProfileController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := middleware.GetUserId(c)
	req := request.UpdateProfile{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	user, err := controller.userUseCase.Update(ctx, req.ToUpdateDomain(), id)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return basereponse.NewSuccessResponse(c, response.FromDomain(user))
}