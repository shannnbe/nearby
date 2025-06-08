package http

import (
	"fmt"

	"gofr.dev/pkg/gofr"
)

func groupV1(path string) string {
	return fmt.Sprintf("/%s/%s", "api/v1", path)
}

func RegisterRoute(route *gofr.App) {
	// Monitoring API
	// route.POST(groupV1("create-accident"), CreateAccidentHandler)

	// Reporting API
	// route.POST(groupV1("retrieve-accident"), RetrieveAccidentHandler)
}
