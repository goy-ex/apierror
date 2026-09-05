package apierror

import "net/http"

var InternalServerError = APIError{
	StatusCode: http.StatusInternalServerError,
	Message:    http.StatusText(http.StatusInternalServerError),
}

var BadRequest = APIError{
	StatusCode: http.StatusBadRequest,
	Message:    http.StatusText(http.StatusBadRequest),
}
