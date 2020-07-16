package dislimit

import (
	"aix.zhhainiao.com/aixpublic/server/core/apiutil/reqresp"
	"bytes"
	"encoding/json"
	"github.com/alicebob/miniredis"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/redis.v5"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestDisLimiter_IpLimit(t *testing.T) {
	uid := "1234567"
	path := "/limit"
	id := uid + path

	run, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{
		Addr: run.Addr(),
	})

	tokenKey, timestampKey := getKeys(id)

	for i := 0; i < 10; i++ {

		run := luaScript.Run(client, []string{tokenKey, timestampKey}, 5, 8, time.Now().Unix(), 1)
		if run.Err() != nil {
			t.Fail()
		}

		result := run.Val().([]interface{})
		if result[0].(int64) == 0 {
			assert.Equal(t, 8, i)
			time.Sleep(1 * time.Second)
		}
	}

	time.Sleep(1 * time.Second)
	for i := 0; i < 5; i++ {
		run := luaScript.Run(client, []string{tokenKey, timestampKey}, 5, 5, time.Now().Unix(), 1)
		if run.Err() != nil {
			t.Fail()
		}

		result := run.Val().([]interface{})
		if result[0].(int64) == 0 {
			t.Fail()
		}
	}
}

func BenchmarkDisLimiter_IpLimit(b *testing.B) {
	uid := "1234567"
	path := "/ip"
	id := uid + path

	run, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{
		Addr: run.Addr(),
	})

	tokenKey, timestampKey := getKeys(id)

	for i := 0; i < b.N; i++ {
		_ = luaScript.Run(client, []string{tokenKey, timestampKey}, 5, 8, time.Now().Unix(), 2)
	}
}

func TestDisLimiter_PathLimit(t *testing.T) {
	run, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{
		Addr: run.Addr(),
	})

	limiter := NewLimiter(client)

	req, _ := http.NewRequest("GET", "/pathlimit", nil)

	router := gin.New()

	router.GET("/pathlimit", limiter.PathLimit(5, 5), func(c *gin.Context) {
		c.String(200, "ok")
	})

	for i := 0; i < 10; i++ {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if i < 5 {
			assert.Equal(t, w.Code, 200)
		} else {
			assert.Equal(t, w.Body.String(), "{\"resp_common\":{\"ret\":5001002,\"msg\":\"server busy\"}}")
		}
	}

	time.Sleep(1 * time.Second)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)

}

func TestDisLimiter_UserLimit(t *testing.T) {
	run, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{
		Addr: run.Addr(),
	})

	limiter := NewLimiter(client)

	router := gin.New()

	router.POST("/userlimit", limiter.UserPathLimit(5, 5), func(c *gin.Context) {
		c.String(200, "ok")
	})

	for i := 0; i < 10; i++ {
		body := &reqresp.Request{
			Common: reqresp.RequestCommon{
				Uid: "uid" + strconv.Itoa(i),
			},
		}

		marshal, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/userlimit", bytes.NewReader([]byte(marshal)))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, 200)
	}

	for i := 0; i < 10; i++ {
		body := &reqresp.Request{
			Common: reqresp.RequestCommon{
				Uid: "uid",
			},
		}

		marshal, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/userlimit", bytes.NewReader([]byte(marshal)))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if i < 5 {
			assert.Equal(t, w.Code, 200)
		} else {
			assert.Equal(t, w.Body.String(), "{\"resp_common\":{\"ret\":5001002,\"msg\":\"server busy\"}}")
		}
	}
}

func TestDisLimiter_IpLimit2(t *testing.T) {
	run, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{
		Addr: run.Addr(),
	})

	limiter := NewLimiter(client)

	req, _ := http.NewRequest("GET", "/iplimit", nil)

	router := gin.New()

	router.GET("/iplimit", limiter.IpLimit(5, 5), func(c *gin.Context) {
		c.String(200, "ok")
	})

	for i := 0; i < 10; i++ {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if i < 5 {
			assert.Equal(t, w.Code, 200)
		} else {
			assert.Equal(t, w.Body.String(), "{\"resp_common\":{\"ret\":5001002,\"msg\":\"server busy\"}}")
		}
	}

	time.Sleep(1 * time.Second)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
}
