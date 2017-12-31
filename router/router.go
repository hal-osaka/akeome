package router

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/images", "./public/images")
	r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/@:id", func(c *gin.Context) {
		c.HTML(200, "post.html", nil)
	})

	apiRouter(r.Group("/api"))

	return r

}
