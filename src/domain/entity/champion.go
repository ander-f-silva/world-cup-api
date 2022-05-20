package entity

type Champion struct {
	ID      uint   `json:"id,omitempty" gorm:"primaryKey"`
	Country string `json:"country" binding:"required"`
	Year    int64  `json:"year" binding:"required"`
}
