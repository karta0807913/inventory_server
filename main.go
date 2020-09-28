package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/karta0807913/inventory_server/model"
	"github.com/karta0807913/inventory_server/router"
	"github.com/karta0807913/inventory_server/serverutil"
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
		pKey, err := serverutil.GenerateKey()
		if err != nil {
			log.Fatal("Generate key error", err)
		}
		serverutil.SavePEMKey(PrivateKeyPath, pKey)
	}
	jwt, err := serverutil.NewJwtHelperFromPem(PrivateKeyPath)
	if err != nil {
		log.Fatal("read private key file error", err)
	}
	storage, err := serverutil.NewGormStorage(db)
	if err != nil {
		log.Fatal("create storage error", err)
	}
	sessionFactory := serverutil.NewGinSessionFactory(jwt, storage)
	serv.Use(sessionFactory.SessionMiddleware("session"))
	router.InitRouter(router.RouterConfig{
		Router: serv.Group("/"),
		DB:     db,
	})

	serv.Run("0.0.0.0:4000")
}
