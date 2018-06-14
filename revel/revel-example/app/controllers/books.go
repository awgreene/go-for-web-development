package controllers

import "github.com/revel/revel"

type Books struct {
    *revel.Controller
}

func (c Books) Index() revel.Result {
	return c.RenderText("Getting all books\n")
}

func (c Books) Get(id string) revel.Result {
	return c.RenderText("Getting book with ID " + id + "\n")
}

func (c Books) Delete(id string) revel.Result {
	return c.RenderText("Deleting book with ID " + id + "\n")
}
