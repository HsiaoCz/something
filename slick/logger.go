package slick

import (
	"log/slog"
	"time"
)

func Logger() Handlerfunc {
	return func(ctx *Context) {
		t := time.Now()
		ctx.Next()
		slog.Info("the request has fished", "status", ctx.StatusCode, "path", ctx.R.URL.Path, "time cost", time.Since(t))
	}
}
