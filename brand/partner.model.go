package brand

type Partner struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Brands []Brand       `json:"brands" gorm:"foreignKey:PartnerId"`
	Users  []PartnerUser `json:"users" gorm:"foreignKey:PartnerId"`
}

type PartnerRole string

const (
	PartnerRoleAdmin      PartnerRole = "admin"
	PartnerRoleManager    PartnerRole = "manager"
	PartnerRoleOperations PartnerRole = "operations"
	PartnerRoleReporting  PartnerRole = "reporting"
	PartnerRoleSupport    PartnerRole = "support"
	PartnerRoleFinance    PartnerRole = "finance"
)

type PartnerUserStatus string

const (
	PartnerUserActive   PartnerUserStatus = "active"
	PartnerUserInactive PartnerUserStatus = "inactive"
	PartnerUserBlocked  PartnerUserStatus = "blocked"
)

type PartnerUser struct {
	ID            int               `json:"id"`
	PartnerId     int               `json:"partnerId"`
	Partner       Partner           `json:"partner" gorm:"foreignKey:PartnerId"`
	Name          string            `json:"name"`
	Email         string            `json:"email"`
	EmailVerified bool              `json:"emailVerified"`
	Phone         string            `json:"phone"`
	Password      string            `json:"password"`
	Role          PartnerRole       `json:"role"`
	Status        PartnerUserStatus `json:"status"`
}
