package handler

//тут будет реализована логика работы с http запросами
import (
	"github.com/darabul/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
)

//штука, которая позволяет обрабатывать запросы на go в виде функций
type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// инициализируем маршруты, к которым будем обращаться с помощью http методов
func (h *Handler) InitRoutes() *gin.Engine {
	//создаем маршрут - это имя, которое отсылает API к определенным endpoints
	//маршрут - это URL, к которому можно обратиться разными HTTP методами
	router := gin.New()
	//создаем endpoints (метод запроса(GET, POST, PUT/PATCH, DELETE) + url)
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SingUp)
		auth.POST("/sing-in", h.SingIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
