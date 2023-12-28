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
	n := 0
	for _, step := range steps {
		// 执行定位操作，如果有的话
		if step.Locator != nil {
			locators.PerformActionBasedOnLocator(page, *step.Locator)
		}

		// 执行断言检查，如果有的话
		if step.Assert != nil {
			// 断言计数器,每次执行断言时加1
			n += 1
			testResult := assertions.AssertLocator(page, *step.Assert, n)
			if testResult.FailMessage != "" {
				if !step.Assert.Continue {
					log.Info().Msgf("断言失败,断言后不继续：%v", step.Assert.Continue)
					return
				}
				log.Info().Msgf("断言失败,断言后继续：%v", step.Assert.Continue)
			} else {
				log.Info().Msgf("断言成功，继续后续步骤")
			}
		}
	}
}
