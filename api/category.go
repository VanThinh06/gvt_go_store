package api

import (
	// db "at01/db/sqlc"

	db "at01/db/sqlc"
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// POST
type createCategorytRequest struct {
	Name     string `json:"name" binding:"required"`
	National string `json:"national"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategorytRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCategoryParams{
		Name:     req.Name,
		National: req.National,
	}

	category, err := server.store.CreateCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, category)
}

// GET
type listCategoryRequest struct {
	PageID   int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"limit" binding:"required,min=5,max=10"`
}

func (server *Server) getListCategory(ctx *gin.Context) {
	var req listCategoryRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		if req.PageID == 0 && req.PageSize == 0 {
			server.getAllCategory(ctx)
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCategoryParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize, // bỏ qua page id đã get
	}

	listCategory, err := server.store.ListCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, listCategory)
}

func (server *Server) getAllCategory(ctx *gin.Context) {

	listCategory, err := server.store.GetAllCategory(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, listCategory)
}

// GET PARAM :id
type getCategoryRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getCategoryById(ctx *gin.Context) {
	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	category, err := server.store.GetCategoryById(ctx, uuid.Must(uuid.Parse(req.ID)))
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, category)
	return
}

// chạy máy chủ HTTP trên đầu vào address để bắt đầu lắng nghe các yêu cầu API.
func (server *Server) Start(address string) error {
	return server.router.Run(":8888")
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
