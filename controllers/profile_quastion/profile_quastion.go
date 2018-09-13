package profile_quastion

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/eure/si2018-second-half-3/entities"
	"github.com/eure/si2018-second-half-3/libs/token"
	"github.com/eure/si2018-second-half-3/models"
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
		var elements []*models.ProfileQuastionElement
		return getProfileQuastionsOKResponse(elements)
	}

	// 未記入項目をもとにProfileQuastionsElementを取得する
	r := repositories.NewProfileQuastionElementRepository(s)
	var profileQuastionElemetens entities.ProfileQuastionElements

	profileQuastionElemetens, err = r.FindByBalkProfileList(balkProfileList)

	if err != nil {
		return getProfileQuastionsInternalServerErrorResponse("ProfileQuastionsElementの取得失敗")
	}

	if len(profileQuastionElemetens) == 0 {
		return getProfileQuastionsInternalServerErrorResponse("ProfileQuastionsElementが追加されてない可能性があります")
	}

	contentRepository := repositories.NewProfileQuastionContentRepository(s)

	//FIXME N + 1起きるけど許して♡
	for i, element := range profileQuastionElemetens {
		profileQuastionElemetens[i].Contents, err = contentRepository.FindByProfileQuastionId(element.ID)
		if err != nil {
			return getProfileQuastionsInternalServerErrorResponse("内部エラー")
		}
	}

	profileQuastionElemetensModel := profileQuastionElemetens.Build()

	return getProfileQuastionsOKResponse(profileQuastionElemetensModel)
}
