package domain

type Vote struct {
	Option         int    `json:"option"`
	ParedaoID      string `json:"paredaoId"`
	RecaptchaToken string `json:"recaptchaToken"`
}
