package locators

import (
	"github.com/ZeroBugHero/web_ui_go/models"
	"github.com/playwright-community/playwright-go"
	"github.com/rs/zerolog/log"
)

var actionFuncs = map[string]func(element playwright.Locator, locator models.Locator, timeout *float64){
	"input": inputAction,
	"enter": enterAction,
	// ... 其他动作 ...
}

func mouseOperation(element playwright.Locator, targetElement playwright.Locator, locator models.Locator, page playwright.Page) {
	switch locator.Operation.Action.Interactive {
	case "drag_and_drop": // 拖拽
		box, err := element.BoundingBox(playwright.LocatorBoundingBoxOptions{
			Timeout: playwright.Float(locator.Timeout)})
		if err != nil {
			log.Error().Err(err).Msg("获取元素边界框失败")
			return
		}
		page.Mouse().Move(box.X+box.Width/2, box.Y+box.Height/2) // 移动到元素中心
		page.Mouse().Down()
		page.Mouse().Move(locator.Operation.Action.Coordinates.X, locator.Operation.Action.Coordinates.Y)
		page.Mouse().Up()
	case "drag_element": // 拖拽元素
		if targetElement == nil {
			log.Error().Msg("目标元素为空")
			return
		}
		err := element.DragTo(targetElement, playwright.LocatorDragToOptions{
			Timeout: playwright.Float(locator.Timeout)})
		if err != nil {
			log.Error().Err(err).Msg("拖拽元素失败")
			return
		}
	case "scroll": // 滚动到元素可见
		err := element.ScrollIntoViewIfNeeded(playwright.LocatorScrollIntoViewIfNeededOptions{
			Timeout: playwright.Float(locator.Timeout),
		})
		if err != nil {
			log.Error().Err(err).Msg("滚动到元素失败")
			return
		}
	case "wheel": // 滚动鼠标
		err := page.Mouse().Wheel(0, locator.Operation.Action.Coordinates.Y*1000)
		if err != nil {
			log.Error().Err(err).Msg("滚动鼠标失败")
			return
		}
	case "hover": // 悬停
		err := element.Hover(playwright.LocatorHoverOptions{Timeout: playwright.Float(locator.Timeout)})
		if err != nil {
			log.Error().Err(err).Msg("悬停失败")
			return
		}
	case "right_click": // 右击
		err := element.Click(playwright.LocatorClickOptions{Button: playwright.MouseButtonRight,
			Timeout: playwright.Float(locator.Timeout)})
		if err != nil {
			log.Error().Err(err).Msg("右击失败")
			return
		}
	case "double_click": // 双击
		err := element.Dblclick(playwright.LocatorDblclickOptions{Timeout: playwright.Float(locator.Timeout)})
		if err != nil {
			log.Error().Err(err).Msg("双击失败")
			return
		}
	case "click": // 按键
		err := element.Click(playwright.LocatorClickOptions{Timeout: playwright.Float(locator.Timeout)})
		if err != nil {
			log.Error().Err(err).Msg("点击失败")
			return
		}
	}

}

func inputAction(element playwright.Locator, locator models.Locator, timeout *float64) {
	element.Fill(locator.Operation.Action.Input, playwright.LocatorFillOptions{Timeout: timeout})
}

func enterAction(element playwright.Locator, locator models.Locator, timeout *float64) {
	element.Press(locator.Operation.Action.Interactive, playwright.LocatorPressOptions{Timeout: timeout})
}

// ... [其他动作的函数实现] ...
