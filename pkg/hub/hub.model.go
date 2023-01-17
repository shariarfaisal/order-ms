package hub

type Hub struct {
	ID        uint    `json:"id" gorm:"primarykey;index"`
	Name      string  `json:"name"`
	Slug      string  `json:"slug"`
	City      string  `json:"city"`
	Area      string  `json:"area"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
