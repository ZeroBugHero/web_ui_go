package models

import "time"

// TestStep 是用于单个测试步骤的结构体
type TestStep struct {
	StartTime time.Time // 默认为当前时间 time.Now()
	Locator   Locator   `yaml:"locator" json:"locator"`
	Assert    Assert    `yaml:"assert" json:"assert"`
}

// Locator 是用于定位元素的结构体
type Locator struct {
	Name      string    `yaml:"name" json:"name"`
	Type      string    `yaml:"type" json:"type"` // 定位器类型 css, role, test-id, text, placeholder, label, xpath
	Exact     bool      `yaml:"exact" json:"exact"`
	Values    []string  `yaml:"values" json:"values"` // 定位器的值
	Operation Operation `yaml:"operation" json:"operation"`
	Timeout   float64   `yaml:"timeout" json:"timeout"`
}

// Operation 是用于操作元素的结构体
type Operation struct {
	Action Action `yaml:"action" json:"action"`
}

// Action 是用于操作元素的结构体
type Action struct {
	Input       string `yaml:"input" json:"input"`
	Interactive string `yaml:"interactive" json:"interactive"` // 操作类型 click, input, enter, select, scroll, hover, right_click, double_click, drag_and_drop, drag_and_drop_by_offset, press, type, upload_file
	Coordinates []int  `yaml:"coordinates" json:"coordinates"`
}

// Assert 是用于断言的结构体
type Assert struct {
	Name           string         `yaml:"name" json:"name"`
	Type           string         `yaml:"type" json:"type"` // 断言类型 css, role, test-id, text, placeholder, label, xpath
	Exact          bool           `yaml:"exact" json:"exact"`
	ElementLocator ElementLocator `yaml:"element_locator" json:"element_locator"` // 元素定位器
	Check          Check          `yaml:"check" json:"check"`                     // 断言检查
	Continue       bool           `yaml:"continue" json:"continue"`               // 断言失败是否继续
	Timeout        int            `yaml:"timeout" json:"timeout"`                 // 断言超时时间
}

// Check 是用于检查的结构体
type Check struct {
	Type   string   `yaml:"type" json:"type"`     // 断言方式 equal, not_equal, contains, not_contains, greater_than, less_than
	Expect []string `yaml:"expect" json:"expect"` // 期望值
}

type ElementLocator struct {
	Values []string `yaml:"values" json:"values"` // 元素定位器的值
	Index  int      `yaml:"index" json:"index"`   // 下标
}
