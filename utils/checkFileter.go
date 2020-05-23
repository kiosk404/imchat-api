/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/5/16
**/
package utils

import (
	"reflect"
)

func CheckEmpty(args ...interface{}) (string, bool) {
	for _, arg := range args {
		t := reflect.TypeOf(arg)
		v := reflect.ValueOf(arg)
		switch t.Kind() {
		case reflect.Float32, reflect.Float64:
			if v.Interface().(float64) == 0.0 {
				return t.Name(), false
			}
		case reflect.Int16, reflect.Int32, reflect.Int64:
			if v.Interface().(int64) == 0 {
				return t.Name(), false
			}
		case reflect.String:
			if v.Interface().(string) == "" {
				return t.Name(), false
			}
		}
	}
	return "", true
}
