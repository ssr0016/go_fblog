package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssr0016/blogbackend/controller"
	"github.com/ssr0016/blogbackend/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticate)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
	app.Get("/api/detailpost/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)

	app.Get("/api/uniquepost", controller.UniquePost)

}
