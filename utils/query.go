package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetLimit(ctx *gin.Context) *int64 {
	limitParam := ctx.DefaultQuery("limit", "10")
	limit, _ := strconv.ParseInt(limitParam, 0, 64)
	return &limit
}

func GetSkip(ctx *gin.Context) *int64 {
	skipParam := ctx.DefaultQuery("skip", "0")
	skip, _ := strconv.ParseInt(skipParam, 0, 64)
	return &skip
}