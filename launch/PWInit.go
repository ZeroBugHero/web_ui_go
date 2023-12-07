package launch

import (
	"github.com/ZeroBugHero/web_ui_go/configs"
	"github.com/ZeroBugHero/web_ui_go/models"
	"github.com/playwright-community/playwright-go"
	"sync"
)

var (
	globalPlaywright *playwright.Playwright
	once             sync.Once
	GlobalConfig     *models.Config
)

func init() {
	once.Do(func() {
		var err error
		globalPlaywright, err = playwright.Run()
		if err != nil {
			panic(err)
		}
	})
	GlobalConfig = configs.ConfigLoad()
}

func Browser(browserType string) playwright.Browser {
	if globalPlaywright == nil {
		panic("Playwright not initialized")
	}

	var browser playwright.Browser
	var err error

	switch browserType {
	case "chromium":
		browser, err = globalPlaywright.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(GlobalConfig.Headless),
		})
	case "firefox":
		browser, err = globalPlaywright.Firefox.Launch()
	case "webkit":
		browser, err = globalPlaywright.WebKit.Launch()
	default:
		browser, err = globalPlaywright.Chromium.Launch()
	}

	if err != nil {
		panic(err)
	}

	return browser
}

// ClosePlaywright 确保在程序结束时关闭 Playwright
func ClosePlaywright() {
	if globalPlaywright != nil {
		globalPlaywright.Stop()
	}
}
