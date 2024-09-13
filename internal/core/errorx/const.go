package errorx

// custom error code/msg to return

//type CodeError struct {
//	errCode uint32
//	errMsg  string
//}
//
//const OK uint32 = 200
//
//const SERVER_COMMON_ERROR uint32 = 500000
//const REUQEST_PARAM_ERROR uint32 = 400000
//const TOKEN_EXPIRE_ERROR uint32 = 403000
//
//var (
//    message map[uint32]string
//	UnAuthorizedError = NewErrCodeMsg(TOKEN_EXPIRE_ERROR, "登录过期失效，请重新登陆")
//	InternalError     = NewErrCodeMsg(SERVER_COMMON_ERROR, "内部异常")
//)
//
//
//func init() {
//	message = make(map[uint32]string)
//	message[OK] = "SUCCESS"
//	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
//	message[REUQEST_PARAM_ERROR] = "参数错误"
//	message[TOKEN_EXPIRE_ERROR] = "登录过期失效，请重新登陆"
//}
//
//func MapErrMsg(errcode uint32) string {
//	if msg, ok := message[errcode]; ok {
//		return msg
//	} else {
//		return "服务器开小差啦,稍后再来试一试"
//	}
//}
//
//func IsCodeErr(errcode uint32) bool {
//	if _, ok := message[errcode]; ok {
//		return true
//	} else {
//		return false
//	}
//}
//
//// GetErrCode 返回给前端的错误码
//func (e *CodeError) GetErrCode() uint32 {
//	return e.errCode
//}
//
//// GetErrMsg 返回给前端显示端错误信息
//func (e *CodeError) GetErrMsg() string {
//	return e.errMsg
//}
//
//func (e *CodeError) Error() string {
//	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
//}
//
//func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
//	return &CodeError{errCode: errCode, errMsg: errMsg}
//}
//func NewErrCode(errCode uint32) *CodeError {
//	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
//}
//
//func NewErrMsg(errMsg string) *CodeError {
//	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
//}
