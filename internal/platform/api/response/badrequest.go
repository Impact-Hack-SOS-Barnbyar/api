package response

import "net/http"

type BadRequestProblemDetails struct {
	ProblemDetails
	Errors []ProblemDetailsError
}

func NewBadRequestProblemDetails(er []ProblemDetailsError) *BadRequestProblemDetails {
	return &BadRequestProblemDetails{
		ProblemDetails: ProblemDetails{
			Status:  http.StatusBadRequest,
			Type:    typeUrlPattern + "400",
			Title:   "Bad Request",
			Details: "The request produced one or more errors",
		},
		Errors: er,
	}
}
