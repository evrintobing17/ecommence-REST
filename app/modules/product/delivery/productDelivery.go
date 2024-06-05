package delivery

import (
	"net/http"

	"github.com/evrintobing17/ecommence-REST/app/helpers/jsonhttpresponse"
	"github.com/evrintobing17/ecommence-REST/app/helpers/requestvalidationerror"
	"github.com/evrintobing17/ecommence-REST/app/helpers/routehelper"
	"github.com/evrintobing17/ecommence-REST/app/middlewares/authmiddleware"
	"github.com/evrintobing17/ecommence-REST/app/models"
	userUsecase "github.com/evrintobing17/ecommence-REST/app/modules/buyer/usecase"
	"github.com/evrintobing17/ecommence-REST/app/modules/product"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productUC      product.ProductUsecase
	authMiddleware authmiddleware.AuthMiddleware
}

func NewProductHTTPHandler(r *gin.Engine, productUC product.ProductUsecase, authMiddleware authmiddleware.AuthMiddleware) {
	handlers := productHandler{
		productUC:      productUC,
		authMiddleware: authMiddleware,
	}

	buyer := r.Group("/product/list", handlers.authMiddleware.AuthorizeJWTWithBuyerContext())
	{
		buyer.GET("", handlers.getAll)
	}
	seller := r.Group("/product", handlers.authMiddleware.AuthorizeJWTWithSellerContext())
	{
		seller.GET("", handlers.productBySellerID)
		seller.POST("", handlers.addProduct)
	}
}

func (handler *productHandler) getAll(c *gin.Context) {

	getProduct, err := handler.productUC.GetAll()
	if err != nil {
		jsonhttpresponse.BadRequest(c, models.APIResponseOptions{
			StatusCode: http.StatusBadRequest,
			Message:    "failed",
			Errors:     err,
			Data:       nil,
		})
		return
	}

	jsonhttpresponse.OK(c, models.APIResponseOptions{
		StatusCode: 200,
		Message:    "success",
		Errors:     err,
		Data:       getProduct,
	})
}
func (handler *productHandler) addProduct(c *gin.Context) {
	var productModels models.Product

	errBind := c.ShouldBind(&productModels)
	if errBind != nil {

		validations := requestvalidationerror.GetvalidationError(errBind)
		if len(validations) > 0 {
			jsonhttpresponse.BadRequest(c, models.APIResponseOptions{
				StatusCode:     http.StatusBadRequest,
				Message:        "failed",
				ErrorInterface: validations,
				Data:           nil,
			})
			return
		}

		jsonhttpresponse.BadRequest(c, "")
		return
	}
	userAuth, err := routehelper.GetSellerFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, models.APIResponseOptions{
			StatusCode: http.StatusBadRequest,
			Message:    "failed",
			Errors:     err,
			Data:       nil,
		})
		return
	}

	productModels.SellerID = userAuth.ID

	product, err := handler.productUC.CreateProduct(productModels)
	if err != nil {

		if err == userUsecase.ErrInvalidCredential {
			jsonhttpresponse.Unauthorized(c, err.Error())
			return
		}

		if err == userUsecase.ErrPhoneAlreadyExist ||
			err == userUsecase.ErrEmailAlreadyExist {
			jsonhttpresponse.Conflict(c, err.Error())
		}

		jsonhttpresponse.InternalServerError(c, err.Error())
		return
	}

	response := models.Product{
		ID:          product.ID,
		ProductName: product.ProductName,
		Description: product.Description,
		Price:       product.Price,
		SellerID:    product.SellerID,
	}

	jsonhttpresponse.OK(c, models.APIResponseOptions{
		StatusCode: 200,
		Message:    "success",
		Errors:     err,
		Data:       response,
	})
}

func (handler *productHandler) productBySellerID(c *gin.Context) {
	userAuth, err := routehelper.GetSellerFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, models.APIResponseOptions{
			StatusCode: http.StatusBadRequest,
			Message:    "failed",
			Errors:     err,
			Data:       nil,
		})
		return
	}
	getProduct, err := handler.productUC.GetBySellerID(userAuth.ID)
	if err != nil {
		jsonhttpresponse.BadRequest(c, models.APIResponseOptions{
			StatusCode: http.StatusBadRequest,
			Message:    "failed",
			Errors:     err,
			Data:       nil,
		})
		return
	}

	jsonhttpresponse.OK(c, models.APIResponseOptions{
		StatusCode: 200,
		Message:    "success",
		Errors:     err,
		Data:       getProduct,
	})
}
