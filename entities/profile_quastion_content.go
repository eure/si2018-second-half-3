package entities

type ProfileQuastionContent struct {
	Message           string `xorm:"message"`
	Type              string `xorm:"type"`
	ProfileQuastionId int64  `xorm:"profile_quastion_element_id"` //xormがこれで合ってるか分からん
}
