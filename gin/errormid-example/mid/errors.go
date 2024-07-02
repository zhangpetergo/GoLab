package mid

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhangpetergo/GoLab/gin/errormid-example/errs"
	"github.com/zhangpetergo/GoLab/slog-example/foundation/logger"
	"net/http"
)

var codeStatus [17]int

// init maps out the error codes to http status codes.
func init() {
	codeStatus[errs.OK.Value()] = http.StatusOK
	codeStatus[errs.Canceled.Value()] = http.StatusGatewayTimeout
	codeStatus[errs.Unknown.Value()] = http.StatusInternalServerError
	codeStatus[errs.InvalidArgument.Value()] = http.StatusBadRequest
	codeStatus[errs.DeadlineExceeded.Value()] = http.StatusGatewayTimeout
	codeStatus[errs.NotFound.Value()] = http.StatusNotFound
	codeStatus[errs.AlreadyExists.Value()] = http.StatusConflict
	codeStatus[errs.PermissionDenied.Value()] = http.StatusForbidden
	codeStatus[errs.ResourceExhausted.Value()] = http.StatusTooManyRequests
	codeStatus[errs.FailedPrecondition.Value()] = http.StatusBadRequest
	codeStatus[errs.Aborted.Value()] = http.StatusConflict
	codeStatus[errs.OutOfRange.Value()] = http.StatusBadRequest
	codeStatus[errs.Unimplemented.Value()] = http.StatusNotImplemented
	codeStatus[errs.Internal.Value()] = http.StatusInternalServerError
	codeStatus[errs.Unavailable.Value()] = http.StatusServiceUnavailable
	codeStatus[errs.DataLoss.Value()] = http.StatusInternalServerError
	codeStatus[errs.Unauthenticated.Value()] = http.StatusUnauthorized
}

func Error(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next() // 处理请求
		ctx := c.Request.Context()
		log.Info(ctx, "hello", "")
		if len(c.Errors) > 0 {
			// 有错误发生
			err := c.Errors[0].Err
			// 记录错误
			log.Error(ctx, "message", "ERROR", err.Error())
			// 返回的是我们自定义的错误
			if errs.IsError(err) {
				var errs errs.Error
				errors.As(err, &errs)
				// 修改返回内容
				c.JSON(codeStatus[errs.Code.Value()], errs)
				return
			}
			errs := errs.Newf(errs.Unknown, errs.Unknown.String())
			c.JSON(codeStatus[errs.Code.Value()], errs)
		}
	}
}
