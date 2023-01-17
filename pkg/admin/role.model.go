package admin

type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name" gorm:"uniqueIndex;not null"`
}
