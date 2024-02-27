package http

import (
	"github.com/stdyum/api-common/hc"
)

type Routes interface {
	ConfigureRoutes() *hc.Engine
}
