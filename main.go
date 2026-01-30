package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.getVacationRouter(r)
	r.run()
}
