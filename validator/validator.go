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
		sb := strings.Builder{}
		for _, e := range errs {
			sb.WriteString("-")
			sb.WriteString(e.Error())
		}

		return apierror.APIError{
			StatusCode: http.StatusBadRequest,
			Message:    sb.String(),
		}
	}

	return apierror.BadRequest
})
