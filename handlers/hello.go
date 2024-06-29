package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{
		l: l,
	}
}

// ServeHTTP is the HandleFunc params func which we are passing for a specific "route path"
// Even though we are not calling this ServeHTTP func but since it is implemented the http.Handler interface
// It will doing all its job under the hood
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	h.l.Println("Hello World from ServeHTTP!")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops!", http.StatusBadRequest)

		//? below codes will do same as above code
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("oops!"))
		return
	}

	fmt.Fprintf(w, "Hello %s\n", d)

}
