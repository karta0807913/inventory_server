package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/karta0807913/inventory_server/model"
	"github.com/karta0807913/inventory_server/router"
	"github.com/karta0807913/inventory_server/server"
)

func main() {
	db, err := model.SqliteDB(SqliteName)
	if err != nil {
		log.Fatal("start db error", err)
	}

	model.InitDB(db)

	serv := gin.Default()
	_, err = os.Stat(PrivateKeyPath)
	if os.IsNotExist(err) {
		pKey, err := server.GenerateKey()
		if err != nil {
			log.Fatal("Generate key error", err)
		}
		server.SavePEMKey(PrivateKeyPath, pKey)
	} else {
		log.Fatal("read private key file error", err)
	}
	jwt, err := server.NewJwtHelper(PrivateKeyPath)
	if err != nil {
		log.Fatal("read private key file error", err)
	}
	storage, err := server.NewGormStorage(db)
	if err != nil {
		log.Fatal("create storage error", err)
	}
	sessionFactory := server.NewGinSessionFactory(jwt, storage)
	serv.Use(sessionFactory.SessionMiddleware("session"))
	router.InitRouter(router.RouterConfig{
		Router: serv.Group("/"),
		DB:     db,
	})

	serv.Run("0.0.0.0:4000")
}
