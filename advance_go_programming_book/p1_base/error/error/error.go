package error

//
//type Error interface {
//	Caller() []CallerInfo
//	Wrapped() []error
//	Code() int
//
//	private()
//}
//
//type CallerInfo struct {
//	FuncName string
//	FileNmae string
//	FileLine string
//}
//
//func New(msg string) error
//func NewWithCode(code int, msg string) error
//func Wrap(err error, msg string) error
//func WrapWithCode(code int, err error, msg string) error
//func FromJson(json string) error
//func ToJson(err error) string
