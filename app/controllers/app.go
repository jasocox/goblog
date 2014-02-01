package controllers

import "github.com/robfig/revel"

type Application struct {
  *revel.Controller
}

func (a Application) Index() revel.Result {
  return a.Render()
}
