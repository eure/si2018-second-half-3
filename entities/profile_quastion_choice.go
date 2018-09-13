package entities

type ProfileQuastionChoice struct {
	Content                  string `xorm:"content"`
	ProfileQuastionElementId int64  `xorm:"profile_quastion_element_id"`
}
