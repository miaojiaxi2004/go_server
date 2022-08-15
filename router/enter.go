package router

import (
	"github.com/miaojiaxi2004/go_server/router/example"
	"github.com/miaojiaxi2004/go_server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
