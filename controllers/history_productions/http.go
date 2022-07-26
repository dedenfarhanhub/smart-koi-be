package history_productions

import (
	"github.com/dedenfarhanhub/smart-koi-be/app/middleware"
	"github.com/dedenfarhanhub/smart-koi-be/business/history_productions"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/history_productions/request"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/history_productions/response"
	"github.com/dedenfarhanhub/smart-koi-be/helper/covert_pointer"
	"github.com/dedenfarhanhub/smart-koi-be/helper/pagination"
	basereponse "github.com/dedenfarhanhub/smart-koi-be/helper/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HistoryProductionController struct {
	historyProductionUseCase history_productions.UseCase
}

func NewHistoryProductionController(hpc history_productions.UseCase) *HistoryProductionController {
	return &HistoryProductionController{
		historyProductionUseCase: hpc,
	}
}

// Fetch All godoc
// @Tags history-productions-controller
// @Summary Fetch
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param start_date query string false "start date"
// @Param end_date query string false "end date"
// @Param limit query string true "limit"
// @Param page query string true "page"
// @Param sort query string true "sort"
// @Accept json
// @Produce json
// @Success 200 {object} response.HistoryProductions
// @Router /history-productions [get]
func (controller *HistoryProductionController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	req := pagination.Pagination{}
	startDate  	:= c.QueryParam("start_date")
	endDate  	:= c.QueryParam("end_date")
	req.Limit 	= covert_pointer.ConvertStringInt(c.QueryParam("limit"))
	req.Page  	= covert_pointer.ConvertStringInt(c.QueryParam("page"))
	req.Sort  	= c.QueryParam("sort")
	paginationResponse, allHistoryProductions, err := controller.historyProductionUseCase.Fetch(ctx, req, startDate, endDate)
	paginationResponse.Contents = response.FromListDomain(allHistoryProductions)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, paginationResponse)
}

// Store All godoc
// @Tags history-productions-controller
// @Summary Store
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param HistoryProduction body request.HistoryProduction true "HistoryProduction"
// @Accept json
// @Produce json
// @Success 200 {object} response.HistoryProductions
// @Router /history-productions [post]
func (controller *HistoryProductionController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	userId := middleware.GetUserId(c)
	req := request.HistoryProduction{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	historyProduction, err := controller.historyProductionUseCase.Store(ctx, req.ToDomain(userId))
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, response.FromDomain(historyProduction))
}

// FindById All godoc
// @Tags history-productions-controller
// @Summary FindById
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param id path string true "historyProduction id"
// @Accept json
// @Produce json
// @Success 200 {object} response.HistoryProductions
// @Router /history-productions/{id} [get]
func (controller *HistoryProductionController) FindById(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.FindHistoryProductionByIdRequest{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	historyProduction, err := controller.historyProductionUseCase.GetByID(ctx, req.Id)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, response.FromDomain(historyProduction))
}

// Update All godoc
// @Tags history-productions-controller
// @Summary Update
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param UpdateHistoryProduction body request.UpdateHistoryProduction true "UpdateHistoryProduction"
// @Param id path string true "user id"
// @Accept json
// @Produce json
// @Success 200 {object} response.HistoryProductions
// @Router /history-productions/{id} [put]
func (controller *HistoryProductionController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	userId := middleware.GetUserId(c)
	req := request.UpdateHistoryProduction{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	historyProduction, err := controller.historyProductionUseCase.Update(ctx, req.ToUpdateDomain(userId), req.Id)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, response.FromDomain(historyProduction))
}

// Destroy All godoc
// @Tags history-productions-controller
// @Summary Destroy
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param id path string true "historyProduction id"
// @Accept json
// @Produce json
// @Success 200 {object} response.HistoryProductions
// @Router /history-productions/{id} [delete]
func (controller *HistoryProductionController) Destroy(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.FindHistoryProductionByIdRequest{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	historyProduction, err := controller.historyProductionUseCase.Destroy(ctx, req.Id)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, response.FromDomain(historyProduction))
}

// Barchart All godoc
// @Tags history-productions-controller
// @Summary Barchart
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Success 200 {object} response.HistoryProductions
// @Router /history-productions/barchart [get]
func (controller *HistoryProductionController) Barchart(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.FindHistoryProductionByIdRequest{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	historyProductions, err := controller.historyProductionUseCase.Barchart(ctx)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, response.FromListDomain(historyProductions))
}

// Stat All godoc
// @Tags history-productions-controller
// @Summary Stat
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Success 200 {object} response.HistoryProductionStats
// @Router /history-productions/stat [get]
func (controller *HistoryProductionController) Stat(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.FindHistoryProductionByIdRequest{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	historyProductionStatDomain, err := controller.historyProductionUseCase.Stat(ctx)
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, response.FromHistoryProductionStatDomain(historyProductionStatDomain))
}