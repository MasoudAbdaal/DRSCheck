package DRS

import "net/http"

type TrackActions struct {
	isPassed           bool
	passedActions      string
	passedActionsCount uint8
}

func DRSHandler(w http.ResponseWriter, r *http.Request) {
	var trackActions TrackActions

	if r.Header.Get("Host") == "" {
		trackActions.isPassed = false
		trackActions.passedActionsCount = +1
		trackActions.passedActions += "\n 920350 Host header is a numeric IP address"

	}

}
