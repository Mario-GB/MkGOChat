package goserver

import (
	"fmt"
	"net/http"
)

func StartWeb() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error listen server")
	}
}
