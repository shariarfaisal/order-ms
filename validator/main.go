package validator

import (
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Date string's validation
func IsValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	if err == nil {
		return true
	}
	_, err = time.Parse("2006/01/02", date)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006-01-02 15:04:05", date)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006-01-02T15:04:05", date)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006-01-02T15:04:05Z", date)
	if err == nil {
		return true
	}

	_, err = time.Parse("2006-01-02T15:04:05.000Z", date)
	if err == nil {
		return true
	}

	return false
}

// Email validation 
func IsValidEmail(v string) bool {
	_, err := mail.ParseAddress(v)
	return err == nil
}

// URL validation
func IsValidURL(v string) bool {
	_, err := url.ParseRequestURI(v)
	return err == nil
}

// IP validation
func IsValidIP(v string) bool {
	return net.ParseIP(v) != nil
}

// IPv4 validation
func IsValidIpV4(v string) bool {
	ip := net.ParseIP(v)
	return ip != nil && ip.To4() != nil
}


// Validation messages
var messages = map[string] string{
	"required": "%s is required.",
	"email": "Invalid email.",
	"date": "Invalid date.",
	"min": "%s must be at least %s.",
	"max": "%s must be at most %s.",
	"enum": "%s must be one of %s.",
	"include": "%s must include one of %s.",
	"eq": "%s must be equal to %s.",
	"ne": "%s must not be equal to %s.",
	"gt": "%s must be greater than %s.",
	"gte": "%s must be greater than or equal to %s.",
	"lt": "%s must be less than %s.",
	"lte": "%s must be less than or equal to %s.",
	"url": "Invalid URL.",
	"ip": "Invalid IP address.",
	"ipv4": "It must be a valid IPv4 address.",
}

type params struct {
	v reflect.Value 
	l string 
	name string 
	errPrefix string 
}

func validateField (p params) string {
	name := p.name
	v := p.v 
	l := p.l 
	ep := p.errPrefix


	fields := make([]string, 2)
	keyValue := strings.Split(l, "=")

	if len(keyValue) == 2 {
		fields[0] = keyValue[0]
		fields[1] = keyValue[1]
	} else {
		fields[0] = keyValue[0]
	}
	key := fields[0]
	value := fields[1]

	if key != "" {
		switch key {
			case "required":
				err := fmt.Sprintf(messages["required"], name)
				if v.Kind() == reflect.String {
					if v.String() == "" {
						return err
					}
				}else if v.Kind() == reflect.Int {
					if v.Int() == 0 {
						return err
					}
				}else if v.Kind() == reflect.Float32 {
					if v.Float() == 0 {
						return err
					}
				} else if v.Kind() == reflect.Float64 {
					if v.Float() == 0 {
						return err
					}
				}else if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
					if v.Len() == 0 {
						return err
					}
				}
			case "include": 
				err := fmt.Sprintf(messages["include"], name, value)
				if ep != "" {
					err = ep + err
				}
				if v.Kind() == reflect.String {
					if !strings.Contains(value, v.String()) {
						return err
					}
				}
			case "min":
				if value != "" {
					err := fmt.Sprintf(messages["min"], name, value)
					if ep != "" {
						err = ep + err
					}
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.String {
						if vv > 0 {
							if len(v.String()) < vv{
								return err
							}
						}
					}else if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() < int64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() < float64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() < float64(vv){
								return err
							}
						}
					}else if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
						if vv > 0 {
							if v.Len() < vv{
								return err
							}
						}
					}
				}
			case "max":
				if value != "" {
					err := fmt.Sprintf(messages["max"], name, value)
					if ep != "" {
						err = ep + err
					}
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.String {
						if vv > 0 {
							if len(v.String()) > vv{
								return err
							}
						}
					}else if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() > int64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() > float64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() > float64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
						if vv > 0 {
							if v.Len() > vv{
								return err
							}
						}
					}
				}
			case "eq": 
				if value != "" {
					err := fmt.Sprintf(messages["eq"], name, value)
					if ep != "" {
						err = ep + err
					}
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() != int64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() != float64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() != float64(vv){
								return err
							}
						}
					}else if v.Kind() == reflect.String {
						if vv > 0 {
							if len(v.String()) != vv{
								return err
							} else {
								if v.String() != value{
									return err
								}
							}
						}
					}
				}
			case "ne":
				if value != "" {
					err := fmt.Sprintf(messages["ne"], name, value)
					if ep != "" {
						err = ep + err
					}

					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() == int64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() == float64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() == float64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.String {
						if vv > 0 {
							if len(v.String()) == vv{
								return err
							} else {
								if v.String() == value{
									return err
								}
							}
						}
					}
				}
			case "gt":
				if value != "" {
					err := fmt.Sprintf(messages["gt"], name, value)
					if ep != "" {
						err = ep + err
					}
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() <= int64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() <= float64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() <= float64(vv){
								return err
							}
						}
					}
				}
			case "gte":
				if value != "" {
					err := fmt.Sprintf(messages["gte"], name, value)
					if ep != "" {
						err = ep + err
					}
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() < int64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() < float64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() < float64(vv){
								return err
							}
						}
					}
				}
			case "lt":
				if value != "" {
					err := fmt.Sprintf(messages["lt"], name, value)
					if ep != "" {
						err = ep + err
					}
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() >= int64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() >= float64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() >= float64(vv){
								return err
							}
						}
					}
				}
			case "lte":
				if value != "" {
					err := fmt.Sprintf(messages["lte"], name, value)
					if ep != "" {
						err = ep + err
					}
					vv, _ := strconv.Atoi(value)
					if v.Kind() == reflect.Int {
						if vv > 0 {
							if v.Int() > int64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float32 {
						if vv > 0 {
							if v.Float() > float64(vv){
								return err
							}
						}
					} else if v.Kind() == reflect.Float64 {
						if vv > 0 {
							if v.Float() > float64(vv){
								return err
							}
						}
					}
				}
			case "enum":
				if value != "" {
					err := fmt.Sprintf(messages["enum"], name, value)
					if ep != "" {
						err = ep + err
					}
					if v.Kind() == reflect.String {
						if is := strings.Contains(value, v.String()); !is {
							return err
						}
					}
				}
			case "email":
				if v.Kind() == reflect.String {
					if !IsValidEmail(v.String()) {
						return messages["email"]
					}
				}
			case "url":
				if v.Kind() == reflect.String {
					if !IsValidURL(v.String()) {
						return messages["url"]
					}
				}
			case "ip":
				if v.Kind() == reflect.String {
					if !IsValidIP(v.String()) {
						return messages["ip"]
					}
				}
			case "ipv4":
				if v.Kind() == reflect.String {
					if !IsValidIpV4(v.String()) {
						return messages["ipv4"]
					}
				}
			case "date":
				if v.Kind() == reflect.String {
					if !IsValidDate(v.String()) {
						return messages["date"]
					}
				}
		}
	}

	return ""
}

