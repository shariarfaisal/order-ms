package admin

type Admin struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Phone    string `json:"phone"`
	Image    string `json:"image"`
	RoleName string `json:"role_name" gorm:"column:role_name"`
	Role     Role   `json:"-" gorm:"foreignKey:RoleName;references:Name"`
}
