package main
import (
    "net/http"
    "io"
)

func writeError(w http.ResponseWriter, err error) {
    io.WriteString(w, err.Error())
}