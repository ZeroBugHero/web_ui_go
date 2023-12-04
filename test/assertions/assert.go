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

	if page == nil || CountLocatorValues(locator) == 0 {
		t.Errorf("Invalid parameters")
		return t
	}

	innerTexts, err := getInnerTextsBasedOnLocatorType(page, locator)
	if err != nil {
		log.Error().Err(err).Msg("获取文本失败")
		t.Errorf("Error fetching texts: %v", err)
		return t
	}

	// 根据断言类型执行断言检查
	assertBasedOnCheckType(t, innerTexts, locator.Check)
	return t
}

// getInnerTextsBasedOnLocatorType 根据定位器类型获取所有内部文本
func getInnerTextsBasedOnLocatorType(page playwright.Page, locator models.Assert) ([]string, error) {
	switch locator.Type {
	case "css":
		return page.Locator(locator.Values[0]).AllInnerTexts()
	case "role":
		return page.GetByRole(playwright.AriaRole(locator.Values[0]), playwright.PageGetByRoleOptions{Name: locator.Values[1], Exact: playwright.Bool(locator.Exact)}).AllInnerTexts()
	case "test-id":
		return page.GetByTestId(locator.Values[0]).AllInnerTexts()
	case "text":
		return page.GetByText(locator.Values[0], playwright.PageGetByTextOptions{Exact: playwright.Bool(locator.Exact)}).AllInnerTexts()
	case "placeholder":
		return page.GetByPlaceholder(locator.Values[0], playwright.PageGetByPlaceholderOptions{Exact: playwright.Bool(locator.Exact)}).AllInnerTexts()
	case "label":
		return page.GetByLabel(locator.Values[0], playwright.PageGetByLabelOptions{Exact: playwright.Bool(locator.Exact)}).AllInnerTexts()
	case "xpath":
		return page.Locator(locator.Values[0]).AllInnerTexts()
	default:
		return nil, errors.New("不支持的定位器类型")
	}
}

// assertBasedOnCheckType 根据断言类型执行断言检查
func assertBasedOnCheckType(t *CustomTestingT, innerTexts []string, check models.Check) {
	if assertionFunc, ok := builtin.Assertions[check.Type]; ok {
		// 调用断言函数
		assertionFunc(t, innerTexts, check.Expect)
	} else {
		// 如果没有找到对应的断言类型
		t.Errorf("不支持的断言类型: %s", check.Type)
	}
}
func CountLocatorValues(assert models.Assert) int {
	return len(assert.Values)
}
