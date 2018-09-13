package profile_quastion

import (
	"github.com/eure/si2018-second-half-3/models"
	si "github.com/eure/si2018-second-half-3/restapi/summerintern"
	"github.com/go-openapi/runtime/middleware"
)

func getProfileQuastionsOKResponse(profileQuastionElemetens []*models.ProfileQuastionElement) middleware.Responder {
	return si.NewGetProfileQuastionsOK().WithPayload(profileQuastionElemetens)
}

func getProfileQuastionsInternalServerErrorResponse(message string) middleware.Responder {
	return si.NewGetProfileQuastionsInternalServerError().WithPayload(
		&si.GetProfileQuastionsInternalServerErrorBody{
			Code:    "500",
			Message: "Internal Server Error :: " + message,
		})
}

func getProfileQuastionsUnauthorizedResponse(message string) middleware.Responder {
	return si.NewGetProfileQuastionsUnauthorized().WithPayload(
		&si.GetProfileQuastionsUnauthorizedBody{
			Code:    "401",
			Message: "Token Is Invalid:: " + message,
		})
}

func getProfileQuastionsBadRequestResponse(message string) middleware.Responder {
	return si.NewGetProfileQuastionsBadRequest().WithPayload(
		&si.GetProfileQuastionsBadRequestBody{
			Code:    "400",
			Message: "Bad Request:: " + message,
		})
}
