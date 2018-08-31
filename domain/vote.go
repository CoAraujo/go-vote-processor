package domain

type Vote struct {
	Option         int    `json:"option"`
	GroupID        string `json:"groupId"`
	RecaptchaToken string `json:"recaptchaToken"`
}
