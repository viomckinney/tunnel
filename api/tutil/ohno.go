package tutil

import (
	"fmt"
	"net/http"
)

func OhNo(w http.ResponseWriter, err error) bool {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		fmt.Println(err.Error())
		return true
	}
	return false
}
