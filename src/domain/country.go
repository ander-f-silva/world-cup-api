package domain

type Country struct {
	ID   uint   `json:"id,omitempty" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
}
