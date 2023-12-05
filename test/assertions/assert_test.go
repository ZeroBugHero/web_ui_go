package assertions

import (
	"testing"
	"web_ui_go/launch"
	"web_ui_go/models"
)

func TestAssertLocator(t *testing.T) {
	browser := launch.Browser("chromium")
	page, _ := browser.NewPage()
	page.Goto("http://localhost:8080/")
	check := models.Check{
		Type:   "eq",
		Expect: "点击我",
	}
	locator := models.ElementLocator{
		Values: []string{"my-button"},
		Index:  -1,
	}
	assert := models.Assert{
		Name:           "aa",
		Type:           "test-id",
		Exact:          false,
		ElementLocator: locator,
		Check:          check,
		Continue:       false,
		Timeout:        0,
	}

	AssertLocator(page, assert)
}
