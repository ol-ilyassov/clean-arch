package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ErrorResponse struct {
	ID     uuid.UUID   `json:"id"`
	Error  string      `json:"message,omitempty"`
	Errors []string    `json:"errors,omitempty"`
	Info   interface{} `json:"info,omitempty"`
}

func SetError(c *gin.Context, statusCode int, errs ...error) {
	var response = ErrorResponse{
		ID: uuid.New(),
	}

	if len(errs) == 0 {
		return
	}

	if len(errs) > 0 {
		response.Error = errs[0].Error()

		if len(errs) > 1 {
			for _, err := range errs {
				response.Errors = append(response.Errors, err.Error())
			}
		}
	}

	c.JSON(statusCode, response)
}
