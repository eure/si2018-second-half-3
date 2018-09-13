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
	if isStringPresent(u.Gender) {
		balkProfileItems = append(balkProfileItems, "Gender")
	}
	if isStringPresent(u.Nickname) {
		balkProfileItems = append(balkProfileItems, "Nickname")
	}
	if isStringPresent(u.ImageURI) {
		balkProfileItems = append(balkProfileItems, "ImageURI")
	}
	if isStringPresent(u.Tweet) {
		balkProfileItems = append(balkProfileItems, "Tweet")
	}
	if isStringPresent(u.Introduction) {
		balkProfileItems = append(balkProfileItems, "Introduction")
	}
	if isStringPresent(u.ResidenceState) {
		balkProfileItems = append(balkProfileItems, "ResidenceState")
	}
	if isStringPresent(u.HomeState) {
		balkProfileItems = append(balkProfileItems, "HomeState")
	}
	if isStringPresent(u.Education) {
		balkProfileItems = append(balkProfileItems, "Education")
	}
	if isStringPresent(u.Job) {
		balkProfileItems = append(balkProfileItems, "Job")
	}
	if isStringPresent(u.AnnualIncome) {
		balkProfileItems = append(balkProfileItems, "AnnualIncome")
	}
	if isStringPresent(u.Height) {
		balkProfileItems = append(balkProfileItems, "Height")
	}
	if isStringPresent(u.BodyBuild) {
		balkProfileItems = append(balkProfileItems, "BodyBuild")
	}
	if isStringPresent(u.MaritalStatus) {
		balkProfileItems = append(balkProfileItems, "MaritalStatus")
	}
	if isStringPresent(u.Child) {
		balkProfileItems = append(balkProfileItems, "Child")
	}
	if isStringPresent(u.WhenMarry) {
		balkProfileItems = append(balkProfileItems, "WhenMarry")
	}
	if isStringPresent(u.WantChild) {
		balkProfileItems = append(balkProfileItems, "WantChild")
	}
	if isStringPresent(u.Smoking) {
		balkProfileItems = append(balkProfileItems, "Smoking")
	}
	if isStringPresent(u.Drinking) {
		balkProfileItems = append(balkProfileItems, "Drinking")
	}
	if isStringPresent(u.Holiday) {
		balkProfileItems = append(balkProfileItems, "Holiday")
	}
	if isStringPresent(u.HowToMeet) {
		balkProfileItems = append(balkProfileItems, "HowToMeet")
	}
	if isStringPresent(u.CostOfDate) {
		balkProfileItems = append(balkProfileItems, "CostOfDate")
	}
	if isStringPresent(u.NthChild) {
		balkProfileItems = append(balkProfileItems, "NthChild")
	}
	if isStringPresent(u.Housework) {
		balkProfileItems = append(balkProfileItems, "Housework")
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
