package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

type In map[string][]interface{}
type Compare map[string]interface{}


func eq(q Compare, separator string, logicalOp string, ex bool, rest ...interface{}) (string, bool) {
	s := ""
	pre := "'"
	suf := "'"


	for i, v := range rest {
		if i == 0 {
			pre = v.(string)
		} else if i == 1 {
			suf = v.(string)
		}
	}

	l := len(q)
	if q != nil && l > 0 {
		if !ex {
			s += " WHERE "
		}else {
			s += fmt.Sprintf(" %s ", separator)
		}

		i := 0
		for k, v := range q {
			kind := reflect.TypeOf(v).Kind()
			if kind == reflect.String {
				v = fmt.Sprintf("%s%s%s", pre, v.(string), suf)
			} else if kind == reflect.Int {
				v = strconv.Itoa(v.(int))
			} else if kind == reflect.Bool {
				v = strconv.FormatBool(v.(bool))
			} else if kind == reflect.Float64 {
				v = strconv.FormatFloat(v.(float64), 'f', -1, 64)
			} else if kind == reflect.Float32 {
				v = strconv.FormatFloat(v.(float64), 'f', -1, 32)
			}
			
			s += fmt.Sprintf("LOWER(%s) %s LOWER(%s) ", k, logicalOp, v.(string))

			if i < l - 1 {
				s += fmt.Sprintf(" %s ", separator)
			}
			i++
		}
	}

	if s != "" {
		ex = true
	}

	return s, ex
}

func in(q In, separator string, logicalOp string, ex bool) (string, bool) {
	s := ""
	
	l := len(q)
	if q != nil && l > 0 {
		if !ex {
			s += " WHERE "
		}else {
			s += fmt.Sprintf(" %s ", separator)
		}

		i := 0
		for k, v := range q {
			kind := reflect.TypeOf(v).Kind()
			
			arr := ""
			if kind == reflect.Slice {
				arr = "("
				for i, value := range v {
					kind := reflect.TypeOf(value).Kind()
					if kind == reflect.String {
						arr += fmt.Sprintf("'%s'", value.(string))
					} else if kind == reflect.Int {
						arr += strconv.Itoa(value.(int))
					}

					if i < len(v) - 1 {
						arr += ", "
					}
				}
				arr += ")"
			}

			s += fmt.Sprintf("%s %s %s ", k, logicalOp, arr)

			if i < l - 1 {
				s += fmt.Sprintf(" %s ", separator)
			}
			i++
		}
	}

	if s != "" {
		ex = true
	}

	return s, ex
}

type Where struct {
	Equal     Compare `json:"equal"`
	EqualOr   Compare `json:"equalOr"`
	NotEqual  Compare `json:"notEqual"`
	NotEqualOr  Compare `json:"notEqualOr"`
	Gt Compare `json:"greaterThan"`
	GtOr Compare `json:"greaterThanOr"`
	Lt  Compare `json:"lessThan"`
	LtOr  Compare `json:"lessThanOr"`
	Gte Compare `json:"greaterThanOrEqual"`
	GteOr Compare `json:"greaterThanOrEqualOr"`
	Lte Compare `json:"lessThanOrEqual"`
	LteOr Compare `json:"lessThanOrEqualOr"`
	StartsWith Compare `json:"startsWith"`
	StartsWithOr Compare `json:"startsWithOr"`
	EndsWith   Compare `json:"endsWith"`
	EndsWithOr   Compare `json:"endsWithOr"`
	Like      Compare `json:"like"`
	LikeOr      Compare `json:"likeOr"`
	In        In `json:"in"`
	NotIn     In `json:"notIn"`
	IsNull    Compare `json:"isNull"`
	IsNullOr  Compare `json:"isNullOr"`
	IsNotNull Compare `json:"isNotNull"`
	IsNotNullOr Compare `json:"isNotNullOr"`
}