func validateStruct(uv reflect.Value) (bool, map[string]interface{}) {
	ut := uv.Type()

	errors := map[string]interface{}{}

	for i := 0; i < ut.NumField(); i++ {
		field := ut.Field(i)

		title := field.Tag.Get("title")
		name := field.Tag.Get("json")

		if name == "" {
			name = field.Name
		}

		if title == "" {
			title = name
		}
		
		tag := field.Tag.Get("v")
		v := uv.Field(i)

		tags := strings.Split(tag, ";")

		for _, l := range tags {
			if v.Kind() == reflect.Struct {
				if isErr, err := validateStruct(v); isErr {
					errors[name] = err
					break
				}
			}else if v.Kind() == reflect.Slice {
				err := validateField(params{v: v, l: l, name: title,},)
				if err != "" {
					errors[name] = err
					break
				}else if v.Len() > 0 {
					errors[name] = map[int]interface{}{}
					for i := 0; i < v.Len(); i++ {
						if v.Index(i).Kind() == reflect.Struct {
							isErr, err := validateStruct(v.Index(i))
							if isErr {
								errors[name].(map[int]interface{})[i] = err
							}
						}else if strings.HasPrefix(l, "item.") {
							err := validateField(params{
								v: v.Index(i), l: strings.TrimPrefix(l, "item."), name: title, errPrefix: "Items of ",
							},)
							if err != "" {
								errors[name] = err 
								break
							}
						}
					}

					if len(errors[name].(map[int]interface{})) == 0 {
						delete(errors, name)
					}
				}
			}else if v.Kind() == reflect.Map {
				err := validateField(params{v: v, l: l, name: title,},)
				if err != "" {
					errors[name] = err
					break
				}else if v.Len() > 0 {
					errors[name] = map[string]interface{}{}
					for _, e := range v.MapKeys() {
						v := v.MapIndex(e)
						if v.Kind() == reflect.Struct {
							isErr, err := validateStruct(v)
							if isErr && e.Kind() == reflect.String {
								errors[name].(map[string]interface{})[e.String()] = err
							}else if isErr && e.Kind() == reflect.Int {
								errors[name].(map[string]interface{})[strconv.Itoa(int(e.Int()))] = err
							}
						}else if strings.HasPrefix(l, "item.") {
							err := validateField(params{
								v: v, l: strings.TrimPrefix(l, "item."), name: "It",
							},)
							if err != "" && e.Kind() == reflect.String {
								errors[name].(map[string]interface{})[e.String()] = err 
							} else if err != "" && e.Kind() == reflect.Int {
								errors[name].(map[string]interface{})[strconv.Itoa(int(e.Int()))] = err 
							}
						}
					}

					if len(errors[name].(map[string]interface{})) == 0 {
						delete(errors, name)
					}
				}
			}else {
				err := validateField(
					params{
						v: v,
						l: l,
						name: title,
					},
				)
				if err != "" {
					errors[name] = err
					break
				}
			}
		}
	}

	isErr := false 
	if len(errors) > 0 {
		isErr = true
	}

	return isErr, errors
}

/*
	Validate struct
	@return bool - true if no error
	@return map[string]interface{} - errors
*/
func Validate(s interface{}) (bool, map[string]interface{}) {
	ut := reflect.TypeOf(s)
	uv := reflect.ValueOf(s)
	
	if ut.Kind() == reflect.Struct {
		isErr, err := validateStruct(uv)
		return !isErr, err
	}

	return true, map[string]interface{}{}
}