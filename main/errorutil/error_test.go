package errorutil

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

func Test_error(t *testing.T) {
	fmt.Println(test("test"))
}

func test(name string) error{
	if name == "test" {
		return Wrap(errors.New("test"))
	}

	return nil
}

// 包括一个错误，可以构建调用链，用于偏上层操作错误
func Wrap(err error) error {
	_, filename, line, _ := runtime.Caller(1)
	//lpsErr := Error{
	//	codeMsg:    cm,
	//	cause:      err.cause,
	//	InvokeLink: NewInvokeNode(param, err.cause.Error()),
	//	Location:   fmt.Sprintf("file %s, line %v", pathutil.GetFileRelativePath(filename), line),
	//}
	//lpsErr.InvokeLink.Next = err.InvokeLink
	return errors.New(filename + ":" + strconv.Itoa(line))
}

func errortest() error{
	return errors.New("测试")
}