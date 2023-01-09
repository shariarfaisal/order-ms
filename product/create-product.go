package product

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/brand"
	"github.com/shariarfaisal/order-ms/utils"
	"github.com/shariarfaisal/validator"
	"gorm.io/gorm"
)

type VariantItemSchema struct {
	Name string `json:"name" v:"required;min=3;max=50"`
	Images []string `json:"images"`
	Details string `json:"details" v:"max=1000"`
	Price float32 `json:"price" v:"required;min=0"`
	UseInventory bool `json:"useInventory" v:"required"`
	InventoryType InventoryType `json:"inventoryType" v:"enum=periodic,perpetual"`
	Variants []VariantSchema `json:"variants"`
}

type VariantSchema struct {
	Title string `json:"title" v:"required;min=2;max=50"`
	MinSelect int `json:"minSelect" v:"min=0"`
	MaxSelect int `json:"maxSelect" v:"min=0"`
	Items []VariantItemSchema `json:"items" v:"min=1"`
}


type ProductSchema struct {
	CategoryId uint `json:"categoryId" title:"Category" v:"required"`
	Name string `json:"name" v:"required;min=3;max=50"`
	Images []string `json:"images"`
	Details string `json:"details" v:"max=1000"`
	Price float32 `json:"price"`
	BrandId uint `json:"brandId" v:"required"`
	UseInventory bool `json:"useInventory" v:"required"`
	InventoryType InventoryType `json:"inventoryType" v:"enum=periodic,perpetual"`
	Variants []VariantSchema `json:"variants"`
}


func createVariantItem(param VariantItemSchema, variantId uint, brandId uint, tx *gorm.DB) error {
	fmt.Println("Variant item created")
	
	product := Product{
		Name: param.Name,
		Details: param.Details,
		Price: param.Price,
		BrandId: brandId,
		UseInventory: param.UseInventory,
		InventoryType: param.InventoryType,
		Type: ProductTypeVariant,
		Status: "active",
		Stock: 0,
		Images: []string{},
	}

	if param.Images != nil {
		for _, img := range param.Images {
			product.Images = append(product.Images, img)
		}
	}

	if len(param.Variants) > 0 {
		product.Type = ProductTypeVariant
	}else {
		product.Type = ProductTypeSingle
	}

	// create product
	er := tx.Create(&product).Error
	if er != nil {
		return er
	}

	if param.Variants != nil {
		er = createVariant(param.Variants, product, tx)
		if er != nil {
			return er
		}
	}

	return nil
}


func createVariant(variants []VariantSchema, product Product, tx *gorm.DB) error {
	
	for _, v := range variants {
		variant := ProductVariant{
			ProductId: product.ID,
			Title: v.Title,
			MinSelect: v.MinSelect,
			MaxSelect: v.MaxSelect,
		}

		// create variant
		er := tx.Create(&variant).Error
		fmt.Println("Variant created")
		if er != nil {
			return er 
		} 

		if len(v.Items) > 0 {
			for _, item := range v.Items {
				er = createVariantItem(item, variant.ID, product.BrandId, tx)
				if er != nil {
					return er
				}
			}
		}
	}

	return nil
}


func createProduct(c *gin.Context) {
	var params ProductSchema
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	price := validator.NewField("price", reflect.ValueOf(10), "min=20", "")
	erx := price.Validate()

	fmt.Println(erx)

	isValid, err := validator.Validate(params)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if len(params.Variants) > 0 && params.Price > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrType{
			"price": "Price not acceptable",
		}})
		return 
	}

	// check brand exists
	_, er := brand.GetBrandById(params.BrandId)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brand not found"})
		return
	}

	// check category exists in brand
	cat, er := brand.GetCategoryById(params.CategoryId)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}

	if cat.BrandId != params.BrandId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}

	// check same name, same brand exists
	_, er = IsSameProductExists(params.Name, params.BrandId)
	if er == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product with same name already exists"})
		return
	}
	// create product
	data := Product{
		Name: params.Name,
		Slug: utils.GetSlug(params.Name+"-"+"-"+strconv.Itoa(int(params.BrandId))),
		Details: params.Details,
		Price: params.Price,
		BrandId: params.BrandId,
		CategoryId: params.CategoryId,
		UseInventory: params.UseInventory,
		InventoryType: params.InventoryType,
		Images: []string{},
		Status: "active",
		Stock: 0,
	}

	er = db.Transaction(func (tx *gorm.DB) error {

		if params.Images != nil {
			for _, img := range params.Images {
				data.Images = append(data.Images, img)
			}
		}

		if len(params.Variants) > 0 {
			data.Type = ProductTypeVariant
		}else {
			data.Type = ProductTypeSingle
		}

		// create product
		er := tx.Create(&data).Error
		if er != nil {
			return er 
		}

		// Create variants 
		fmt.Println("=============variant total", len(params.Variants))
		if len(params.Variants) > 0 {
			er = createVariant(params.Variants, data, tx)
			if er != nil {
				return er
			}
		}

		return nil 
	})

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": er.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"data": data })
}