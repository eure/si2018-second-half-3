package entities

import (
	"github.com/eure/si2018-second-half-3/models"
	"github.com/go-openapi/strfmt"
)

type User struct {
	ID             int64           `xorm:"id"`
	Gender         string          `xorm:"gender"`
	Birthday       strfmt.Date     `xorm:"birthday"`
	Nickname       string          `xorm:"nickname"`
	Tweet          string          `xorm:"tweet"`
	Introduction   string          `xorm:"introduction"`
	ResidenceState string          `xorm:"residence_state"`
	HomeState      string          `xorm:"home_state"`
	Education      string          `xorm:"education"`
	Job            string          `xorm:"job"`
	AnnualIncome   string          `xorm:"annual_income"`
	Height         string          `xorm:"height"`
	BodyBuild      string          `xorm:"body_build"`
	MaritalStatus  string          `xorm:"marital_status"`
	Child          string          `xorm:"child"`
	WhenMarry      string          `xorm:"when_marry"`
	WantChild      string          `xorm:"want_child"`
	Smoking        string          `xorm:"smoking"`
	Drinking       string          `xorm:"drinking"`
	Holiday        string          `xorm:"holiday"`
	HowToMeet      string          `xorm:"how_to_meet"`
	CostOfDate     string          `xorm:"cost_of_date"`
	NthChild       string          `xorm:"nth_child"`
	Housework      string          `xorm:"housework"`
	Hobby          string          `xorm:"hobby"`
	CreatedAt      strfmt.DateTime `xorm:"created_at"`
	UpdatedAt      strfmt.DateTime `xorm:"updated_at"`

	ImageURI string `xorm:"-"`
}

func (u User) Build() models.User {
	return models.User{
		ID:             u.ID,
		Gender:         u.Gender,
		Birthday:       u.Birthday,
		Nickname:       u.Nickname,
		ImageURI:       u.ImageURI,
		Tweet:          u.Tweet,
		Introduction:   u.Introduction,
		ResidenceState: u.ResidenceState,
		HomeState:      u.HomeState,
		Education:      u.Education,
		Job:            u.Job,
		AnnualIncome:   u.AnnualIncome,
		Height:         u.Height,
		BodyBuild:      u.BodyBuild,
		MaritalStatus:  u.MaritalStatus,
		Child:          u.Child,
		WhenMarry:      u.WhenMarry,
		WantChild:      u.WantChild,
		Smoking:        u.Smoking,
		Drinking:       u.Drinking,
		Holiday:        u.Holiday,
		HowToMeet:      u.HowToMeet,
		CostOfDate:     u.CostOfDate,
		NthChild:       u.NthChild,
		Housework:      u.Housework,
		Hobby:          u.Hobby,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}

func (u User) GetOppositeGender() string {
	if u.Gender == "F" {
		return "M"
	}
	return "F"
}

type Users []User

func (users *Users) Build() []*models.User {
	var sUsers []*models.User

	for _, u := range *users {
		swaggerUser := u.Build()
		sUsers = append(sUsers, &swaggerUser)
	}
	return sUsers
}

// ユーザーの未記入項目のリストを返す関数
// 時間がないので、クソコードで行く
func (u User) GetBalkProfileItems() []string {
	var balkProfileItems []string

	// FIXME Birthdayの空判定がぱっと出てこないので後回し
	//  if u.Birthday == strfmt.Date {
	//    balkProfileItems = append(balkProfileItems, "Birthday")
	//  }
	if !isStringPresent(u.Gender) {
		balkProfileItems = append(balkProfileItems, "gender")
	}
	if !isStringPresent(u.Nickname) {
		balkProfileItems = append(balkProfileItems, "nickname")
	}
	if isStringPresent(u.ImageURI) {
		balkProfileItems = append(balkProfileItems, "image_uri")
	}
	if !isStringPresent(u.Tweet) {
		balkProfileItems = append(balkProfileItems, "tweet")
	}
	if !isStringPresent(u.Introduction) {
		balkProfileItems = append(balkProfileItems, "introduction")
	}
	if !isStringPresent(u.ResidenceState) {
		balkProfileItems = append(balkProfileItems, "residence_state")
	}
	if !isStringPresent(u.HomeState) {
		balkProfileItems = append(balkProfileItems, "home_state")
	}
	if !isStringPresent(u.Education) {
		balkProfileItems = append(balkProfileItems, "education")
	}
	if !isStringPresent(u.Job) {
		balkProfileItems = append(balkProfileItems, "job")
	}
	if !isStringPresent(u.AnnualIncome) {
		balkProfileItems = append(balkProfileItems, "annual_income")
	}
	if !isStringPresent(u.Height) {
		balkProfileItems = append(balkProfileItems, "height")
	}
	if !isStringPresent(u.BodyBuild) {
		balkProfileItems = append(balkProfileItems, "body_build")
	}
	if !isStringPresent(u.MaritalStatus) {
		balkProfileItems = append(balkProfileItems, "marital_status")
	}
	if !isStringPresent(u.Child) {
		balkProfileItems = append(balkProfileItems, "child")
	}
	if !isStringPresent(u.WhenMarry) {
		balkProfileItems = append(balkProfileItems, "when_marry")
	}
	if !isStringPresent(u.WantChild) {
		balkProfileItems = append(balkProfileItems, "want_child")
	}
	if !isStringPresent(u.Smoking) {
		balkProfileItems = append(balkProfileItems, "smoking")
	}
	if !isStringPresent(u.Drinking) {
		balkProfileItems = append(balkProfileItems, "drinking")
	}
	if !isStringPresent(u.Holiday) {
		balkProfileItems = append(balkProfileItems, "holiday")
	}
	if !isStringPresent(u.HowToMeet) {
		balkProfileItems = append(balkProfileItems, "how_to_meet")
	}
	if !isStringPresent(u.CostOfDate) {
		balkProfileItems = append(balkProfileItems, "cost_of_date")
	}
	if !isStringPresent(u.Hobby) {
		balkProfileItems = append(balkProfileItems, "hobby")
	}
	if !isStringPresent(u.NthChild) {
		balkProfileItems = append(balkProfileItems, "nth_child")
	}
	if !isStringPresent(u.Housework) {
		balkProfileItems = append(balkProfileItems, "housework")
	}

	return balkProfileItems
}

// stringがnil or ""かを判定する
func isStringPresent(str string) bool {
	// このままだと " "こういう文字列も入力済みとして判定される
	// 判定できる関数が標準パッケージにありそう
	if len(str) == 0 {
		return false
	}

	return true
}
