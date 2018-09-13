package entities

type ProfileQuastionElement struct {
	ID       int64  `xorm:"id"`
	Priority int64  `xorm:"priority"`
	Name     string `xorm:"name"`
}
