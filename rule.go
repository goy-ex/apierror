package apierror

import "errors"

type Rule struct {
	Match  func(err error) bool
	Mapper MapperComponent
}

func NewIsRule(target error, mapper MapperComponent) Rule {
	return Rule{
		Match: func(err error) bool {
			return errors.Is(err, target)
		},
		Mapper: mapper,
	}
}

func NewAsTypeRule[T error](mapper MapperComponent) Rule {
	return Rule{
		Match: func(err error) bool {
			_, ok := errors.AsType[T](err)
			return ok
		},
		Mapper: mapper,
	}
}
