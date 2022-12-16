package api

import (
	db "at01/db/sqlc"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

// POST
type createProductRequest struct {
	IDCategory  string   `json:"idCategory" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Price       int64    `json:"price" binding:"required"`
	Image       string   `json:"image"`
	ListImage   []string `json:"listImage"`
	Description string   `json:"description"`
	Sold        int64    `json:"sold"`
	Status      int64    `json:"status"`
	Sale        int64    `json:"sale"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProductParams{
		IDCategory:  uuid.Must(uuid.Parse(req.IDCategory)),
		Name:        req.Name,
		Price:       null.IntFrom(req.Price),
		Image:       null.StringFrom(req.Image),
		ListImage:   req.ListImage,
		Description: null.StringFrom(req.Description),
		Sold:        null.IntFrom(req.Sold),
		Status:      null.IntFrom(req.Status),
		Sale:        null.IntFrom(req.Sale),
	}

	product, err := server.store.CreateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// /
// / PAGINATE AND GET ALL
// /
type listProductRequest struct {
	Populate string `form:"populate"`
	PageID   int32  `form:"page" binding:"required,min=1"`
	PageSize int32  `form:"limit" binding:"required,min=5,max=10"`
}

func (server *Server) getListProduct(ctx *gin.Context) {
	var req listProductRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		if req.PageID == 0 && req.PageSize == 0 && req.Populate == "" {
			server.getAllProduct(ctx)
			return
		}
	}

	arg := db.ListProductParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize, // bỏ qua page id đã get
	}

	listProduct, err := server.store.ListProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, listProduct)
}

func (server *Server) getAllProduct(ctx *gin.Context) {

	listProduct, err := server.store.GetAllProduct(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, listProduct)
}

// /
// / GET PARAM :id
// /
type getProductRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getProductById(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	product, err := server.store.GetProductById(ctx, uuid.Must(uuid.Parse(req.ID)))
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, product)
	return
}
