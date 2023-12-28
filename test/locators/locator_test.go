package locators

import (
	"github.com/ZeroBugHero/web_ui_go/launch"
	"github.com/ZeroBugHero/web_ui_go/models"
	"testing"
)

func TestPerformActionBasedOnLocator(t *testing.T) {

	operation := models.Operation{Action: models.Action{
		Interactive: "click",
		Coordinates: models.Coordinates{},
	}}
	locator := models.Locator{
		Name:      "名称",
		Type:      "test-id",
		Exact:     false,
		Values:    []string{"my-button"},
		Operation: operation,
		Timeout:   0,
	}

	browser := launch.Browser("chromium")
	page, _ := browser.NewPage()
	page.Goto("http://localhost:8080/")
	PerformActionBasedOnLocator(page, locator)

}
