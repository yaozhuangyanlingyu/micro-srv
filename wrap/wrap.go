package wrap

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/micro/go-micro/v2/server"
	"github.com/yaozhuangyanlingyu/micro-srv/logger"
	"go.uber.org/zap"
)

// 日志中间件
func LogRespWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Println(reflect.ValueOf(rsp), reflect.TypeOf(rsp))

		// 验证proto字段
		v := reflect.ValueOf(rsp).Elem()
		code := v.FieldByName("Code")
		msg := v.FieldByName("Msg")
		if !code.IsValid() || !msg.IsValid() {
			return errors.New("rsp必须有Code和Msg两个字段")
		}

		// 记录请求日志
		logger.Info(
			"REQUEST",
			zap.String("method", req.Method()),
			zap.Reflect("body", req.Body()),
			zap.Int64("request_time", time.Since(time.Now()).Milliseconds()),
		)

		// 调用业务代码
		err := fn(ctx, req, rsp)
		return err
	}
}
