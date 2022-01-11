package gocardless

import "context"

type PagingIterator interface {
	Next() bool
	Value(context.Context) interface{}
}
