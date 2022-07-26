package base_reponse

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseResponse struct {
	Code 	  int 	  	  `json:"code"`
	Messages  []string 	  `json:"messages,omitempty"`
	Data 	  interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Code = 200
	response.Messages = []string { "Success"}
	response.Data = param

	return c.JSON(http.StatusOK, response)
}


func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Code = status
	response.Messages = []string{err.Error()}

	return c.JSON(status, response)
}
