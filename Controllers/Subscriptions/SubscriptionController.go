package subscriptions

import (
	"fmt"
	"io"
	"net/http"
)

func NotifyNewSub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//data is array of bytes
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error with data retrieval")
	}
	asString := string(data)
	fmt.Print(asString)
}
