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
	if err != nil || len(innerTexts) == 0 {
		log.Error().Err(err).Msg("获取文本失败")
		t.Errorf("Error fetching texts: %v", err)
		return t
	}
	if locator.ElementLocator.Index < 0 {
		log.Info().Msg("索引不能为负数，将其设置为0。")
		locator.ElementLocator.Index = 0
		t.FailMessage = "索引不能为负数，将其设置为0。"
		return t
	} else if locator.ElementLocator.Index == 0 || locator.ElementLocator.Index == 1 {
		innerText := innerTexts[0]
		assertBasedOnCheckType(t, innerText, locator.Check)
	} else {
		// 循环断言
		for i := 0; i < locator.ElementLocator.Index; i++ {
			innerText := innerTexts[i]
			assertBasedOnCheckType(t, innerText, locator.Check)
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
func assertBasedOnCheckType(t *CustomTestingT, innerTexts interface{}, check models.Check) {
	if assertionFunc, ok := builtin.Assertions[check.Type]; ok {
		// 调用断言函数
		assertionFunc(t, innerTexts, check.Expect)
	} else {
		// 如果没有找到对应的断言类型
		t.Errorf("不支持的断言类型: %s", check.Type)
	}
}
func CountLocatorValues(assert models.Assert) int {
	return len(assert.ElementLocator.Values)
}
