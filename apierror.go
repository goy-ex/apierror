package apierror

import (
	"slices"
)

type APIError struct {
	StatusCode int
	Message    string
}

type MapperComponent interface {
	Map(err error) APIError
}

type MapperComposite struct {
	rules []Rule
}

func NewMapperComposite(rules ...Rule) MapperComposite {
	newRules := make([]Rule, len(rules))
	copy(newRules, rules)

	return MapperComposite{rules: newRules}
}

func (mc MapperComposite) With(rules ...Rule) MapperComposite {
	newRules := make([]Rule, len(mc.rules)+len(rules))
	copy(newRules[:len(mc.rules)], mc.rules)
	copy(newRules[len(mc.rules):], rules)

	return MapperComposite{rules: newRules}
}

func (mc MapperComposite) Map(err error) APIError {
	for _, rule := range slices.Backward(mc.rules) {
		if rule.Match(err) {
			return rule.Mapper.Map(err)
		}
	}

	return InternalServerError
}

type MapperFunc func(error) APIError

func (f MapperFunc) Map(err error) APIError {
	return f(err)
}
