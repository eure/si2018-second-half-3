package main

import (
	"github.com/eure/si2018-second-half-3/entities"
	"github.com/eure/si2018-second-half-3/repositories"
)

func dummyProfileQuastion() {
	s := repositories.NewSession()
	r := repositories.NewProfileQuastionElementRepository(s)

	count := 24 //elementの項目数

	elementList := []string{
		"gender",
		"birthday",
		"nickname",
		"tweet",
		"introduction",
		"residence_state",
		"home_state",
		"education",
		"job",
		"annual_income",
		"height",
		"body_build",
		"marital_status",
		"child",
		"when_marry",
		"want_child",
		"smoking",
		"drinking",
		"holiday",
		"how_to_meet",
		"cost_of_date",
		"nth_child",
		"housework",
		"hobby",
	}

	for i := 0; i < count; i++ {
		profile := entities.ProfileQuastionElement{
			Name:     elementList[i],
			Priority: 1, //とりあえず1でやっとく,priorityの意味ないけど
		}
		r.Create(profile)
	}
}
