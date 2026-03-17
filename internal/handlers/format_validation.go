package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// BindAndValidate binds the request body and validates it in one step.
// Returns false if it fails, and writes the error response itself —
// so the handler just checks the bool and returns.
func BindAndValidate[T any](c *gin.Context, v *validator.Validate, dst *T) bool {
	if err := c.ShouldBindJSON(dst); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "malformed request body",
			"details": err.Error(),
		})
		return false
	}

	if err := v.Struct(dst); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "validation failed",
			"details": formatValidationErrors(err),
		})
		return false
	}

	return true
}

// Formats go-playground/validator errors into a clean slice
func formatValidationErrors(err error) []gin.H {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return []gin.H{{"field": "unknown", "message": err.Error()}}
	}

	out := make([]gin.H, len(ve))
	for i, fe := range ve {
		out[i] = gin.H{
			"field":   fe.Field(),
			"message": fe.Tag(),
		}
	}
	return out
}
