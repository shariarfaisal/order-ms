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

type Field struct {
	Name string
	Value reflect.Value
	Prefix string
	TagValue string
}

func NewField(name string, value reflect.Value, tagValue string, prefix string,) *Field {
	return &Field{
		Name: name,
		Value: value,
		Prefix: prefix,
		TagValue: tagValue,
	}
}

func (f *Field) Required() (err string) {
	v := f.Value 
	name := f.Name

	err = fmt.Sprintf(messages["required"], name)
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

	return "" 
}

func (f *Field) Include() (err string) {
	v := f.Value
	name := f.Name
	tagValue := f.TagValue
	errorPrefix := f.Prefix

	err = fmt.Sprintf(messages["include"], name, tagValue)
	if errorPrefix != "" {
		err = errorPrefix + err
	}
	if v.Kind() == reflect.String {
		if !strings.Contains(tagValue, v.String()) {
			return err
		}
	}
	return ""
}

func (f *Field) Min() (err string) {
	v := f.Value
	name := f.Name
	tagValue := f.TagValue
	errorPrefix := f.Prefix

	if tagValue != "" {
		err = fmt.Sprintf(messages["min"], name, tagValue)
		if errorPrefix != "" {
			err = errorPrefix + err
		}
		tagIntValue, er := strconv.Atoi(tagValue)
		if er != nil {
			return err
		}
		if v.Kind() == reflect.String {
			if len(v.String()) < tagIntValue{
				return err
			}
		}else if v.Kind() == reflect.Int {
			if v.Int() < int64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float32 {
			if v.Float() < float64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float64 {
			if v.Float() < float64(tagIntValue){
				return err
			}
		}else if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
			if v.Len() < tagIntValue{
				return err
			}
		}
	}
	return ""
}

func (f *Field) Max() (err string) {
	v := f.Value
	name := f.Name
	tagValue := f.TagValue
	errorPrefix := f.Prefix

	if tagValue != "" {
		err = fmt.Sprintf(messages["max"], name, tagValue)
		if errorPrefix != "" {
			err = errorPrefix + err
		}
		tagIngValue, er := strconv.Atoi(tagValue)
		if er != nil {
			return err
		}
		if v.Kind() == reflect.String {
			if len(v.String()) > tagIngValue{
				return err
			}
		}else if v.Kind() == reflect.Int {
			if v.Int() > int64(tagIngValue){
				return err
			}
		} else if v.Kind() == reflect.Float32 {
			if v.Float() > float64(tagIngValue){
				return err
			}
		} else if v.Kind() == reflect.Float64 {
			if v.Float() > float64(tagIngValue){
				return err
			}
		} else if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
			if v.Len() > tagIngValue{
				return err
			}
		}
	}
	return ""
}

