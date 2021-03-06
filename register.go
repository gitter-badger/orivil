package orivil

import (
	"gopkg.in/orivil/event.v0"
	"gopkg.in/orivil/middle.v0"
	"gopkg.in/orivil/router.v0"
	"gopkg.in/orivil/service.v0"
)

// every bundle register should implement Register interface
type Register interface {
	RegisterRoute(c *router.Container)

	RegisterService(c *service.Container)

	RegisterMiddle(c *middle.Container)

	Boot(c *service.Container)
}

type MiddlewareConfigure interface {
	SetMiddle(bag *middle.Bag)
}

type ServerEventListener interface {
	AddServerListener(d *event.Dispatcher)
}
