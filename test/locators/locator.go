package locators

import (
	"github.com/ZeroBugHero/web_ui_go/internal/builtin"
	"github.com/ZeroBugHero/web_ui_go/launch"
	"github.com/ZeroBugHero/web_ui_go/models"
	"github.com/playwright-community/playwright-go"
	"github.com/rs/zerolog/log"
	"time"
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
	ExecuteInteractiveAction(element, element, locator, page)
}

// 在程序启动时对builtin.MouseOperationList进行排序和转换为map

var MouseOperationSet map[string]bool

func init() {
	MouseOperationSet = make(map[string]bool)
	for _, operation := range builtin.MouseOperationList {
		MouseOperationSet[operation] = true
	}
}

// ExecuteInteractiveAction 根据定位器执行相应动作
func ExecuteInteractiveAction(element playwright.Locator, targetElement playwright.Locator, locator models.Locator, page playwright.Page) {
	if element == nil {
		log.Error().Msg("定位器为空")
		return
	}
	log.Info().Msgf("执行动作: %s", locator.Operation.Action.Interactive)
	// 判断locator.Operation.Action是否在MouseOperationSet中
	if _, ok := MouseOperationSet[locator.Operation.Action.Interactive]; ok {
		mouseOperation(element, targetElement, locator, page)
	} else {
		// 如果在映射中找到了对应的动作，执行该动作
		if actionFunc, ok := actionFuncs[locator.Operation.Action.Interactive]; ok {
			actionFunc(element, locator, &locator.Timeout)
		} else {
			log.Error().Msg("不支持的动作")
		}
	}
}

func countLocatorValues(step models.Locator) int {
	return len(step.Values)
}
