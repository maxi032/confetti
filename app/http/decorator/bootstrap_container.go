package decorator

import (
	"github.com/confetti-framework/framework/contract/inter"
	"github.com/confetti-framework/framework/foundation"
	"github.com/confetti-framework/framework/foundation/decorator/container_decorator"
)

var bootstraps = []inter.Bootstrap{
	registerProviders{},
	bootProviders{},
}

// Bootstrap ensures that the application is started
func Bootstrap(container *foundation.Container) inter.Container {
	handler := container_decorator.Handler{Bootstraps: bootstraps}

	return handler.BootstrapWith(container)
}
