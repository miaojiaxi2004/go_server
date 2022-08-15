package v1

import (
	"github.com/miaojiaxi2004/go_server/api/v1/example"
	"github.com/miaojiaxi2004/go_server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
