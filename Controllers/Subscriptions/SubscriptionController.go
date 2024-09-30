package subscriptions

import (
	"fmt"
	"net/http"
)

func NotifyNewSub(w http.ResponseWriter, r *http.Request) {
	fmt.Print("NEW SUBSCRIPTION")
}
