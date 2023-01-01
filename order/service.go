package order

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/shariarfaisal/order-ms/brand"
	"github.com/shariarfaisal/order-ms/product"
)

func isValidItemsForOrder(items []OrderItemSchema, prods []product.Product) (bool, string) {
	/*
		- check products existence
		- check products status
		- check products inventory
	*/

	byId := map[int]product.Product{}
	for _, prod := range prods {
		byId[int(prod.ID)] = prod
	}

	for _, item := range items {
		prod, ok := byId[item.Id]
		if !ok || prod.Status != product.ProductStatusApproved {
			return false, fmt.Sprintf("%s not found", prod.Name)
		}

		if !prod.IsAvailable {
			return false, fmt.Sprintf("%s not available", prod.Name)
		}

		if prod.UseInventory && prod.Stock < item.Quantity {
			return false, fmt.Sprintf("%s out of stock", prod.Name)
		}
	}

	return true, ""
}

func getOrderItems(params OrderSchema) ([]product.Product, error) {
	items := params.Items
	itemsId := []int{}
	for _, item := range items {
		itemsId = append(itemsId, item.Id)
	}

	prods, err := product.GetByIds(itemsId)
	if err != nil {
		return nil, err
	}
	
	_, errMsg := isValidItemsForOrder(items, prods)
	if errMsg != "" {
		return nil, errors.New(errMsg)
	}

	return prods,nil 
}


func InBetweenHours (start /* start hour */, end /* end hour*/, hour, minute float32) bool {
	if minute > 0 {
		hour += minute / 100
	}

	if start < end {
		return start <= hour && end >= hour
	} else {
		return start <= hour || end >= hour
	}
}

func isBrandsOperating(brands []brand.Brand) (bool, string) {

	for _, item := range brands {
		if item.Status != brand.BrandStatusActive {
			return false, fmt.Sprintf("%s is not operating", item.Name)
		}

		if !item.IsAvailable {
			return false, fmt.Sprintf("%s is not operating", item.Name)
		}

		opTime := item.OperatingTimes

		if len(opTime) > 0 {
			day := time.Now().Day()
	
			times, exists := opTime[strconv.Itoa(day)]
			if !exists {
				return false, fmt.Sprintf("%s is not operating", item.Name)
			}
			
			hour := float32(time.Now().Hour())
			minute := float32(time.Now().Minute())
			isOperating := false
			for _, t := range times.([]brand.OperatingTime) {
				from := t.From.Hour + t.From.Minute / 100
				to := t.To.Hour + t.To.Minute / 100
				if InBetweenHours(float32(from), float32(to), hour, minute) {
					isOperating = true
					break
				}
			}
	
			if !isOperating {
				return false, fmt.Sprintf("%s is not operating", item.Name)
			}
		}
	}

	return true, ""
}