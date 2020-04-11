package connector

import (
	"fmt"
	"net/http"
)

type response struct {
	Status bool
	Message string
	Content interface{}
}

type route struct {
	method string
	endpoint string
}

type CoreConnector struct {
	baseUrl string
	protectedToken string

	resetRequestCountRoute route
	checkUserInactivityRoute route

	httpClient *http.Client
}

func NewCoreConnector(podIp string, port int, protectedToken string) CoreConnector {
	baseUrl := fmt.Sprintf("http://%s:%d", podIp, port)

	return CoreConnector{
		baseUrl: baseUrl,
		protectedToken: protectedToken,

		resetRequestCountRoute: route{
			method:   "DELETE",
			endpoint: "/subs",
		},
		checkUserInactivityRoute: route{
			method: "GET",
			endpoint: "/profile/inactivity",
		},

		httpClient: &http.Client{},
	}
}
