package controllers

import (
	"github.com/revel/revel"
)

type Home struct {
	*revel.Controller
}

func (C Home) Index() revel.Result {
	return C.Render()
}
