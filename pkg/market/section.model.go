package market

import (
	"time"

	"github.com/shariarfaisal/order-ms/pkg/utils"
	"gorm.io/gorm"
)

type SectionImage struct {
	Url      string `json:"url"`
	Platform string `json:"platform"`
}

/*
* Video, Image Story
* Brand Slide Section
* Product Slide Section
* Category Slide Section
* Image Slide Section
 */

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
	CreatedAt   time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Title       string         `json:"title"`
	ShowTitle   bool           `json:"showTitle"`
	Slug        string         `json:"slug"`
	IsActive    bool           `json:"isActive"`
	Platform    utils.Platform `json:"platform"`
	ContentType ContentType    `json:"contentType"`
	Ratio       string         `json:"ratio"` // 16:9, 4:3, 1:1
	Validity    time.Time      `json:"validity" gorm:"type:timestamptz;"`
	// TODO: Zone
}

type SectionItem struct {
	ID              uint                   `json:"id"`
	Platform        utils.Platform         `json:"platform"`
	SectionID       uint                   `json:"sectionId"`
	Section         Section                `json:"section" gorm:"foreignKey:SectionID"`
	ReferenceId     uint                   `json:"referenceId"`
	Images          map[string]interface{} `json:"image" gorm:"type:jsonb;"`
	RedirectLink    string                 `json:"redirectLink"`
	VisitorCount    int                    `json:"visitorCount"`
	OrderCount      int                    `json:"orderCount"`
	StartDate       time.Time              `json:"startDate" gorm:"type:timestamptz;"`
	EndDate         time.Time              `json:"endDate" gorm:"type:timestamptz;"`
	ActiveDays      []int                  `json:"activeDays" gorm:"type:integer[]"`
	ActiveHours     []int                  `json:"activeHours" gorm:"type:integer[]"`
	SortingPosition int                    `json:"sortingPosition"`
	IsActive        bool                   `json:"isActive"`
}
