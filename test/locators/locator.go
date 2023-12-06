package locators

import (
	"github.com/playwright-community/playwright-go"
	"github.com/rs/zerolog/log"
	"time"
	"web_ui_go/launch"
	"web_ui_go/models"
)

// LocatorTestStep 定位器
type LocatorTestStep models.TestStep

// CreateAndReturnNewPage 创建并返回新的页面对象
func (l *LocatorTestStep) CreateAndReturnNewPage() playwright.Page {
	browser := launch.Browser(launch.GlobalConfig.Browser)
	page, err := browser.NewPage(playwright.BrowserNewPageOptions{BaseURL: playwright.String(launch.GlobalConfig.BaseUrl)})
	if err != nil {
		log.Error().Err(err).Msg("创建新页面失败")
		browser.Close()
		launch.ClosePlaywright()
		return nil
	}
	defer func() {
		browser.Close()
		launch.ClosePlaywright()
	}()
	l.StartTime = time.Now().Local()
	return page
}

// PerformActionBasedOnLocator 根据定位器执行动作
func PerformActionBasedOnLocator(page playwright.Page, locator models.Locator) {
	// 参数验证
	if page == nil || countLocatorValues(locator) == 0 {
		log.Error().Msg("无效的参数")
		return
	}

	// 根据类型选择定位器
	var element playwright.Locator

	switch locator.Type {
	case "css":
		element = page.Locator(locator.Values[0])
	case "role":
		element = page.GetByRole(playwright.AriaRole(locator.Values[0]), playwright.PageGetByRoleOptions{Name: locator.Values[1], Exact: playwright.Bool(locator.Exact)})
	case "test-id":
		element = page.GetByTestId(locator.Values[0])
	case "text":
		element = page.GetByText(locator.Values[0], playwright.PageGetByTextOptions{Exact: playwright.Bool(locator.Exact)})
	case "placeholder":
		element = page.GetByPlaceholder(locator.Values[0], playwright.PageGetByPlaceholderOptions{Exact: playwright.Bool(locator.Exact)})
	case "label":
		element = page.GetByLabel(locator.Values[0], playwright.PageGetByLabelOptions{Exact: playwright.Bool(locator.Exact)})
	default:
		element = page.Locator(locator.Values[0])
	}

	// 执行动作
	ExecuteInteractiveAction(element, locator)
}

// ExecuteInteractiveAction 根据定位器执行相应动作
func ExecuteInteractiveAction(element playwright.Locator, locator models.Locator) {
	if element == nil {
		log.Error().Msg("定位器为空")
		return
	}

	log.Info().Msgf("执行动作: %s", locator.Operation.Action.Interactive)
	// 参考assert.go中的assertBasedOnCheckType函数，减少switch case的嵌套 todo
	switch locator.Operation.Action.Interactive {
	case "click":
		clickAction(element, playwright.Float(locator.Timeout))
	case "input":
		inputAction(element, locator.Operation.Action.Input, playwright.Float(locator.Timeout))
	case "enter":
		enterAction(element, locator.Operation.Action.Interactive, playwright.Float(locator.Timeout))
	case "hover":
		hoverAction(element, playwright.Float(locator.Timeout))
	case "right_click":
		rightClickAction(element, playwright.Float(locator.Timeout))
	case "double_click":
		doubleClickAction(element, playwright.Float(locator.Timeout))
	// ... [其他动作] ...
	default:
		log.Warn().Msgf("未知动作: %s", locator.Operation.Action.Interactive)
	}
}

// 以下是具体动作的函数实现
func clickAction(element playwright.Locator, timeout *float64) {
	element.Click(playwright.LocatorClickOptions{Timeout: timeout})
}

func inputAction(element playwright.Locator, input string, timeout *float64) {
	element.Fill(input, playwright.LocatorFillOptions{Timeout: timeout})
}

func enterAction(element playwright.Locator, key string, timeout *float64) {
	element.Press(key, playwright.LocatorPressOptions{Timeout: timeout})
}

func hoverAction(element playwright.Locator, timeout *float64) {
	element.Hover(playwright.LocatorHoverOptions{Timeout: timeout})
}

func rightClickAction(element playwright.Locator, timeout *float64) {
	element.Click(playwright.LocatorClickOptions{Button: playwright.MouseButtonRight, Timeout: timeout})
}

func doubleClickAction(element playwright.Locator, timeout *float64) {
	element.Dblclick(playwright.LocatorDblclickOptions{Timeout: timeout})
}

// ... [其他动作的函数实现] ...

func countLocatorValues(step models.Locator) int {
	return len(step.Values)
}
