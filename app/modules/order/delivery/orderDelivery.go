package delivery

import (
	"net/http"
	"strconv"

	"github.com/evrintobing17/ecommence-REST/app/helpers/jsonhttpresponse"
	"github.com/evrintobing17/ecommence-REST/app/helpers/requestvalidationerror"
	"github.com/evrintobing17/ecommence-REST/app/helpers/routehelper"
	"github.com/evrintobing17/ecommence-REST/app/middlewares/authmiddleware"
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/evrintobing17/ecommence-REST/app/modules/order"
	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderUC        order.OrderUsecase
	authMiddleware authmiddleware.AuthMiddleware
}

func NewOrderHTTPHandler(r *gin.Engine, orderUC order.OrderUsecase, authMiddleware authmiddleware.AuthMiddleware) {
	handlers := orderHandler{
		orderUC:        orderUC,
		authMiddleware: authMiddleware,
	}

	buyer := r.Group("/order", handlers.authMiddleware.AuthorizeJWTWithBuyerContext())
	{
		buyer.GET("/list", handlers.getAll)
		buyer.POST("", handlers.createOrder)
	}
	seller := r.Group("/orders", handlers.authMiddleware.AuthorizeJWTWithSellerContext())
	{
		seller.PUT("/:id", handlers.acceptOrder)
		seller.GET("/list", handlers.getListOrder)
	}
}

func (handler *orderHandler) getAll(c *gin.Context) {

	userAuth, err := routehelper.GetBuyerFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, models.APIResponseOptions{
			StatusCode: http.StatusBadRequest,
			Message:    "failed",
			Errors:     err,
			Data:       nil,
		})
		return
	}

	getOrder, err := handler.orderUC.GetListOrder(userAuth.ID)
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
		Data:       getOrder,
	})
}

func (handler *orderHandler) createOrder(c *gin.Context) {

	var orderModels models.OrderRequest

	errBind := c.ShouldBind(&orderModels)
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

	userAuth, err := routehelper.GetBuyerFromJWTContext(c)
	if err != nil {
		jsonhttpresponse.BadRequest(c, models.APIResponseOptions{
			StatusCode: http.StatusBadRequest,
			Message:    "failed",
			Errors:     err,
			Data:       nil,
		})
		return
	}

	getOrder, err := handler.orderUC.CreateOrder(userAuth.ID, orderModels.ItemID, orderModels.Quantity, userAuth.Address)
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
		Data:       getOrder,
	})
}

func (handler *orderHandler) acceptOrder(c *gin.Context) {
	var orderModels models.AcceptOrderRequest
	orderModels.ItemID, _ = strconv.Atoi(c.Param("photoId"))
	errBind := c.ShouldBind(&orderModels)
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

	getOrder, err := handler.orderUC.AcceptOrder(userAuth.ID, orderModels.ItemID, userAuth.Address)
	if err != nil {
		jsonhttpresponse.BadRequest(c, models.APIResponseOptions{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Errors:     err,
			Data:       nil,
		})
		return
	}

	jsonhttpresponse.OK(c, models.APIResponseOptions{
		StatusCode: 200,
		Message:    "success",
		Errors:     err,
		Data:       getOrder,
	})
}

func (handler *orderHandler) getListOrder(c *gin.Context) {

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

	getOrder, err := handler.orderUC.GetSellerListOrder(userAuth.ID)
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
		Data:       getOrder,
	})
}
