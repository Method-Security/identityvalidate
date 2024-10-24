package azure

import (
	"context"
	"time"

	identityvalidate "github.com/Method-Security/identityvalidate/generated/go"
	azure "github.com/Method-Security/identityvalidate/internal/portal/azure"
)

type Module interface {
	ModuleRun(config *identityvalidate.PortalConfig) (*identityvalidate.Trigger, []string)
}

type Engine struct {
	Library Module
	Config  *identityvalidate.PortalConfig
	Modules map[identityvalidate.PortalType]map[identityvalidate.ModuleName]Module
}

func NewEngine(config *identityvalidate.PortalConfig) *Engine {
	return &Engine{
		Config: config,
		Modules: map[identityvalidate.PortalType]map[identityvalidate.ModuleName]Module{
			identityvalidate.PortalTypeAzure: {
				identityvalidate.ModuleNameOwaLogin: &azure.OWALibrary{},
			},
		},
	}
}

func (e *Engine) Run(ctx context.Context) (*identityvalidate.Trigger, []string) {
	return e.Library.ModuleRun(e.Config)
}

func (e *Engine) Launch(ctx context.Context) (*identityvalidate.PortalReport, error) {
	resources := identityvalidate.PortalReport{PortalType: e.Config.PortalType, ModuleName: e.Config.ModuleName}
	errors := []string{}

	e.Library = e.Modules[e.Config.PortalType][e.Config.ModuleName]

	if e.Config.Attempts == 1 {
		e.Config.Interval = 0
	}

	var triggers []*identityvalidate.Trigger
	for i := 0; i < e.Config.Attempts; i++ {
		trigger, errs := e.Run(ctx)
		if trigger != nil {
			triggers = append(triggers, trigger)
		}
		errors = append(errors, errs...)

		time.Sleep(time.Duration(e.Config.Interval) * time.Second)
	}

	resources.Triggers = triggers
	resources.Errors = errors
	return &resources, nil
}
