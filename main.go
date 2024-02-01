package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhvdt/vcs-week2/config"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//Database connection
	db := config.ConnectDatabase()
	err := db.AutoMigrate(&Product{})
	if err != nil {
		return
	}
	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	//server := &http.Server{
	//	Addr:           ":8888",
	//	Handler:        routes,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//err = server.ListenAndServe()

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
