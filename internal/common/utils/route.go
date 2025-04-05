package utils

import (
	"net/http"
	"search-vector/internal/common/schema"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func Route[P any, R any](route func(P) (R, error)) func(c *gin.Context) {
	validate, err := schema.New()
	if err != nil {
		panic(err.Error())
	}

	return func(c *gin.Context) {
		var params P
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})

			return
		}

		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})

			return
		}

		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})

			return
		}

		if err := validate.Struct(params); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})

			return
		}

		response, err := route(params)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})

			return
		}

		c.JSON(http.StatusOK, response)
	}
}
