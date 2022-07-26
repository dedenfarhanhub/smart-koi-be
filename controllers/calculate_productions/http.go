package calculate_productions

import (
	"github.com/dedenfarhanhub/smart-koi-be/app/middleware"
	"github.com/dedenfarhanhub/smart-koi-be/business/calculate_productions"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/calculate_productions/request"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/calculate_productions/response"
	basereponse "github.com/dedenfarhanhub/smart-koi-be/helper/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CalculateProductionController struct {
	calculateProductionUseCase calculate_productions.UseCase
}

func NewCalculateProductionController(cpc calculate_productions.UseCase) *CalculateProductionController {
	return &CalculateProductionController{
		calculateProductionUseCase: cpc,
	}
}

// Store All godoc
// @Tags calculate-productions-controller
// @Summary Store
// @Description Put all mandatory parameter
// @Param Authorization header string true "Bearer"
// @Param CalculateProduction body request.CalculateProduction true "CalculateProduction"
// @Accept json
// @Produce json
// @Success 200 {object} response.CalculateProductions
// @Router /calculate-productions [post]
func (controller *CalculateProductionController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	userId := middleware.GetUserId(c)
	req := request.CalculateProduction{}
	if err := c.Bind(&req); err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	historyProduction, err := controller.calculateProductionUseCase.Store(ctx, req.ToDomain(userId))
	if err != nil {
		return basereponse.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return basereponse.NewSuccessResponse(c, response.FromDomain(historyProduction))
}