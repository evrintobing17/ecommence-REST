package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// docs "github.com/evrintobing17/docs"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/evrintobing17/ecommence-REST/app/middlewares/authmiddleware"

	buyerdelivery "github.com/evrintobing17/ecommence-REST/app/modules/buyer/delivery"
	buyerrepository "github.com/evrintobing17/ecommence-REST/app/modules/buyer/repository"
	buyerusecase "github.com/evrintobing17/ecommence-REST/app/modules/buyer/usecase"

	sellerdelivery "github.com/evrintobing17/ecommence-REST/app/modules/seller/delivery"
	sellerrepository "github.com/evrintobing17/ecommence-REST/app/modules/seller/repository"
	sellerusecase "github.com/evrintobing17/ecommence-REST/app/modules/seller/usecase"

	productdelivery "github.com/evrintobing17/ecommence-REST/app/modules/product/delivery"
	productrepository "github.com/evrintobing17/ecommence-REST/app/modules/product/repository"
	productusecase "github.com/evrintobing17/ecommence-REST/app/modules/product/usecase"

	orderdelivery "github.com/evrintobing17/ecommence-REST/app/modules/order/delivery"
	orderrepository "github.com/evrintobing17/ecommence-REST/app/modules/order/repository"
	orderusecase "github.com/evrintobing17/ecommence-REST/app/modules/order/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func Initialize(Dbdriver, DbUser, DbPassword, DbHost, DbName string, DbPort int) (DB *gorm.DB) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", DbHost, DbPort, DbUser, DbName, DbPassword)
		DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Println("Cannot connect to database")
			log.Fatal("Error: ", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", Dbdriver)
		}
	}

	return DB

}

func run() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error opening env, %v", err)
	} else {
		fmt.Println(".env file loaded")
	}
	db_driver := os.Getenv("DB_DRIVER")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	db_port, _ := strconv.Atoi(port)
	DB := Initialize(db_driver, db_user, db_password, db_host, db_name, db_port)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaulTableName string) string {
		return "" + defaulTableName
	}

	// docs.SwaggerInfo.BasePath = ""

	r := gin.New()
	// orderDelivery.NewHttpDelivery(r, orderUC)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//Repository
	buyerRepository := buyerrepository.NewBuyerRepository(DB)
	sellerRepository := sellerrepository.NewSellerRepository(DB)
	productrepository := productrepository.NewProductRepository(DB)
	orderRepository := orderrepository.NewOrderRepository(DB)

	//Usecase
	sellerUC := sellerusecase.NewSellerUsecase(sellerRepository)
	buyerUC := buyerusecase.NewBuyerUsecase(buyerRepository)
	productUC := productusecase.NewProductUsecase(productrepository)
	orderUC := orderusecase.NewOrderUsecase(orderRepository, productrepository)

	//Middleware
	authMiddleware := authmiddleware.NewAuthMiddleware(buyerRepository, sellerRepository)

	//Presenter
	sellerdelivery.NewAuthHTTPHandler(r, sellerUC, authMiddleware)
	buyerdelivery.NewAuthHTTPHandler(r, buyerUC, authMiddleware)
	productdelivery.NewProductHTTPHandler(r, productUC, authMiddleware)
	orderdelivery.NewOrderHTTPHandler(r, orderUC, authMiddleware)

	fmt.Println("Listening to port 8081")
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}

func main() {

	run()
}
