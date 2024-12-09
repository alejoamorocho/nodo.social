package function

import (
	"net/http"

	"github.com/kha0sys/nodo.social/internal/handlers"
)

// HelloWorld is an HTTP Cloud Function
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	handlers.HelloWorld(w, r)
}
