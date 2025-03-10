package response

import "net/http"

type InternalServerErrorProblemDetails struct {
	ProblemDetails
}

func NewInternalServerErrorProblemDetails() *InternalServerErrorProblemDetails {
	return &InternalServerErrorProblemDetails{
		ProblemDetails{
			Status:  http.StatusInternalServerError,
			Type:    typeUrlPattern + "500",
			Title:   "Internal Server Error",
			Details: "An unexpected error occured on the server and has been logged",
		},
	}
}
