package assertions

import (
	"github.com/ZeroBugHero/web_ui_go/launch"
	"github.com/ZeroBugHero/web_ui_go/models"
	"testing"
)

func TestAssertLocator(t *testing.T) {
	browser := launch.Browser("chromium")
	page, _ := browser.NewPage()
	page.Goto("http://localhost:8080/")
	check := models.Check{
		Type:   "eq",
		Expect: []string{"my-vue-project"},
	}
	locator := models.ElementLocator{
		Values: []string{},
		Index:  -1,
	}
	assert := models.Assert{
		Name:           "aa",
		Type:           "title",
		Exact:          false,
		ElementLocator: locator,
		Check:          check,
		Continue:       false,
		Timeout:        0,
	}

	AssertLocator(page, assert)
}
