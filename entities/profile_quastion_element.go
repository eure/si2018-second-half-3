package entities

import (
	"github.com/eure/si2018-second-half-3/models"
)

type ProfileQuastionElement struct {
	ID       int64  `xorm:"id"`
	Priority int64  `xorm:"priority"`
	Name     string `xorm:"name"`
	FormType string `xorm:"form_type"`

	Choices  []ProfileQuastionChoice  `xorm:"-"`
	Contents []ProfileQuastionContent `xorm:"-"`
}

func (p ProfileQuastionElement) Build() models.ProfileQuastionElement {
	profileQuastionContent := make(models.ProfileQuastionContent, 0)

	// 本来はentities.ProfileQuastionContentないでモデルへの変換をやるべきだけどここでやっちゃう
	for _, content := range p.Contents {
		contentItem := models.ProfileQuastionContentItems0{
			Message: content.Message,
			Type:    content.Type,
		}

		profileQuastionContent = append(profileQuastionContent, &contentItem)
	}

	// 本来はentities.ProfileQuastionChoiceないでモデルへの変換をやるべきだけどここでやっちゃう
	choiceContent := make([]string, len(p.Choices))

	for i, choice := range p.Choices {
		choiceContent[i] = choice.Content
	}

	return models.ProfileQuastionElement{
		Contents: profileQuastionContent,
		Choices:  choiceContent,
		FormType: p.FormType,
		Name:     p.Name,
		Priority: p.Priority,
	}
}

type ProfileQuastionElements []ProfileQuastionElement

func (elements *ProfileQuastionElements) Build() []*models.ProfileQuastionElement {
	var elements_model []*models.ProfileQuastionElement

	for _, e := range *elements {
		elementModel := e.Build()
		elements_model = append(elements_model, &elementModel)
	}

	return elements_model
}
