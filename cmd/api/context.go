package main

import (
	"context"
	"net/http"

	"github.com/Oskr11/A3_SE2201_ALIASKAR_UTESHKALIEV/internal/data"
)

type contextKey string

const userContextKey = contextKey("user")

func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

// contextGetUser retrieves the User struct from the request context. The only time that
// this helper should be used is when we logically expect there to be a User struct value
// in the context, and if it doesn't exist it will firmly be an 'unexpected' error, upon we panic.
func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}

	return user
}
