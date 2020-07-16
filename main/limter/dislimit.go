package dislimit

import (
	"aix.zhhainiao.com/aixpublic/server/core/apiutil/reqresp"
	"aix.zhhainiao.com/aixpublic/server/core/ctxlog"
	"aix.zhhainiao.com/aixpublic/utils/code"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v5"
	"time"
)

//DisLimiter 分布式限流.
type DisLimiter struct {
	Client *redis.Client
}

//NewLimiter 新建limiter.
func NewLimiter(client *redis.Client) *DisLimiter {
	return &DisLimiter{
		Client: client,
	}
}

//PathLimit 对接口进行限流.
func (d *DisLimiter) PathLimit(rate, capacity int) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		d.doLimit(c, path, rate, capacity)

		c.Next()
	}
}

//UserPathLimit uid+接口进行限流.
func (d *DisLimiter) UserPathLimit(rate, capacity int) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &reqresp.Request{}
		if !reqresp.ParseBodyWithErrorHandle(c, req) {
			reqresp.ReturnError(c, code.BadRequest)
			c.Abort()
		}

		uid := req.Common.Uid
		path := c.Request.URL.Path
		key := uid + path

		d.doLimit(c, key, rate, capacity)

		c.Next()
	}
}

//IpLimit ip地址进行限流.
func (d *DisLimiter) IpLimit(rate, capacity int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		d.doLimit(c, ip, rate, capacity)

		c.Next()
	}
}

func (d *DisLimiter) doLimit(c *gin.Context, key string, rate, capacity int) {
	tokenKey, timestampKey := getKeys(key)
	run := luaScript.Run(d.Client, []string{tokenKey, timestampKey}, rate, capacity, time.Now().Second(), 1)
	if run.Err() != nil {
		ctxlog.WithCtx(c).Errorf("PathLimit execute failed:%v", run.Err())
	}

	allowed, ok := run.Val().([]interface{})
	if !ok {
		return
	}

	if allowed[0].(int64) == 0 {
		reqresp.ReturnError(c, code.ServerBusy)
		c.Abort()
	}
}

func getKeys(id string) (tokenKey string, timestampKey string) {
	return fmt.Sprintf(redisTokenKey, id), fmt.Sprintf(redisTimestampKey, id)
}
