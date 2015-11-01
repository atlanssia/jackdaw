package utils
import (
    "net/http"
    "io"
)

func WriteError(w http.ResponseWriter, err error) {
    io.WriteString(w, err.Error())
}