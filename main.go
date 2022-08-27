package main

import (
	"github.com/gin-gonic/gin"
	book2 "go-exercise/book"
	"go-exercise/db/tables"
	handler "go-exercise/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "shinobi:shinobi@tcp(54.179.86.168:3306)/go_exercise?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&tables.Book{})
	router := gin.Default()

	service := book2.BookService(book2.RepoToDb(db))
	bookHandler := handler.NewBookHandler(service)
	//service.Create(book2.BookRequest{
	//	Title:    "buku from new service",
	//	Price:    "50000",
	//	Discount: 30,
	//})

	//cn := book2.RepoToDb(db)
	//bookCreate := tables.Book{
	//	Title: "buku baru banget",
	//	Price: 50000,
	//	Discount: 50,
	//}
	//b, _ := cn.Create(bookCreate)

	//fmt.Println("find by id ",b)

	/*
	* create data on db
	 */
	//b := tables.Book{}
	//b.Title = "man of power"
	//b.Discount = 20
	//b.Price = 50000
	//err = db.Create(&b).Error
	//
	//if err != nil {
	//	fmt.Println("============== error create a book ==============")
	//	fmt.Println(err)
	//}

	/*
	* retrieve data on db
	 */
	//var b tables.Book
	//err = db.Last(&b).Error
	//if err != nil {
	//	fmt.Println("============= failed to retrieved data ===============")
	//	fmt.Println(err)
	//}
	//fmt.Println("titile of book::: ", b.Title)
	//fmt.Println("book::: ",b)

	/*
	* update data on db
	 */
	//var b tables.Book
	//err = db.Debug().Where("title = ?", "power ranger").First(&b).Error
	//if err != nil {
	//	fmt.Println("========== error when find data ============")
	//	fmt.Println(err)
	//}
	//
	//b.Title = "power ranger update"
	//err = db.Save(&b).Error
	//if err != nil {
	//	fmt.Println("========== error when update data ============")
	//	fmt.Println(err)
	//}
	//fmt.Println("title ", b.Title)
	//fmt.Println("object book ", b)

	/*
	* delete on db
	 */
	//var b tables.Book
	//err = db.Debug().Where("id = ? ", 2).First(&b).Error
	//if err != nil {
	//	fmt.Println("=============== failed find the data =================")
	//	fmt.Println(err)
	//}
	//
	//// delete it
	//err = db.Delete(&b).Error
	//if err != nil {
	//	fmt.Println("=============== failed find the data =================")
	//	fmt.Println(err)
	//}

	v1 := router.Group("v1")

	v1.GET("/", bookHandler.BookGetHandler)

	v1.POST("/book", bookHandler.BookHandlerPost)

	router.Run(":80")
}
