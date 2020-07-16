package dislimit

import (
	"fmt"
	"github.com/alicebob/miniredis"
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v5"
	"net/http"
	"net/http/httptest"
	"time"
)

func Example() {
	run, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{
		Addr: run.Addr(),
	})

	limiter := &DisLimiter{
		Client: client,
	}

	req, _ := http.NewRequest("GET", "/iplimit", nil)

	router := gin.New()

	router.GET("/iplimit", limiter.IpLimit(5, 5), func(c *gin.Context) {
		c.String(200, "ok")
	})

	for i := 0; i < 6; i++ {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if i >= 5 {
			fmt.Println(w.Body.String())
		}
	}

	time.Sleep(1 * time.Second)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	fmt.Println(w.Code)

	// Output:
	// {"resp_common":{"ret":5001002,"msg":"server busy"}}
	// 200
}
