package launch

import (
	"testing"
)

func TestInitializePlaywright(t *testing.T) {
	// 启动浏览器
	browser := Browser("chromium")
	defer browser.Close() // 确保在测试结束时关闭浏览器

	// 创建新页面
	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("Failed to create new page: %v", err)
	}

	// 访问网站
	_, err = page.Goto("https://www.baidu.com")
	if err != nil {
		t.Fatalf("Failed to navigate: %v", err)
	}
	defer ClosePlaywright() // 测试结束时关闭 Playwright
}
