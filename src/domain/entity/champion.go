package entity

type Champion struct {
	ID      uint   `json:"id,omitempty" gorm:"primaryKey"`
	Country string `json:"country" binding:"required"`
	Year    int    `json:"year" binding:"required"`
}

func (champion Champion) IsExist() bool {
	return champion.ID > 0
}
