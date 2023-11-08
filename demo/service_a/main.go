package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"servicea/middlewares"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/module/apmlogrus"
	"go.elastic.co/ecslogrus"
)

func init() {
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&ecslogrus.Formatter{})
	logrus.SetLevel(logrus.ErrorLevel)
	logrus.AddHook(&apmlogrus.Hook{})
}

func main() {
	r := gin.New()

	// Setup middleware
	r.Use(gin.Recovery())
	r.Use(middlewares.LoggingMiddleware())
	r.Use(apmgin.Middleware(r))

	r.GET("/hello", func(c *gin.Context) {
		span, ctx := apm.StartSpan(c.Request.Context(), "HelloHandler", "request")
		defer span.End()

		// 1. Long time processing request
		processingRequest(ctx)

		// 2. Call extenal api
		todo, err := getTodoFromAPI(ctx)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"todo": todo,
		})
	})
	r.Run() // Start server on 0.0.0.0:8080
}

func processingRequest(ctx context.Context) {
	span, ctx := apm.StartSpan(ctx, "processingRequest", "custom")
	defer span.End()

	doSomething(ctx)

	// time sleep simulate some processing time
	time.Sleep(15 * time.Millisecond)
	return
}

func doSomething(ctx context.Context) {
	span, ctx := apm.StartSpan(ctx, "doSomething", "custom")
	defer span.End()

	// time sleep simulate some processing time
	time.Sleep(20 * time.Millisecond)
	return
}

func getTodoFromAPI(ctx context.Context) (map[string]interface{}, error) {
	span, ctx := apm.StartSpan(ctx, "getTodoFromAPI", "custom")
	defer span.End()

	var result map[string]interface{}

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, err
}
