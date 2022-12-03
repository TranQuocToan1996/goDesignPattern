package structuralPatterns

import "net/http"

//Real subject

type Application struct {
}

func (a *Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return http.StatusOK, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return http.StatusAccepted, "User Created"
	}
	return http.StatusNotFound, "Not Ok"
}
