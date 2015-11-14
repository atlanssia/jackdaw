package utils
import (
    "net/http"
    "io"
)

func WriteError(w http.ResponseWriter, err error) {
    io.WriteString(w, err.Error())
}

func Write404(w http.ResponseWriter) {

    w.WriteHeader(http.StatusNotFound)
}