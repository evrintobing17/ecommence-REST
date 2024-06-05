package authmiddleware

import "github.com/gin-gonic/gin"

type AuthMiddleware interface {
	AuthorizeJWTWithBuyerContext() gin.HandlerFunc
	AuthorizeJWTWithSellerContext() gin.HandlerFunc
}
