package api

import (
	db "at01/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Đây là nơi chúng tôi triển khai máy chủ API HTTP của mình.
type Server struct {
	store  *db.Store   // Tương tác với cơ sở dữ liệu khi xử lý các yêu cầu API từ máy khách.
	router *gin.Engine // Gửi từng yêu cầu API đến trình xử lý chính xác để xử lý.
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/category", server.createCategory)
	router.GET("/category/:id", server.getCategoryById)
	router.GET("/category", server.getListCategory)
	// router.GET("/category", server.getAllCategory)

	router.POST("/product", server.createProduct)
	router.GET("/product/:id", server.getProductById)
	router.GET("/product", server.getListProduct)

	router.POST("/singup", server.createAccount)
	router.POST("/singin", server.signIn)

	server.router = router
	return server
}
