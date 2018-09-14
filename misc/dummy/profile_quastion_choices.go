package main

import (
	"github.com/eure/si2018-second-half-3/entities"
	"github.com/eure/si2018-second-half-3/repositories"
)

func dummyProfileQuastionChoices() {
	s := repositories.NewSession()
	r := repositories.NewProfileQuastionChoiceRepository(s)

	profileQuastionContents := []entities.ProfileQuastionChoice{
		{
			Content:                  "東京",
			ProfileQuastionElementId: 11,
		},
		{
			Content:                  "それいがい",
			ProfileQuastionElementId: 11,
		},
		{
			Content:                  "ムキムキ",
			ProfileQuastionElementId: 3,
		},
		{
			Content:                  "それいがい",
			ProfileQuastionElementId: 3,
		},
		{
			Content:                  "東京",
			ProfileQuastionElementId: 12,
		},
		{
			Content:                  "それいがい",
			ProfileQuastionElementId: 12,
		},
		{
			Content:                  "1億円",
			ProfileQuastionElementId: 15,
		},
		{
			Content:                  "それいがい",
			ProfileQuastionElementId: 15,
		},
		{
			Content:                  "180cm",
			ProfileQuastionElementId: 16,
		},
		{
			Content:                  "それいがい",
			ProfileQuastionElementId: 16,
		},
		{
			Content:                  "バツ1",
			ProfileQuastionElementId: 18,
		},
		{
			Content:                  "未婚",
			ProfileQuastionElementId: 18,
		},
		{
			Content:                  "いない",
			ProfileQuastionElementId: 19,
		},
		{
			Content:                  "それいがい",
			ProfileQuastionElementId: 19,
		},
		{
			Content:                  "今すぐ",
			ProfileQuastionElementId: 20,
		},
		{
			Content:                  "それいがい",
			ProfileQuastionElementId: 20,
		},
		{
			Content:                  "ほしい！",
			ProfileQuastionElementId: 21,
		},
		{
			Content:                  "ほしくない",
			ProfileQuastionElementId: 21,
		},
		{
			Content:                  "吸う",
			ProfileQuastionElementId: 22,
		},
		{
			Content:                  "吸わない",
			ProfileQuastionElementId: 22,
		},

		{
			Content:                  "飲む",
			ProfileQuastionElementId: 23,
		},
		{
			Content:                  "飲まない",
			ProfileQuastionElementId: 23,
		},

		{
			Content:                  "土曜",
			ProfileQuastionElementId: 24,
		},
		{
			Content:                  "それ以外",
			ProfileQuastionElementId: 24,
		},

		{
			Content:                  "メッセージを重ねてから",
			ProfileQuastionElementId: 25,
		},
		{
			Content:                  "それ以外",
			ProfileQuastionElementId: 25,
		},

		{
			Content:                  "6万円",
			ProfileQuastionElementId: 26,
		},
		{
			Content:                  "それ以外",
			ProfileQuastionElementId: 26,
		},
		{
			Content:                  "長男",
			ProfileQuastionElementId: 27,
		},
		{
			Content:                  "それ以外",
			ProfileQuastionElementId: 27,
		},

		{
			Content:                  "積極的に参加",
			ProfileQuastionElementId: 28,
		},
		{
			Content:                  "絶対しない",
			ProfileQuastionElementId: 28,
		},
	}

	for _, content := range profileQuastionContents {
		r.Create(content)
	}
}
