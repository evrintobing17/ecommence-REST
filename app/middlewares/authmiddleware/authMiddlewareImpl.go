package authmiddleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/evrintobing17/ecommence-REST/app/helpers/jsonhttpresponse"
	"github.com/evrintobing17/ecommence-REST/app/helpers/jwthelper"
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/evrintobing17/ecommence-REST/app/modules/buyer"
	"github.com/evrintobing17/ecommence-REST/app/modules/seller"

	"github.com/gin-gonic/gin"
)

var (
	AdminAccess        = "admin"
	AdminAndUserAccess = "admin|user"

	ErrInvalidToken          = errors.New("invalid token")
	ErrUserContextNotSet     = errors.New("user context is empty. Use AuthorizeJWTWithUserContext instead")
	ErrInvalidResourceAccess = errors.New("this user has no rights to access this resource")
)

type authMiddleware struct {
	buyerService  buyer.BuyerRepository
	sellerService seller.SellerRepository
}

func NewAuthMiddleware(buyerService buyer.BuyerRepository, sellerService seller.SellerRepository) AuthMiddleware {
	return &authMiddleware{sellerService: sellerService, buyerService: buyerService}
}

// AuthorizeJWTWithUserContext - Authorize JWT with User Context (Need to look up for user in DB in every request)
func (auth *authMiddleware) AuthorizeJWTWithBuyerContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user interface{}
		bearerToken := c.GetHeader("Authorization")

		//Get User Claims
		if bearerToken == "" {
			jsonhttpresponse.Unauthorized(c, models.APIResponseOptions{
				StatusCode: http.StatusUnauthorized,
				Message:    ErrInvalidToken.Error(),
				Data:       nil,
			})
			c.Abort()
			return
		}

		//Extract JWT Token from Bearer
		jwtTokenSplit := strings.Split(bearerToken, "Bearer ")
		if jwtTokenSplit[1] == "" {
			jsonhttpresponse.Unauthorized(c, models.APIResponseOptions{
				StatusCode: http.StatusUnauthorized,
				Message:    ErrInvalidToken.Error(),
				Data:       nil,
			})
			c.Abort()
			return
		}
		jwtToken := jwtTokenSplit[1]

		jwtTokenClaims, err := jwthelper.VerifyTokenWithClaims(jwtToken)
		if err != nil {
			jsonhttpresponse.Unauthorized(c, models.APIResponseOptions{
				StatusCode: http.StatusUnauthorized,
				Message:    ErrInvalidToken.Error(),
				Data:       nil,
			})
			c.Abort()
			return
		}

		user, err = auth.buyerService.GetByID(jwtTokenClaims.Id)
		if err != nil {
			jsonhttpresponse.Unauthorized(c, models.APIResponseOptions{
				StatusCode: http.StatusUnauthorized,
				Message:    err.Error(),
				Data:       nil,
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func (auth *authMiddleware) AuthorizeJWTWithSellerContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user interface{}
		bearerToken := c.GetHeader("Authorization")

		//Get User Claims
		if bearerToken == "" {
			jsonhttpresponse.Unauthorized(c, models.APIResponseOptions{
				StatusCode: http.StatusUnauthorized,
				Message:    ErrInvalidToken.Error(),
				Data:       nil,
			})
			c.Abort()
			return
		}

		//Extract JWT Token from Bearer
		jwtTokenSplit := strings.Split(bearerToken, "Bearer ")
		if jwtTokenSplit[1] == "" {
			jsonhttpresponse.Unauthorized(c, models.APIResponseOptions{
				StatusCode: http.StatusUnauthorized,
				Message:    ErrInvalidToken.Error(),
				Data:       nil,
			})
			c.Abort()
			return
		}
		jwtToken := jwtTokenSplit[1]

		jwtTokenClaims, err := jwthelper.VerifyTokenWithClaims(jwtToken)
		if err != nil {
			jsonhttpresponse.Unauthorized(c, models.APIResponseOptions{
				StatusCode: http.StatusUnauthorized,
				Message:    ErrInvalidToken.Error(),
				Data:       nil,
			})
			c.Abort()
			return
		}

		user, err = auth.sellerService.GetByID(jwtTokenClaims.Id)
		if err != nil {
			jsonhttpresponse.Unauthorized(c, models.APIResponseOptions{
				StatusCode: http.StatusUnauthorized,
				Message:    err.Error(),
				Data:       nil,
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
