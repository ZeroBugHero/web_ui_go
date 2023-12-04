package playwright

import (
	"encoding/json"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type TestResult struct {
	TestName string `json:"testName"`
	Passed   bool   `json:"passed"`
	ErrorMsg string `json:"errorMsg,omitempty"`
}

func TestLocatorAllTextContents(t *testing.T) {
	// 初始化 Playwright 和浏览器
	err := playwright.Install()
	pw, err := playwright.Run()
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)})

	page, err := browser.NewPage()

	// 测试逻辑
	_, err = page.Goto("https://www.baidu.com")

	//err = page.Locator("#kw").Fill("百度", playwright.LocatorFillOptions{Timeout: playwright.Float(500)})
	title, err := page.Title()
	fmt.Println("title:", title)
	fmt.Println("err:", err)
	require.NoError(t, err)
	//fmt.Println(innerHTML)
	//page.Pause()

	// 收集测试结果
	result := TestResult{
		TestName: "TestLocatorAllTextContents",
		Passed:   err == nil,
	}

	if err != nil {
		result.ErrorMsg = err.Error()
	}

	// 将结果保存为 JSON
	saveTestResultAsJSON(result, "test_results.json")
	browser.Close()

}

func saveTestResultAsJSON(result TestResult, filename string) {
	jsonData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}
