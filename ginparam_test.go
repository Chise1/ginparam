package ginparam

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

type Address struct {
	City    *string `gp:"json"`
	Country *string
}
type Mobile struct {
	AreaCode *string
	Number   *string
}
type User struct {
	ID       *int64 `gp:"name=id;form;"`
	name     *string
	*Address `gp:"super"`
	*Mobile  //字段不与父类同级
}

//func GetUser()
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
