package profile_quastion

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/eure/si2018-second-half-3/libs/token"
	"github.com/eure/si2018-second-half-3/repositories"
	si "github.com/eure/si2018-second-half-3/restapi/summerintern"
)

func GetProfileQuastions(p si.GetProfileQuastionsParams) middleware.Responder {
	s := repositories.NewSession()
	t := p.Token

	user, err := token.GetUserByToken(s, t)
	if err != nil {
		return getProfileQuastionsInternalServerErrorResponse("userの取得に失敗")
	}
	if user == nil {
		return getProfileQuastionsUnauthorizedResponse("無効なトークン")
	}

	balkProfileList := user.GetBalkProfileItems()
	if len(balkProfileList) == 0 {
		return getProfileQuastionsOKResponse()
	}

	// 未記入項目をもとにProfileQuastionsElementを取得する
	r := repositories.NewProfileQuastionElementRepository(s)
	profileQuastionElemetens, err := r.FindByBalkProfileList(balkProfileList)

	if err != nil {
		return getProfileQuastionsInternalServerErrorResponse("ProfileQuastionsElementの取得失敗")
	}

	if len(profileQuastionElemetens) == 0 {
		return getProfileQuastionsInternalServerErrorResponse("ProfileQuastionsElementが追加されてない可能性があります")
	}
}
