package delivery

import (
	"net/http"

	"github.com/evrintobing17/ecommence-REST/app/helpers/jsonhttpresponse"
	"github.com/evrintobing17/ecommence-REST/app/helpers/requestvalidationerror"
	"github.com/evrintobing17/ecommence-REST/app/middlewares/authmiddleware"
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/evrintobing17/ecommence-REST/app/modules/seller"
	userDTO "github.com/evrintobing17/ecommence-REST/app/modules/seller/delivery/userDto"
	userUsecase "github.com/evrintobing17/ecommence-REST/app/modules/seller/usecase"
	"github.com/gin-gonic/gin"
)

type sellerHandler struct {
	sellerUC       seller.SellerUsecase
	authMiddleware authmiddleware.AuthMiddleware
}

func NewAuthHTTPHandler(r *gin.Engine, sellerUC seller.SellerUsecase, authMiddleware authmiddleware.AuthMiddleware) {
	handlers := sellerHandler{
		sellerUC:       sellerUC,
		authMiddleware: authMiddleware,
	}

	authorized := r.Group("/seller")
	{
		authorized.POST("/login", handlers.loginSeller)
		authorized.POST("/register", handlers.registerSeller)
	}
}

func (handler *sellerHandler) loginSeller(c *gin.Context) {

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

	_, jwt, err := handler.sellerUC.Login(loginReq.Email, loginReq.Password)
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

func (handler *sellerHandler) registerSeller(c *gin.Context) {
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

	user, err := handler.sellerUC.Register(registerReq.Name, registerReq.Email, registerReq.Password, registerReq.Address)

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
