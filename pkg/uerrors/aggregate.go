package uerrors

import (
	"errors"
	"strings"

	"gitlab.com/kyle_anderson/go-utils/pkg/linkedlist"
)

/* Represents multiple errors aggregated into a single error. */
type Aggregate interface {
	error
	Is(err error) bool
	As(target interface{}) bool
	/* Materializes an error out of the aggregate, returning nil if the aggregate is empty,
	or the aggregate itself if not. */
	Materialize() Aggregate
}

type LinkedAggregate struct {
	errs linkedlist.ForwardList[error]
}

func (a *LinkedAggregate) Error() string {
	builder := strings.Builder{}
	n := a.errs.First
	for {
		builder.WriteString(n.Item.Error())
		if n = n.Next; n == nil {
			break
		} else {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}

func (a *LinkedAggregate) Add(err error) {
	a.errs.Prepend(err)
}

func (a *LinkedAggregate) Materialize() Aggregate {
	if a.IsEmpty() {
		return nil
	}
	return a
}

func (a *LinkedAggregate) IsEmpty() bool {
	return a.errs.IsEmpty()
}

func (a *LinkedAggregate) Is(err error) bool {
	for n := a.errs.First; n != nil; n = n.Next {
		if errors.Is(n.Item, err) {
			return true
		}
	}
	return false
}

func (a *LinkedAggregate) As(target interface{}) bool {
	for n := a.errs.First; n != nil; n = n.Next {
		if errors.As(n.Item, target) {
			return true
		}
	}
	return false
}
