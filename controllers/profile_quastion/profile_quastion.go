package profile_quastion

import (
	"github.com/go-openapi/runtime/middleware"

	si "github.com/eure/si2018-second-half-3/restapi/summerintern"
)

func GetProfileQuastions(p si.GetProfileQuastionsParams) middleware.Responder {
  return getProfileQuastionsOKResponse()
}
