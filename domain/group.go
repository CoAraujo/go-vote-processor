package domain

import "time"

type Group struct {
	ID          string    `json:"id"`
	CreatedDate time.Time `json:"createdDate"`
	EndTime     time.Time `json:"endTime"`
	Options     []int     `json:"options"`
}
