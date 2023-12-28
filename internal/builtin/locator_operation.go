package builtin

import "github.com/playwright-community/playwright-go"

var LocatorOperation = map[string]func(locator playwright.Locator, locatorType string){
	"click":        LocatorClick,
	"input":        LocatorInput,
	"enter":        LocatorPress,
	"hover":        LocatorHover,
	"right_click":  LocatorClick,
	"double_click": LocatorClick,
}

func LocatorClick(locator playwright.Locator, locatorType string) {
	if len(locatorType) > 0 {
		locator.Click(playwright.LocatorClickOptions{Button: playwright.MouseButtonRight})
		return
	}
	locator.Click()
}

func LocatorInput(locator playwright.Locator, locatorType string) {
	locator.Fill(locatorType)
}

func LocatorPress(locator playwright.Locator, locatorType string) {
	locator.Press(locatorType)

}

func LocatorHover(locator playwright.Locator, locatorType string) {
	locator.Hover()
}

// 定义一个映射，将动作名称映射到对应的执行函数

var MouseOperationList = []string{"drag_and_drop", "drag_element", "scroll", "wheel", "hover", "right_click", "double_click", "click"}
