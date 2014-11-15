package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, val interface{}) {
	b, _ := json.Marshal(val)
	fmt.Fprintf(w, string(b))
}
