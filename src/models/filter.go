package models

import (
	"regexp"
	"strconv"

	"gorm.io/gorm"
)

const (
	TYPE_EQUAL        = "="
	TYPE_LIKE         = "Like"
	TYPE_SMALLER_THAN = "<"
	TYPE_BIGGER_THAN  = ">"
)

func getTypes() []string {
	return []string{TYPE_EQUAL, TYPE_SMALLER_THAN, TYPE_BIGGER_THAN, TYPE_LIKE}
}

type Filter struct {
	gorm.Model
	Type string
}

func (f *Filter) Check(scheme string, value string) bool {
	switch f.Type {
	case TYPE_EQUAL:
		return scheme == value
	case TYPE_LIKE:
		regex := regexp.MustCompile(scheme)
		return regex.MatchString(value)
	case TYPE_SMALLER_THAN:
		scheme_float, err := strconv.ParseFloat(scheme, 64)
		value_float, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return false
		}
		return scheme_float > value_float
	case TYPE_BIGGER_THAN:
		scheme_float, err := strconv.ParseFloat(scheme, 64)
		value_float, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return false
		}
		return scheme_float < value_float
	}
	return false
}

func InitFilter(db *gorm.DB) {
	types := getTypes()
	for _, t := range types {
		db.Create(&Filter{Type: t})
	}
}
