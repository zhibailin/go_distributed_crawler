// 功能函数
package rpcdemo // 包名避免与 rpc 冲突

import "errors"

type DemoService struct {
}

type Args struct {
	A, B int
}

// 固定参数：第一个是输入，第二个是输出
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
