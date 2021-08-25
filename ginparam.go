package ginparam

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

func CheckParam(obj interface{}, g *gin.Context) {
	r := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj)
	for i := 0; i < r.NumField(); i++ {
		gp := r.Field(i).Get("gp")
		if gp == "" {
			g.PostForm()
			v.FieldByName(r.Field(i).Name).Set()
		}
	}
}
