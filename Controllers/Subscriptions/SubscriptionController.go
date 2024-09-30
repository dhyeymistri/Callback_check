package subscriptions

import (
	"encoding/json"
	"net/http"
)

func NotifyNewSub(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("New subscription made")
}
