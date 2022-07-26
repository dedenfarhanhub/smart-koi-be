package users

import (
	"github.com/dedenfarhanhub/smart-koi-be/business/users"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/users/request"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/users/response"
	"github.com/dedenfarhanhub/smart-koi-be/helper/covert_pointer"
	"github.com/dedenfarhanhub/smart-koi-be/helper/pagination"
	basereponse "github.com/dedenfarhanhub/smart-koi-be/helper/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userUseCase users.UseCase
}

func NewUserController(uc users.UseCase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

// Fetch All godoc
// @Tags users-controller
// @Summary Fetch
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param keyword query string false "keyword"
// @Param limit query string true "limit"
// @Param page query string true "page"
// @Param sort query string true "sort"
// @Accept json
// @Produce json
// @Success 200 {object} response.Users
// @Router /users [get]
func (controller *UserController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	req := pagination.Pagination{}
	keyword  := c.QueryParam("keyword")
	req.Limit = covert_pointer.ConvertStringInt(c.QueryParam("limit"))
	req.Page  = covert_pointer.ConvertStringInt(c.QueryParam("page"))
	req.Sort  = c.QueryParam("sort")
	paginationResponse, allUsers, err := controller.userUseCase.Fetch(ctx, req, keyword)
	paginationResponse.Contents = response.FromListDomain(allUsers)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, paginationResponse)
}

// Store All godoc
// @Tags users-controller
// @Summary Store
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param UserRequest body request.Users true "UserRequest"
// @Accept json
// @Produce json
// @Success 200 {object} response.Users
// @Router /users [post]
func (controller *UserController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	user, err := controller.userUseCase.Store(ctx, req.ToDomain(), false)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, response.FromDomain(user))
}

// FindById All godoc
// @Tags users-controller
// @Summary FindById
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param id path string true "user id"
// @Accept json
// @Produce json
// @Success 200 {object} response.Users
// @Router /users/{id} [get]
func (controller *UserController) FindById(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.FindUserByIdRequest{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	user, err := controller.userUseCase.GetByID(ctx, req.Id)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, response.FromDomain(user))
}

// Update All godoc
// @Tags users-controller
// @Summary Update
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param UpdateUsers body request.UpdateUsers true "UpdateUsers"
// @Param id path string true "user id"
// @Accept json
// @Produce json
// @Success 200 {object} response.Users
// @Router /users/{id} [put]
func (controller *UserController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.UpdateUsers{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	user, err := controller.userUseCase.Update(ctx, req.ToUpdateDomain(), req.Id)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, response.FromDomain(user))
}

// Destroy All godoc
// @Tags users-controller
// @Summary Destroy
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param id path string true "user id"
// @Accept json
// @Produce json
// @Success 200 {object} response.Users
// @Router /users/{id} [delete]
func (controller *UserController) Destroy(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.FindUserByIdRequest{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	user, err := controller.userUseCase.Destroy(ctx, req.Id)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, response.FromDomain(user))
}