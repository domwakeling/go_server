package handlers

import (
	"net/http"
)

// RootHandler provides a message advising to check docs at /docs
func (p *Products) RootHandler(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/docs", 302)
}