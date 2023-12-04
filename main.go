package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

// customTestingT 用于存储失败的断言信息
type customTestingT struct {
	FailMessage string
}

// Errorf 记录断言失败的简略信息
func (t *customTestingT) Errorf(format string, args ...interface{}) {
	// 仅保存 "Not equal" 断言失败的基本信息
	if format == "Not equal: \n"+
		"expected: %v\n"+
		"actual  : %v" {
		t.FailMessage = fmt.Sprintf(format, args...)
	}

}

func main() {

	t := new(customTestingT)
	Demo(t)

	// 如果有失败信息，则打印
	if t.FailMessage != "" {
		fmt.Println("Error:", t.FailMessage)
	}
}

func Demo(t *customTestingT) {
	result := assert.Equal(t, "aa", "a")
	fmt.Println(result)
}
