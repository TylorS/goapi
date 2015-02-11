package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Chain struct {
	constructors []Middleware
}

func New(middlewares ...Middleware) Chain {
	c := Chain{}
	c.constructors = append(c.constructors, middlewares...)
	return c
}

func (c Chain) Then(h http.Handler) http.Handler {
	var final http.Handler
	if h != nil {
		final = h
	} else {
		final = http.DefaultServeMux
	}

	for i := len(c.constructors) - 1; i >= 0; i-- {
		final = c.constructors[i](final)
	}

	return final

}

func (c Chain) ThenFunc(fn http.HandlerFunc) http.Handler {
	if fn == nil {
		return c.Then(nil)
	}

	return c.Then(http.HandlerFunc(fn))

}

func (c Chain) Append(middlewares ...Middleware) Chain {
	newMiddlewares := make([]Middleware, len(c.constructors)+len(middlewares))
	copy(newMiddlewares, c.constructors)
	copy(newMiddlewares[len(c.constructors):], middlewares)

	newChain := New(newMiddlewares...)
	return newChain
}
