package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoController struct{}

func (c *VideoController) FindAll() {
}

func (c *VideoController) Save(ctx *gin.Context) error {

	return nil
}

func (c *VideoController) Update(ctx *gin.Context) error {
	return nil
}

func (c *VideoController) Delete(ctx *gin.Context) error {
	return nil
}

func (c *VideoController) ShowAll(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", "")
}
