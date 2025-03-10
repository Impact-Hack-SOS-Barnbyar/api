package response

import "net/http"

type NotFoundProblemDetails struct {
	ProblemDetails
}

func NewNotFoundProblemDetails() *InternalServerErrorProblemDetails {
	return &InternalServerErrorProblemDetails{
		ProblemDetails: ProblemDetails{
			Status:  http.StatusNotFound,
			Type:    typeUrlPattern + "404",
			Title:   "Not Found",
			Details: "The requested resource could not be found",
		},
	}
}
