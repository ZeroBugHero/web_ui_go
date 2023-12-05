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

func (l *LocatorTestStep) stepLocatorPage() playwright.Page {
	browser := launch.Browser(launch.GlobalConfig.Browser)
	page, err := browser.NewPage(playwright.BrowserNewPageOptions{BaseURL: playwright.String(launch.GlobalConfig.BaseUrl)})
	if err != nil {
		log.Error().Err(err).Msg("创建新页面失败")
	}
	defer browser.Close()
	defer launch.ClosePlaywright()
	l.StartTime = time.Now().Local()
	return page
}

// stepLocatorMethod 定位器方法
func stepLocatorMethod(testSteps *models.TestStep) {
	if testSteps == nil {
		return
	}
	l := LocatorTestStep{}
	page := l.stepLocatorPage()
	switch testSteps.Locator.Type {
	case "id":
		locatorAction(page, testSteps.Locator)
	case "xpath":
		locatorAction(page, testSteps.Locator)
	case "role":
		locatorRole(page, testSteps.Locator)
	case "text":
		locatorAction(page, testSteps.Locator)

	}

}

// locatorAction 定位器动作
func locatorAction(page playwright.Page, locator models.Locator) {
	if lenValues(locator) == 1 {
		switch locator.Operation.Action.Interactive {
		case "click":
			page.Locator(locator.Values[0]).Click()
		case "input":
			page.Locator(locator.Values[0]).Fill(locator.Operation.Action.Input)
		case "enter":
			page.Locator(locator.Values[0]).Press(locator.Operation.Action.Interactive)
		case "hover":
			page.Locator(locator.Values[0]).Hover()
		case "right_click":
			page.Locator(locator.Values[0]).Click(playwright.LocatorClickOptions{Button: playwright.MouseButtonRight})
		case "double_click":
			page.Locator(locator.Values[0]).Dblclick()

		}

	} else { // 待优化，自动匹配values值的数量
		switch locator.Operation.Action.Interactive {
		case "click":
			page.Locator(locator.Values[0]).Locator(locator.Values[1]).Click()

		}
	}

}

func locatorRole(page playwright.Page, locator models.Locator) {
	if lenValues(locator) == 1 {
		switch locator.Operation.Action.Interactive {
		case "click":
			page.GetByRole(playwright.AriaRole(locator.Values[0]), playwright.PageGetByRoleOptions{Name: locator.Values[1], Exact: playwright.Bool(locator.Exact)}).Click()
		case "input":
			page.GetByRole(playwright.AriaRole(locator.Values[0]), playwright.PageGetByRoleOptions{Name: locator.Values[1], Exact: playwright.Bool(locator.Exact)}).Fill(locator.Operation.Action.Input)
		case "enter":
			page.GetByRole(playwright.AriaRole(locator.Values[0]), playwright.PageGetByRoleOptions{Name: locator.Values[1], Exact: playwright.Bool(locator.Exact)}).Press(locator.Operation.Action.Interactive)
		case "hover":
			page.GetByRole(playwright.AriaRole(locator.Values[0]), playwright.PageGetByRoleOptions{Name: locator.Values[1], Exact: playwright.Bool(locator.Exact)}).Hover()
		case "right_click":
			page.GetByRole(playwright.AriaRole(locator.Values[0]), playwright.PageGetByRoleOptions{Name: locator.Values[1], Exact: playwright.Bool(locator.Exact)}).Click(playwright.LocatorClickOptions{Button: playwright.MouseButtonRight})
		case "double_click":
			page.GetByRole(playwright.AriaRole(locator.Values[0]), playwright.PageGetByRoleOptions{Name: locator.Values[1], Exact: playwright.Bool(locator.Exact)}).Dblclick()
		}

	} else { // 待优化，自动匹配values值的数量
		switch locator.Operation.Action.Interactive {
		case "click":
			page.Locator(locator.Values[0]).Locator(locator.Values[1]).Click()

		}
	}
}
func lenValues(step models.Locator) int {
	return len(step.Values)
}
