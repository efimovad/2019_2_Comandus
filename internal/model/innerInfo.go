package model

type InnerInfo struct {
	UserID int `json:"user_id"`
	WhoSeeProfile string `json:"who_see_profile"`
	ControlQuestion string `json:"control_question"`
	ControlAnswer string `json:"-"`
}