type Query struct {
	Select    []string               `json:"select"`
	Table     string                 `json:"table"`
	Where Where `json:"where"`
	GroupBy   string               `json:"groupBy"`
	Limit     int                    `json:"limit"`
	Offset    int                    `json:"offset"`
	Order     string                 `json:"order"`
	OrderBy   string                 `json:"orderBy"`
}


func (q *Where) whereClause() string {
	s := ""

	ex := false

	eqS, eqEx := eq(q.Equal, "AND", "=", ex)
	s += eqS
	ex = eqEx

	eqOrS, eqOrEx := eq(q.EqualOr, "OR", "=", ex)
	s += eqOrS
	ex = eqOrEx

	notEqS, notEqEx := eq(q.NotEqual, "AND", "!=", ex)
	s += notEqS
	ex = notEqEx

	notEqOrS, notEqOrEx := eq(q.NotEqualOr, "OR", "!=", ex)
	s += notEqOrS
	ex = notEqOrEx

	gtS, gtEx := eq(q.Gt, "AND", ">", ex)
	s += gtS
	ex = gtEx

	gtOrS, gtOrEx := eq(q.GtOr, "OR", ">", ex)
	s += gtOrS
	ex = gtOrEx

	ltS, ltEx := eq(q.Lt, "AND", "<", ex)
	s += ltS
	ex = ltEx

	ltOrS, ltOrEx := eq(q.LtOr, "OR", "<", ex)
	s += ltOrS
	ex = ltOrEx

	gteS, gteEx := eq(q.Gte, "AND", ">=", ex)
	s += gteS
	ex = gteEx

	gteOrS, gteOrEx := eq(q.GteOr, "OR", ">=", ex)
	s += gteOrS
	ex = gteOrEx

	lteS, lteEx := eq(q.Lte, "AND", "<=", ex)
	s += lteS
	ex = lteEx

	startsWithS, startsWithEx := eq(q.StartsWith, "AND", "LIKE", ex, "'", "%'")
	s += startsWithS
	ex = startsWithEx

	startsWithOrS, startsWithOrEx := eq(q.StartsWithOr, "OR", "LIKE", ex, "'", "%'")
	s += startsWithOrS
	ex = startsWithOrEx

	endsWithS, endsWithEx := eq(q.EndsWith, "AND", "LIKE", ex, "'%", "'")
	s += endsWithS
	ex = endsWithEx

	endsWithOrS, endsWithOrEx := eq(q.EndsWithOr, "OR", "LIKE", ex, "'%", "'")
	s += endsWithOrS
	ex = endsWithOrEx

	likeS, likeEx := eq(q.Like, "AND", "LIKE", ex, "'%", "%'")
	s += likeS
	ex = likeEx

	likeOrS, likeOrEx := eq(q.LikeOr, "OR", "LIKE", ex, "'%", "%'")
	s += likeOrS
	ex = likeOrEx

	inS, inEx := in(q.In, "AND", "IN", ex)
	s += inS
	ex = inEx

	notInS, notInEx := in(q.NotIn, "AND", "NOT IN", ex)
	s += notInS
	ex = notInEx

	return s
}

func (q *Query) sql() string {
	s := "SELECT "

	if q.Select != nil {
		for i, v := range q.Select {
			s += v 
			if i < len(q.Select)-1 {
				s += ", "
			}
		}
	} else {
		s += "* "
	}

	s += " FROM " + q.Table + " "

	w := q.Where.whereClause()
	s += w

	if q.GroupBy != "" {
		s += " GROUP BY " + q.GroupBy + " "
	}

	if q.Order != "" {
		s += " ORDER BY " + q.OrderBy 
	}

	if q.OrderBy != "" {
		s += " " + q.Order + " "
	}else {
		s += " ASC "
	}

	if q.Limit != 0 {
		s += " LIMIT " + strconv.Itoa(q.Limit) + " "
	}

	if q.Offset != 0 {
		s += " OFFSET " + strconv.Itoa(q.Offset) + " "
	}

	return s+";"
}
