package assertions

import (
	"errors"
	"fmt"
	"github.com/ZeroBugHero/web_ui_go/internal/builtin"
	"github.com/ZeroBugHero/web_ui_go/models"
	"github.com/ZeroBugHero/web_ui_go/test/reports"
	"github.com/playwright-community/playwright-go"
	"github.com/rs/zerolog/log"
	"slices"
	"strconv"
	"strings"
)

// CustomTestingT 用于存储失败的断言信息
type CustomTestingT struct {
	FailMessage string
	Actual      interface{}
	Expect      interface{}
}

// Errorf 记录断言失败的简略信息
func (t *CustomTestingT) Errorf(format string, args ...interface{}) {
	// 仅保存断言失败的基本信息
	t.FailMessage = fmt.Sprintf(format, args...)
}

// AssertLocator 根据定位器执行断言
func AssertLocator(page playwright.Page, locator models.Assert, n int) *CustomTestingT {
	t := new(CustomTestingT)
	if t.FailMessage == "" {
		t.FailMessage = "断言成功"
	} // 断言成功默认信息
	if page == nil || countLocatorValues(locator) == 0 {
		t.Errorf("Invalid parameters")
		return t
	}
	// 判断locator.Check.type是否在builtin.ValidateWithPlaywright数组中
	if _, ok := slices.BinarySearch(builtin.ValidateWithPlaywright, locator.Check.Type); ok {
		err := ValidateWithPlaywright(page, locator)
		if err != nil {
			t.Errorf("Error validating with playwright: %v", err)
			validatorStatic := reports.ValidatorStatic{
				Screenshot: "",
				Video:      "",
				Trace:      "",
			}
			result := reports.Validator{
				Id:              0,
				Type:            locator.Check.Type,
				Locator:         locator.ElementLocator.Values[0],
				Expect:          locator.Check.Expect,
				Actual:          nil,
				Message:         err.Error(),
				ValidatorStatic: validatorStatic,
			}
			result.AddResult()
			return t
		}
		return t
	}
	innerTexts, err := getInnerTextsBasedOnLocatorType(page, locator)
	if err != nil || len(innerTexts) == 0 {
		// 错误处理
		t.Errorf("Error fetching texts: %v", err)
		validatorStatic := reports.ValidatorStatic{
			Screenshot: "",
			Video:      "",
			Trace:      "",
		}
		result := reports.Validator{
			Id:              0,
			Type:            locator.Check.Type,
			Locator:         locator.ElementLocator.Values[0],
			Expect:          locator.Check.Expect,
			Actual:          nil,
			Message:         err.Error(),
			ValidatorStatic: validatorStatic,
		}
		result.AddResult()
		return t
	} else {
		// 检查索引是否在innerTexts数组范围内
		if locator.ElementLocator.Index > 0 && locator.ElementLocator.Index < len(locator.Check.Expect) {
			// 检查特定索引处的元素
			innerText := innerTexts[locator.ElementLocator.Index]
			assertBasedOnCheckType(t, innerText, locator.Check.Expect[locator.ElementLocator.Index], locator.Check)
			validatorStatic := reports.ValidatorStatic{
				Screenshot: "",
				Video:      "",
				Trace:      "",
			}
			result := reports.Validator{
				Id:              0,
				Type:            locator.Check.Type,
				Locator:         locator.ElementLocator.Values[0],
				Expect:          locator.Check.Expect,
				Actual:          innerText,
				Message:         captureMessage(t.FailMessage),
				ValidatorStatic: validatorStatic,
			}
			t.Actual = fmt.Sprintf("实际值:%v", innerText)
			t.Expect = fmt.Sprintf("期望值:%v", locator.Check.Expect[locator.ElementLocator.Index])
			result.AddResult()
		} else {
			// 根据locator.Check.Expect的长度进行匹配
			for i := 0; i < len(locator.Check.Expect); i++ {
				innerText := innerTexts[i]
				assertBasedOnCheckType(t, innerText, locator.Check.Expect[i], locator.Check)
				validatorStatic := reports.ValidatorStatic{
					Screenshot: "",
					Video:      "",
					Trace:      "",
				}
				result := reports.Validator{
					Id:              0,
					Type:            locator.Check.Type,
					Locator:         locator.ElementLocator.Values[0],
					Expect:          locator.Check.Expect,
					Actual:          innerText,
					Message:         captureMessage(t.FailMessage),
					ValidatorStatic: validatorStatic,
				}
				t.Actual = fmt.Sprintf("实际值:%v", innerText)
				t.Expect = fmt.Sprintf("期望值:%v", locator.Check.Expect[locator.ElementLocator.Index])
				result.AddResult()
			}
		}
		// 根据断言类型执行断言检查
		//log.Debug().Msgf("详细断言失败的信息:%v", *t) //详细的断言失败日志打印
		t.FailMessage = captureMessage(t.FailMessage)
		log.Info().Msg(fmt.Sprintf("第%d次断言结束：%s", n, *t))
		return t
	}

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
	case "title":
		title, err := page.Title()
		return []string{title}, err
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
	if assert.Type == "title" {
		return 1
	}
	return len(assert.ElementLocator.Values)
}

func ValidateWithPlaywright(page playwright.Page, locator models.Assert) error {
	playwrightAssertions := playwright.NewPlaywrightAssertions()
	pwLocator := page.Locator(locator.ElementLocator.Values[0])
	switch locator.Type {
	case "to_have_text":
		return playwrightAssertions.Locator(pwLocator).ToHaveText(locator.Check.Expect[0])
	case "to_have_value":
		return playwrightAssertions.Locator(pwLocator).ToHaveValue(locator.Check.Expect[0])
	case "to_be_checked":
		return playwrightAssertions.Locator(pwLocator).ToBeChecked()
	case "to_be_disabled":
		return playwrightAssertions.Locator(pwLocator).ToBeDisabled()
	case "to_be_editable":
		return playwrightAssertions.Locator(pwLocator).ToBeEditable()
	case "to_be_enabled":
		return playwrightAssertions.Locator(pwLocator).ToBeEnabled()
	case "to_be_focused":
		return playwrightAssertions.Locator(pwLocator).ToBeFocused()
	case "to_be_visible":
		return playwrightAssertions.Locator(pwLocator).ToBeVisible()
	case "to_be_hidden":
		return playwrightAssertions.Locator(pwLocator).ToBeHidden()
	case "to_be_selected":
		num, _ := strconv.Atoi(locator.Check.Expect[0])
		return playwrightAssertions.Locator(pwLocator).ToHaveCount(num)
	case "to_contain_text":
		return playwrightAssertions.Locator(pwLocator).ToContainText(locator.Check.Expect[0])
	case "to_have_id":
		return playwrightAssertions.Locator(pwLocator).ToHaveId(locator.Check.Expect[0])

	}
	return nil
}

// captureMessage 截取消息
func captureMessage(message string) string {
	startIndex := strings.Index(message, "Error:")
	if startIndex != -1 {
		// 从"Error:"开始截取
		message = message[startIndex:]
	}
	return message
}
