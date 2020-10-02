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
