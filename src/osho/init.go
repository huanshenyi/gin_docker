package osho

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin_docker/src/domain/user"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)

}

type HTTPParams struct {
	URL    string
	User   *user.UserData
	Body   map[string]interface{}
	Params []gin.Param
	Query  string
	Method string
}

func (hp *HTTPParams) AddParam(key string, value string) {
	hp.Params = append(hp.Params, gin.Param{Key: key, Value: value})
}

func (hp *HTTPParams) AddIntParam(key string, value int) {
	hp.AddParam(key, strconv.Itoa(value))
}

// gin.Contextを作る用
func GetGinContext(p HTTPParams) (c *gin.Context, err error) {
	body, err := json.Marshal(p.Body)
	if err != nil {
		return
	}
	c, _ = gin.CreateTestContext(httptest.NewRecorder())
	c.Params = p.Params
	c.Request, err = http.NewRequest(p.Method, fmt.Sprintf("%s?%s", p.URL, p.Query), bytes.NewReader(body))
	if err != nil {
		return
	}
	c.Request.Header.Set("Content-Type", "application/json")
	if p.User != nil {
		c.Set("user", *p.User)
	}
	return
}

// TestChecker は、エラーを考慮してEqualityをチェックします
func TestChecker(t *testing.T, expData interface{}, expError error, actual interface{}, err error) {
	t.Helper()
	if err != nil && expError != nil {
		assert.Equal(t, expError.Error(), err.Error())
	} else {
		assert.Equal(t, expError, err)
		assert.Equal(t, expData, actual)
	}
}

func SortCase(m interface{}) []string {
	rt := reflect.TypeOf(m)
	if rt.Kind() != reflect.Map {
		panic(fmt.Errorf("type should be map, got %v", rt.Kind()))
	}
	if rt.Key().Kind() != reflect.String {
		panic(fmt.Errorf("map key type should be string, got %v", rt.Key().Kind()))
	}

	var keys []string
	for _, key := range reflect.ValueOf(m).MapKeys() {
		keys = append(keys, key.String())
	}
	sort.Strings(keys)
	return keys
}
