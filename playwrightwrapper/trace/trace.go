package trace

import (
	"github.com/ZeroBugHero/web_ui_go/launch"
	"github.com/ZeroBugHero/web_ui_go/models"
	"github.com/playwright-community/playwright-go"
	"github.com/rs/zerolog/log"
)

// TraceTestStep 跟踪
func TraceTestStep(page playwright.Page, assert models.Assert) {

	context := page.Context()
	tracing := context.Tracing()
	err := tracing.Start(playwright.TracingStartOptions{
		Screenshots: playwright.Bool(true),
		Snapshots:   playwright.Bool(true),
		Sources:     playwright.Bool(true),
		Title:       playwright.String(assert.Name),
	})
	if err != nil {
		log.Error().Msgf("Error starting trace: %v", err)
		return
	}
	defer func() {
		stopTrace := tracing.Stop(launch.GlobalConfig.Trace.Path)
		if stopTrace != nil {
			log.Error().Msgf("Error stopping trace: %v", stopTrace)
			return
		}
	}()
}
