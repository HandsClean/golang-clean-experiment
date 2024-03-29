package main

import (
	// "database/sql"
	// "fmt"
	// "log"
	// "net/url"
	// "os"
	// "time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	//"github.com/spf13/viper"
	// _articleHttpDeliver "github.com/bxcodec/go-clean-arch/article/delivery/http"
	// _articleRepo "github.com/bxcodec/go-clean-arch/article/repository"
	// _articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	// _authorRepo "github.com/bxcodec/go-clean-arch/author/repository"
	//_orderRepo
	//"github.com/bxcodec/go-clean-arch/middleware"
)

// func init() {
// 	viper.SetConfigFile(`config.json`)
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}

// 	if viper.GetBool(`debug`) {
// 		fmt.Println("Service RUN on DEBUG mode")
// 	}
// }
func main() {
	// dbHost := viper.GetString(`database.host`)
	// dbPort := viper.GetString(`database.port`)
	// dbUser := viper.GetString(`database.user`)
	// dbPass := viper.GetString(`database.pass`)
	// dbName := viper.GetString(`database.name`)
	// connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	// val := url.Values{}
	// val.Add("parseTime", "1")
	// val.Add("loc", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	// dbConn, err := sql.Open(`mysql`, dsn)
	// if err != nil && viper.GetBool("debug") {
	// 	fmt.Println(err)
	// }
	// err = dbConn.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// 	os.Exit(1)
	// }

	// defer func() {
	// 	err := dbConn.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// e := echo.New()
	// middL := middleware.InitMiddleware()
	// e.Use(middL.CORS)
	// authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	// ar := _articleRepo.NewMysqlArticleRepository(dbConn)
	// //orderRepo
	// //may be ItemRepo

	// timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	// au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	// _articleHttpDeliver.NewArticleHandler(e, au)

	// log.Fatal(e.Start(viper.GetString("server.address")))
	e := echo.New()

	// // Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// // Routes
	// e.POST("/users", createUser)
	// e.GET("/users/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
