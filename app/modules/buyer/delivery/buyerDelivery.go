package delivery

import (
	"net/http"

	"github.com/evrintobing17/ecommence-REST/app/helpers/jsonhttpresponse"
	"github.com/evrintobing17/ecommence-REST/app/helpers/requestvalidationerror"
	"github.com/evrintobing17/ecommence-REST/app/helpers/routehelper"
	"github.com/evrintobing17/ecommence-REST/app/helpers/structsconverter"
	"github.com/evrintobing17/ecommence-REST/app/middlewares/authmiddleware"
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/evrintobing17/ecommence-REST/app/modules/buyer"
	userDTO "github.com/evrintobing17/ecommence-REST/app/modules/buyer/delivery/userDto"
	userUsecase "github.com/evrintobing17/ecommence-REST/app/modules/buyer/usecase"
	"github.com/gin-gonic/gin"
)

type buyerHandler struct {
	buyerUC        buyer.BuyerUsecase
	authMiddleware authmiddleware.AuthMiddleware
}

func NewAuthHTTPHandler(r *gin.Engine, buyerUC buyer.BuyerUsecase, authMiddleware authmiddleware.AuthMiddleware) {
	handlers := buyerHandler{
		buyerUC:        buyerUC,
		authMiddleware: authMiddleware,
	}

	authorized := r.Group("/buyer")
	{
		authorized.POST("/login", handlers.loginBuyer)
		authorized.POST("/register", handlers.registerBuyer)
	}
	users := r.Group("/buyer", handlers.authMiddleware.AuthorizeJWTWithBuyerContext())
	{
		users.DELETE("", handlers.deleteBuyer)
		users.PUT("", handlers.updateBuyer)
	}
}

func (handler *buyerHandler) loginBuyer(c *gin.Context) {

	var loginReq userDTO.ReqLogin

	errBind := c.ShouldBind(&loginReq)
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

	_, jwt, err := handler.buyerUC.Login(loginReq.Email, loginReq.Password)
	if err != nil {

		if err == userUsecase.ErrInvalidCredential {
			jsonhttpresponse.Unauthorized(c, jsonhttpresponse.NewFailedResponse(err.Error()))
			return
		}

		jsonhttpresponse.InternalServerError(c, jsonhttpresponse.NewFailedResponse(err.Error()))
		return
	}

	response := userDTO.ResLogin{
		Jwt: jwt,
	}
	jsonhttpresponse.OK(c, models.APIResponseOptions{
		StatusCode: 200,
		Message:    "success",
		Errors:     err,
		Data:       response,
	})
}
func (handler *buyerHandler) registerBuyer(c *gin.Context) {
	var registerReq userDTO.Register

	errBind := c.ShouldBind(&registerReq)
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

	user, err := handler.buyerUC.Register(registerReq.Name, registerReq.Email, registerReq.Password, registerReq.Address)

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

	response := userDTO.ResRegister{
		Address: user.Address,
		Email:   user.Email,
		ID:      user.ID,
		Name:    user.Name,
	}

	jsonhttpresponse.OK(c, models.APIResponseOptions{
		StatusCode: 200,
		Message:    "success",
		Errors:     err,
		Data:       response,
	})
}

func (handler *buyerHandler) updateBuyer(c *gin.Context) {
	var request userDTO.ReqUpdate
	errBind := c.ShouldBindJSON(&request)
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

		jsonhttpresponse.BadRequest(c, errBind.Error())
		return
	}

	updatedDriverData, err := structsconverter.ToMap(request)
	if err != nil {
		jsonhttpresponse.InternalServerError(c, err.Error())
	}

	//get user ID
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

	updatedDriverData["id"] = userAuth.ID

	updatedUser, err := handler.buyerUC.Update(updatedDriverData)
	if err != nil {
		jsonhttpresponse.BadRequest(c, models.APIResponseOptions{
			StatusCode: http.StatusBadRequest,
			Message:    "failed",
			Errors:     err,
			Data:       nil,
		})
		return
	}

	resp := userDTO.RespUpdate{
		ID:    updatedUser.ID,
		Email: updatedUser.Email,
		Name:  updatedUser.Name,
		Age:   updatedUser.Address,
	}

	jsonhttpresponse.OK(c, resp)
}

func (handler *buyerHandler) deleteBuyer(c *gin.Context) {
	//get user ID
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

	err = handler.buyerUC.DeleteByID(userAuth.ID)
	if err != nil {
		jsonhttpresponse.BadRequest(c, models.APIResponseOptions{
			StatusCode: http.StatusBadRequest,
			Message:    "failed",
			Errors:     err,
			Data:       nil,
		})
	}

	resp := userDTO.DeleteResp{
		Message: "Your account has been succesfully deleted",
	}

	jsonhttpresponse.OK(c, resp)
}
