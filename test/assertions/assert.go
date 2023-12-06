package assertions

import (
	"errors"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"github.com/rs/zerolog/log"
	"web_ui_go/internal/builtin"
	"web_ui_go/models"
)

// CustomTestingT 用于存储失败的断言信息
type CustomTestingT struct {
	FailMessage string
}

// Errorf 记录断言失败的简略信息
func (t *CustomTestingT) Errorf(format string, args ...interface{}) {
	// 仅保存断言失败的基本信息
	t.FailMessage = fmt.Sprintf(format, args...)
}

// AssertLocator 根据定位器执行断言
func AssertLocator(page playwright.Page, locator models.Assert) *CustomTestingT {
	t := new(CustomTestingT)

	if page == nil || countLocatorValues(locator) == 0 {
		t.Errorf("Invalid parameters")
		return t
	}

	innerTexts, err := getInnerTextsBasedOnLocatorType(page, locator)
	if err != nil || len(innerTexts) == 0 {
		// 错误处理
		t.Errorf("Error fetching texts: %v", err)
		return t
	}

	// 检查索引是否在innerTexts数组范围内
	if locator.ElementLocator.Index > 0 && locator.ElementLocator.Index < len(locator.Check.Expect) {
		// 检查特定索引处的元素
		innerText := innerTexts[locator.ElementLocator.Index]
		assertBasedOnCheckType(t, innerText, locator.Check.Expect[locator.ElementLocator.Index], locator.Check)
	} else {
		// 根据locator.Check.Expect的长度进行匹配
		for i := 0; i < len(locator.Check.Expect); i++ {
			innerText := innerTexts[i]
			assertBasedOnCheckType(t, innerText, locator.Check.Expect[i], locator.Check)
		}
	}
	// 根据断言类型执行断言检查
	log.Info().Msg(fmt.Sprintf("断言成功，断言类型为：%s", t))

	return t
}

// getInnerTextsBasedOnLocatorType 根据定位器类型获取所有内部文本
func getInnerTextsBasedOnLocatorType(page playwright.Page, locator models.Assert) ([]string, error) {
	switch locator.Type {
	case "css":
		return page.Locator(locator.ElementLocator.Values[0]).AllInnerTexts()
	case "role":
		return page.GetByRole(playwright.AriaRole(locator.ElementLocator.Values[0]), playwright.PageGetByRoleOptions{Name: locator.ElementLocator.Values[1], Exact: playwright.Bool(locator.Exact)}).AllInnerTexts()
	case "test-id":
		return page.GetByTestId(locator.ElementLocator.Values[0]).AllInnerTexts()
	case "text":
		return page.GetByText(locator.ElementLocator.Values[0], playwright.PageGetByTextOptions{Exact: playwright.Bool(locator.Exact)}).AllInnerTexts()
	case "placeholder":
		return page.GetByPlaceholder(locator.ElementLocator.Values[0], playwright.PageGetByPlaceholderOptions{Exact: playwright.Bool(locator.Exact)}).AllInnerTexts()
	case "label":
		return page.GetByLabel(locator.ElementLocator.Values[0], playwright.PageGetByLabelOptions{Exact: playwright.Bool(locator.Exact)}).AllInnerTexts()
	case "xpath":
		return page.Locator(locator.ElementLocator.Values[0]).AllInnerTexts()
	default:
		return nil, errors.New("不支持的定位器类型")
	}
}

// assertBasedOnCheckType 根据断言类型执行断言检查
func assertBasedOnCheckType(t *CustomTestingT, innerTexts, expect interface{}, check models.Check) {
	if assertionFunc, ok := builtin.Assertions[check.Type]; ok {
		// 调用断言函数
		assertionFunc(t, innerTexts, expect)
	} else {
		// 如果没有找到对应的断言类型
		t.Errorf("不支持的断言类型: %s", check.Type)
	}
}
func countLocatorValues(assert models.Assert) int {
	return len(assert.ElementLocator.Values)
}
