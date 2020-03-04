package glike

import (
	"fmt"
	"time"
)

func DefaultLogger() HandleFunc {

	return func(ctx *Context) {
		startTime := time.Now()
		path := ctx.request.URL.Path
		method := ctx.request.Method

		ctx.Next()

		endTime := time.Now()
		fmt.Printf("method: %v, path: %v, time cost: %v\n", method, path, endTime.Sub(startTime))
	}
}

func Recovery() HandleFunc {
	return func(ctx *Context) {

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %v\n", err)
				ctx.responseWriter.WriteHeader(400)
			}
		}()
		ctx.Next()
	}
}
