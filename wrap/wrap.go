package wrap

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/server"
	"github.com/yaozhuangyanlingyu/micro-srv/logger"
	"go.uber.org/zap"
)

// 日志中间件
func LogRespWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		logger.Info("REQUEST", zap.String("method", req.Method()), zap.Reflect("body", req.Body()), zap.Int64("request_time", time.Since(time.Now()).Milliseconds()))
		err := fn(ctx, req, rsp)
		return err
	}
}
