package errors

var (
	ErrInternalServer = NewResponse(500, 500, "服务器发生错误")
)
