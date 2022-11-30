package healthcheckhandler

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "Healthy")
}
