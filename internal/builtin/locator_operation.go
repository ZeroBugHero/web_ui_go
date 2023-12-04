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
