package main

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	//"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
)

// 第一步: 开一个全局变量
var zkTracer opentracing.Tracer

func main() {
	// 第二步: 初始化 tracer
	{
		reporter := zkHttp.NewReporter("http://134.175.227.112:9411/api/v2/spans")
		defer reporter.Close()
		endpoint, err := zipkin.NewEndpoint("study", "localhost:80")
		if err != nil {
			log.Fatalf("unable to create local endpoint: %+v\n", err)
		}
		nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		if err != nil {
			log.Fatalf("unable to create tracer: %+v\n", err)
		}
		zkTracer = zkOt.Wrap(nativeTracer)
		opentracing.SetGlobalTracer(zkTracer)
	}

	r := gin.Default()
	// 第三步: 添加一个 middleWare, 为每一个请求添加span
	r.Use(func(c *gin.Context) {
		log.Println("==========use===========")
		span := opentracing.StartSpan("root")
		ctx := opentracing.ContextWithSpan(context.Background(), span)
		c.Set("zctx", ctx)
		c.Next()
		span.Finish()
	})
	r.GET("/",
		func(c *gin.Context) {
			time.Sleep(500 * time.Millisecond)
			c.JSON(200, gin.H{"code": 200, "msg": "OK"})
		})
	r.GET("/app",
		func(c *gin.Context) {
			log.Println("==========app===========")
			zctx, ok := c.Get("zctx")
			if ok {
				span := opentracing.SpanFromContext(zctx.(context.Context))
				span.SetOperationName("app-test")
				defer span.Finish()
			}
			SpanTest11(zctx.(context.Context))
			c.JSON(200, gin.H{"code": 200, "msg": "OK"})
		})
	r.Run(":80")
}

func SpanTest11(ctx context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "test11")
	defer span.Finish()

	time.Sleep(1000 * time.Millisecond)
	SpanTest22(ctx)
}

func SpanTest22(ctx context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "test22")
	defer span.Finish()

	time.Sleep(1000 * time.Millisecond)
}

func SpanTest1(ctx opentracing.SpanContext) {
	span := opentracing.StartSpan("SpanTest1", opentracing.ChildOf(ctx))
	defer span.Finish()

	time.Sleep(1000 * time.Millisecond)
	SpanTest2(span.Context())
}

func SpanTest2(ctx opentracing.SpanContext) {
	span := opentracing.StartSpan("SpanTest2", opentracing.ChildOf(ctx))
	defer span.Finish()

	time.Sleep(1000 * time.Millisecond)
}
