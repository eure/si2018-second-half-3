package main

import (
	"fmt"

	"github.com/eure/si2018-second-half-3/entities"
	"github.com/eure/si2018-second-half-3/repositories"
)

func dummyToken() {
	s := repositories.NewSession()
	r := repositories.NewUserTokenRepository(s)

	for i := firstUserID; i <= lastUserID; i++ {
		token := entities.UserToken{
			UserID: int64(i),
			Token:  fmt.Sprintf("USERTOKEN%v", i),
		}
		r.Create(token)
	}
}
