package section

import (
	"time"

	"github.com/shariarfaisal/order-ms/pkg/utils"
	"gorm.io/gorm"
)

type SectionImage map[string]struct {
	Url      string `json:"url"`
	Platform string `json:"platform"`
}

type ContentType string

const (
	ContentTypeImage    ContentType = "image"
	ContentTypeVideo    ContentType = "video"
	ContentTypeProduct  ContentType = "product"
	ContentTypeCategory ContentType = "category"
	ContentTypeBrand    ContentType = "brand"
)

type Section struct {
	ID          uint           `json:"id"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title       string         `json:"title"`
	Subtitle    string         `json:"subtitle"`
	Slug        string         `json:"slug"`
	IsActive    bool           `json:"is_active"`
	Images      SectionImage   `json:"images"`
	Platform    utils.Platform `json:"platform"`
	ContentType ContentType    `json:"content_type"`
}

type SectionItem struct {
	ID              uint           `json:"id"`
	Platform        utils.Platform `json:"platform"`
	SectionID       uint           `json:"section_id"`
	Section         Section        `json:"section" gorm:"foreignKey:SectionID"`
	ReferenceId     uint           `json:"reference_id"`
	Image           string         `json:"image"`
	RedirectLink    string         `json:"redirect_link"`
	VisitorCount    int            `json:"visitor_count"`
	OrderCount      int            `json:"order_count"`
	StartDate       time.Time      `json:"start_date" gorm:"type:timestamptz;"`
	EndDate         time.Time      `json:"end_date" gorm:"type:timestamptz;"`
	ActiveDays      []int          `json:"active_days" gorm:"type:integer[]"`
	ActiveHours     []int          `json:"active_hours" gorm:"type:integer[]"`
	SortingPosition int            `json:"sorting_position"`
	IsActive        bool           `json:"is_active"`
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&Section{}, &SectionItem{})
}