func (f *Field) Equal() (err string) {
	v := f.Value
	name := f.Name
	tagValue := f.TagValue
	errorPrefix := f.Prefix


	if tagValue != "" {
		err = fmt.Sprintf(messages["eq"], name, tagValue)
		if errorPrefix != "" {
			err = errorPrefix + err
		}
		tagIntValue, er := strconv.Atoi(tagValue)
		if er != nil {
			return err
		}
		if v.Kind() == reflect.Int {
			if v.Int() != int64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float32 {
			if v.Float() != float64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float64 {
			if v.Float() != float64(tagIntValue){
				return err
			}
		}else if v.Kind() == reflect.String {
			if len(v.String()) != tagIntValue{
				return err
			} else {
				if v.String() != tagValue{
					return err
				}
			}
		}
	}
	return ""
}

func (f *Field) NotEqual() (err string) {
	v := f.Value
	name := f.Name
	tagValue := f.TagValue
	errorPrefix := f.Prefix

	if tagValue != "" {
		err := fmt.Sprintf(messages["ne"], name, tagValue)
		if errorPrefix != "" {
			err = errorPrefix + err
		}
		
		tagIntValue, er := strconv.Atoi(tagValue)
		if er != nil {
			return err
		}
		if v.Kind() == reflect.Int {
			if v.Int() == int64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float32 {
			if v.Float() == float64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float64 {
			if v.Float() == float64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.String {
			if len(v.String()) == tagIntValue{
				return err
			} else {
				if v.String() == tagValue{
					return err
				}
			}
		}else if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
			if v.Len() == tagIntValue {
				return err
			}
		}
	}
	return ""
}

func (f *Field) GreaterThan() (err string) {
	v := f.Value
	name := f.Name
	tagValue := f.TagValue
	errorPrefix := f.Prefix

	if tagValue != "" {
		err := fmt.Sprintf(messages["gte"], name, tagValue)
		if errorPrefix != "" {
			err = errorPrefix + err
		}
		tagIntValue, er := strconv.Atoi(tagValue)
		if er != nil {
			return err
		}
		if v.Kind() == reflect.Int {
			if v.Int() <= int64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float32 {
			if v.Float() <= float64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float64 {
			if v.Float() <= float64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
			if v.Len() <= tagIntValue{
				return err
			}
		}
	}
	return ""
}

func (f *Field) GreaterThanOrEqual() (err string) {
	v := f.Value
	name := f.Name
	tagValue := f.TagValue
	errorPrefix := f.Prefix


	if tagValue != "" {
		err = fmt.Sprintf(messages["gt"], name, tagValue)
		if errorPrefix != "" {
			err = errorPrefix + err
		}
		tagIntValue, er := strconv.Atoi(tagValue)
		if er != nil {
			return err
		}
		if v.Kind() == reflect.Int {
			if v.Int() < int64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float32 {
			if v.Float() < float64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float64 {
			if v.Float() < float64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
			if v.Len() < tagIntValue {
				return err
			}
		}
	}
	return ""
}

func (f *Field) LessThan() (err string) {
	v := f.Value
	name := f.Name
	tagValue := f.TagValue
	errorPrefix := f.Prefix
	if tagValue != "" {
		err = fmt.Sprintf(messages["lte"], name, tagValue)
		if errorPrefix != "" {
			err = errorPrefix + err
		}
		tagIntValue, er := strconv.Atoi(tagValue)
		if er != nil {
			return err
		}
		if v.Kind() == reflect.Int {
			if v.Int() >= int64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float32 {
			if v.Float() >= float64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float64 {
			if v.Float() >= float64(tagIntValue){
				return err
			}
		}else if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
			fmt.Println(v.Len(), tagIntValue, v.Len() > tagIntValue)
			if v.Len() >= tagIntValue{
				return err
			}
		}
	}
	return ""
}

func (f *Field) LessThanOrEqual() (err string) {
	v := f.Value
	name := f.Name
	tagValue := f.TagValue
	errorPrefix := f.Prefix
	if tagValue != "" {
		err = fmt.Sprintf(messages["lt"], name, tagValue)
		if errorPrefix != "" {
			err = errorPrefix + err
		}
		tagIntValue, er := strconv.Atoi(tagValue)
		if er != nil {
			return err
		}
		if v.Kind() == reflect.Int {
			if v.Int() > int64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float32 {
			if v.Float() > float64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Float64 {
			if v.Float() > float64(tagIntValue){
				return err
			}
		} else if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
			if v.Len() > tagIntValue{
				return err
			}
		}
	}
	return ""
}

func (f *Field) Enum() (err string) {
	v := f.Value
	name := f.Name
	tagValue := f.TagValue
	errorPrefix := f.Prefix

	if tagValue != "" {
		err := fmt.Sprintf(messages["enum"], name, tagValue)
		if errorPrefix != "" {
			err = errorPrefix + err
		}
		if v.Kind() == reflect.String {
			if is := strings.Contains(tagValue, v.String()); !is {
				return err
			}
		}
	}
	return ""
}

func (f *Field) Validate() string {
	tag := f.TagValue
	tagSplit := strings.Split(tag, "=")

	var key string
	if len(tagSplit) > 1 {
		key = tagSplit[0]
		f.TagValue = tagSplit[1]
	} else {
		key = tagSplit[0]
	}

	if key != "" {
		switch key {
			case "required":
				err := f.Required()
				if err != "" {
					return err 
				}
			case "include": 
				err := f.Include()
				if err != "" {
					return err
				}
			case "min":
				err := f.Min()
				if err != "" {
					return err
				}
			case "max":
				err := f.Max()
				if err != "" {
					return err
				}
			case "eq": 
				err := f.Equal()
				if err != "" {
					return err
				}
			case "ne":
				err := f.NotEqual()
				if err != "" {
					return err
				}
			case "gte":
				err := f.GreaterThanOrEqual()
				if err != "" {
					return err
				}
			case "gt":
				err := f.GreaterThan()
				if err != "" {
					return err
				}
			case "lte":
				err := f.LessThanOrEqual()
				if err != "" {
					return err
				}
			case "lt":
				err := f.LessThan()
				if err != "" {
					return err
				}
			case "enum":
				err := f.Enum()
				if err != "" {
					return err
				}
			case "email":
				if f.Value.Kind() == reflect.String {
					if !IsValidEmail(f.Value.String()) {
						return messages["email"]
					}
				}
			case "url":
				if f.Value.Kind() == reflect.String {
					if !IsValidURL(f.Value.String()) {
						return messages["url"]
					}
				}
			case "ip":
				if f.Value.Kind() == reflect.String {
					if !IsValidIP(f.Value.String()) {
						return messages["ip"]
					}
				}
			case "ipv4":
				if f.Value.Kind() == reflect.String {
					if !IsValidIpV4(f.Value.String()) {
						return messages["ipv4"]
					}
				}
			case "date":
				if f.Value.Kind() == reflect.String {
					if !IsValidDate(f.Value.String()) {
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

		for _, tag := range tags {

			field := NewField(title, v, tag, "")
			fmt.Println(field.Value)
			if v.Kind() == reflect.Struct {
				if isErr, err := validateStruct(v); isErr {
					errors[name] = err
					break
				}
			}else if v.Kind() == reflect.Slice {
				fmt.Println("slice", v.Len(), tag, name, title)
				err := field.Validate()
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
						}else if strings.HasPrefix(tag, "item.") {
							field.Prefix = "Items of "
							err := field.Validate()
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
				err := field.Validate()
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
						}else if strings.HasPrefix(tag, "item.") {
							
							field.TagValue = strings.TrimPrefix(tag, "item.")
							field.Name = "It"
							err := field.Validate()

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
				err := field.Validate()
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
func Validate(s interface{}) (valid bool, errors map[string]interface{}) {
	ut := reflect.TypeOf(s)
	uv := reflect.ValueOf(s)

	valid = true 
	errors = map[string]interface{}{}
	
	if ut.Kind() == reflect.Struct {
		isErr, err := validateStruct(uv)
		valid = !isErr
		errors = err
	}

	return valid, errors
}