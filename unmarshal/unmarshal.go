package unmarshal

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/goy-ex/apierror"
)

var Mapper apierror.MapperComponent = apierror.MapperFunc(func(err error) apierror.APIError {
	base := apierror.BadRequest
	if errors.Is(err, io.EOF) {
		base.Message = io.EOF.Error()
	} else if e, ok := errors.AsType[*json.SyntaxError](err); ok {
		base.Message = e.Error()
	} else if e, ok := errors.AsType[*json.UnmarshalTypeError](err); ok {
		base.Message = e.Error()
	} else if _, ok := errors.AsType[*json.InvalidUnmarshalError](err); ok {
		base = apierror.InternalServerError
	}

	return base
})
