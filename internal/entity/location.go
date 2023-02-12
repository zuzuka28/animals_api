package entity

type Location struct {
	Id        int64   `json:"id"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
