package models

type Invoice struct {
	Id     string
	Client string
	Price  int64
	Status string
	// CreateAt time.Time
	// UpdateAt time.Time
}
