package response

const (
	// SUCCESS 成功
	SUCCESS = 0
	// FAIL 失败
	FAIL = -1
)

// Result 统一返回结果
type Result struct {
	Code int         `json:"code" binding:"required"`
	Data interface{} `json:"data" binding:"required"`
	Msg  string      `json:"msg" binding:"required"`
}

// Success 成功结果构造函数
func Success(data interface{}, msg string) *Result {
	return &Result{Code: SUCCESS, Data: data, Msg: msg}
}

//Fail 失败的结果
func Fail(data interface{}, msg string) *Result {
	return &Result{Code: FAIL, Data: data, Msg: msg}
}

// New 创建一个新的普通的结果
func New(code int, data interface{}, msg string) *Result {
	return &Result{Code: code, Data: data, Msg: msg}
}
