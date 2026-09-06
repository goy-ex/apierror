package validator

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/goy-ex/apierror"
)

var Mapper apierror.MapperComponent = apierror.MapperFunc(func(err error) apierror.APIError {
	if _, ok := errors.AsType[*validator.InvalidValidationError](err); ok {
		return apierror.InternalServerError
	}

	if errs, ok := errors.AsType[validator.ValidationErrors](err); ok {
		var sb strings.Builder
		for i, e := range errs {
			sb.WriteString(e.Error())
			if i != len(errs)-1 {
				sb.WriteString("; ")
			}
		}

		return apierror.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    sb.String(),
		}
	}

	return apierror.BadRequest
})
