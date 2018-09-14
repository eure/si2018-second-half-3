package main

import (
	"github.com/eure/si2018-second-half-3/entities"
	"github.com/eure/si2018-second-half-3/repositories"
)

func dummyProfileQuastionContent() {
	s := repositories.NewSession()
	r := repositories.NewProfileQuastionContentRepository(s)

	profileQuastionContents := []entities.ProfileQuastionContent{
		{
			Message:           "Q あなたの性別を教えてください",
			Type:              "main",
			ProfileQuastionId: 1,
		},
		{
			Message:           "Q あなたの誕生日を教えてください",
			Type:              "main",
			ProfileQuastionId: 2,
		},
		{
			Message:           "Q あなたのニックネームを教えてください",
			Type:              "main",
			ProfileQuastionId: 3,
		},
		{
			Message:           "Q つぶやきを入力しましょう",
			Type:              "main",
			ProfileQuastionId: 4,
		},
		{
			Message:           "つぶやきとは、短い投稿のことです。\n人柄がわかるようなつぶやきを入力しましょう。\n例：「クラブにいきませんか」、「肉食べたい」",
			Type:              "sub",
			ProfileQuastionId: 4,
		},
		{
			Message:           "Q 自己紹介を入力しましょう",
			Type:              "main",
			ProfileQuastionId: 5,
		},
		{
			Message:           "Q 居住地を教えてください",
			Type:              "main",
			ProfileQuastionId: 6,
		},
		{
			Message:           "Q 出身地を教えてください",
			Type:              "main",
			ProfileQuastionId: 7,
		},
		{
			Message:           "Q 学歴を教えてください",
			Type:              "main",
			ProfileQuastionId: 8,
		},
		{
			Message:           "Q 職業を教えてください",
			Type:              "main",
			ProfileQuastionId: 9,
		},
		{
			Message:           "Q 年収を教えてください",
			Type:              "main",
			ProfileQuastionId: 10,
		},
		{
			Message:           "Q 身長を教えてください",
			Type:              "main",
			ProfileQuastionId: 11,
		},
		{
			Message:           "Q 体型を教えてください",
			Type:              "main",
			ProfileQuastionId: 12,
		},
		{
			Message:           "Q 結婚歴を教えてください",
			Type:              "main",
			ProfileQuastionId: 13,
		},
		{
			Message:           "Q 子供がいるかを教えてください",
			Type:              "main",
			ProfileQuastionId: 14,
		},
		{
			Message:           "Q いつ結婚したいのか教えてください",
			Type:              "main",
			ProfileQuastionId: 15,
		},
		{
			Message:           "Q 子供がほしいかを教えてください",
			Type:              "main",
			ProfileQuastionId: 16,
		},
		{
			Message:           "Q タバコを吸いますか？",
			Type:              "main",
			ProfileQuastionId: 17,
		},
		{
			Message:           "Q お酒を飲みますか?",
			Type:              "main",
			ProfileQuastionId: 18,
		},
		{
			Message:           "Q 休日を教えてください",
			Type:              "main",
			ProfileQuastionId: 19,
		},
		{
			Message:           "Q 出会うまでの希望を教えてください",
			Type:              "main",
			ProfileQuastionId: 20,
		},
		{
			Message:           "Q デートにいくらぐらい賭けますか",
			Type:              "main",
			ProfileQuastionId: 21,
		},
		{
			Message:           "Q 兄弟、姉妹関係を教えてください",
			Type:              "main",
			ProfileQuastionId: 22,
		},
		{
			Message:           "Q 家事育児には積極的に参加しますか？",
			Type:              "main",
			ProfileQuastionId: 23,
		},
		{
			Message:           "Q 趣味を教えてください",
			Type:              "main",
			ProfileQuastionId: 24,
		},
	}

	for _, content := range profileQuastionContents {
		r.Create(content)
	}
}
