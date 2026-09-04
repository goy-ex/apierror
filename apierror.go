package apierror

import (
	"encoding/json"
	"errors"
	"net/http"
)

type APIError struct {
	StatusCode int
	Message    string
}

type Mapper interface {
	Map(err error) APIError
}

type MapperFunc func(err error) APIError

func (f MapperFunc) Map(err error) APIError {
	return f(err)
}

type Rule struct {
	Target error
	Mapper Mapper
}

type MapperChain struct {
	rules []Rule
}

func NewMapperChain(rules ...Rule) MapperChain {
	rulesCopy := make([]Rule, len(rules))
	copy(rulesCopy, rules)

	return MapperChain{
		rules: rulesCopy,
	}
}

func (r MapperChain) Map(err error) APIError {
	for _, rule := range r.rules {
		if errors.Is(err, rule.Target) {
			return rule.Mapper.Map(err)
		}
	}

	return APIError{
		StatusCode: http.StatusInternalServerError,
		Message:    http.StatusText(http.StatusInternalServerError),
	}
}

var UnmarshalResolver Mapper = MapperFunc(func(err error) APIError {
	base := APIError{
		StatusCode: http.StatusBadRequest,
		Message:    http.StatusText(http.StatusBadRequest),
	}
	if e, ok := errors.AsType[*json.SyntaxError](err); ok {
		base.Message = e.Error()
	}
	if e, ok := errors.AsType[*json.UnmarshalTypeError](err); ok {
		base.Message = e.Error()
	}
	if e, ok := errors.AsType[*json.InvalidUnmarshalError](err); ok {
		base.Message = e.Error()
	}

	return base
})
