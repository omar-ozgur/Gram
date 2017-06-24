package middleware

import (
	"fmt"
	"github.com/omar-ozgur/gram/utilities"
	"net/http"
)

func CustomMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	utilities.Logger.Info("Started request")
	next(rw, r)
	utilities.Logger.Info("Got response")
	fmt.Println()
}
