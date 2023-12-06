package scenarios

import (
	"testing"
	"time"
	"web_ui_go/models"
)

func TestRun(t *testing.T) {

	operation := models.Operation{Action: models.Action{

		Interactive: "click",
		Coordinates: nil,
	}}
	locator := models.Locator{
		Name:      "单测定位",
		Type:      "test-id",
		Exact:     false,
		Values:    []string{"my-button"},
		Operation: operation,
		Timeout:   0,
	}
	elementLocator := models.ElementLocator{
		Values: []string{"my-button"},
		Index:  0,
	}
	check := models.Check{
		Type:   "eq",
		Expect: []string{"my-button"},
	}

	assert := models.Assert{
		Name:           "断言",
		Type:           "test-id",
		Exact:          false,
		ElementLocator: elementLocator,
		Check:          check,
		Continue:       false,
		Timeout:        0,
	}
	step := models.TestStep{
		StartTime: time.Now(),
		Locator:   &locator,
		Assert:    &assert,
	}

	check1 := models.Check{
		Type:   "eq",
		Expect: []string{"按钮已被点击！"},
	}
	elementLocator1 := models.ElementLocator{
		Values: []string{"my-message"},
		Index:  0,
	}
	assert1 := models.Assert{
		Name:           "断言",
		Type:           "test-id",
		Exact:          false,
		ElementLocator: elementLocator1,
		Check:          check1,
		Continue:       true,
		Timeout:        0,
	}
	step1 := models.TestStep{
		StartTime: time.Now(),
		Assert:    &assert1,
	}
	steps := []models.TestStep{step, step1}

	Run(steps, "/login")

}
