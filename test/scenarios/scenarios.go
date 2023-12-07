package scenarios

import (
	"github.com/ZeroBugHero/web_ui_go/launch"
	"github.com/ZeroBugHero/web_ui_go/models"
	"github.com/ZeroBugHero/web_ui_go/test/assertions"
	"github.com/ZeroBugHero/web_ui_go/test/locators"
	"github.com/playwright-community/playwright-go"
	"github.com/rs/zerolog/log"
)

func Run(steps []models.TestStep, uri string) {
	browser := launch.Browser(launch.GlobalConfig.Browser)
	page, err := browser.NewPage(
		playwright.BrowserNewPageOptions{
			BaseURL: playwright.String(launch.GlobalConfig.BaseUrl),
		})
	if err != nil {
		log.Error().Err(err).Msg("创建新页面失败")
		return
	}
	defer page.Close()
	defer launch.ClosePlaywright()

	page.Goto(uri)

	for _, step := range steps {
		// 执行定位操作，如果有的话
		if step.Locator != nil {
			locators.PerformActionBasedOnLocator(page, *step.Locator)
		}

		// 执行断言检查，如果有的话
		if step.Assert != nil {
			testResult := assertions.AssertLocator(page, *step.Assert)
			if testResult.FailMessage != "" {
				log.Error().Msgf("断言失败：%s", testResult.FailMessage)
				if !step.Assert.Continue {
					return
				}
			}
		}
	}
}
