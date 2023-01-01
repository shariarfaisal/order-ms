package utils

// func UrlToQuery(c *gin.Context) string {

// 	query := Query{}

// 	types := reflect.TypeOf(query)
// 	values := reflect.ValueOf(query)

// 	for i := 0; i < types.NumField(); i++ {
// 		field := types.Field(i)
// 		name := field.Tag.Get("json")
// 		kind := field.Type.Kind()

// 		v := c.Query(name)
// 		if v != "" {

// 		}
// 	}

// 	return ""
// }