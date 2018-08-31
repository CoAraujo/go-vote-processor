package domain

type Paredao struct {
	ID          string `json:"id"`
	EndTime     string `json:"endTime"`
	CreatedDate string `json:"createdDate"`
	Options     []int  `json:"options"`
}
