package helpers

import (
	"io"
	"net/http"
)

//JSON encoder
func JSON(str string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, str)
	}
}
