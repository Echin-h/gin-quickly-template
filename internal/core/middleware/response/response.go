package response

import (
	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Code    int32  `json:"code"`
	Error   any    `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func http(c *gin.Context, code int32, msg string, data any, err any) {
	if code == 0 {
		c.JSON(200, &JsonResponse{
			Code:    code,
			Message: msg,
			Data:    data,
		})
		return
	}
	c.JSON(int(code/1000), &JsonResponse{
		Code:    code / 1000,
		Message: msg,
		Data:    data,
		Error:   err,
	})
}

func Success(c *gin.Context, data any) {
	http(c, 0, "success", data, nil)
}

func Fail(c *gin.Context, code int, msg string, err ...any) {
	newErrs := make([]string, 0, len(err))

	for _, e := range err {
		switch v := e.(type) {
		case []error:
			for _, ea := range v {
				newErrs = append(newErrs, ea.Error())
			}
		case error:
			newErrs = append(newErrs, v.Error())
		case string:
			newErrs = append(newErrs, v)
		case []string:
			newErrs = append(newErrs, v...)
		case []interface{}:
			for _, item := range v {
				switch item := item.(type) {
				case error:
					newErrs = append(newErrs, item.Error())
				case string:
					newErrs = append(newErrs, item)
				default:
					newErrs = append(newErrs, "Unknown error")
				}
			}
		default:
			newErrs = append(newErrs, "Unknown error")
		}
	}

	// 返回错误响应
	http(c, int32(code), msg, nil, newErrs)
}

func FailWithData(c *gin.Context, code int, msg string, data any, errs ...any) {
	for i, e := range errs {
		if v, ok := e.(error); ok {
			errs[i] = v.Error()
		}
	}
	http(c, int32(code), msg, data, errs)
}

func UnAuthorization(c *gin.Context) {
	Fail(c, 401000, "登录过期失效，请重新登陆")
}

func Forbidden(c *gin.Context) {
	Fail(c, 403000, "未授权")
}

func InValidParam(c *gin.Context, err ...any) {
	Fail(c, 400000, "请求校验失败", err...)
}

func ServerError(c *gin.Context, err error) {
	Fail(c, 500000, "内部异常", err)
}
