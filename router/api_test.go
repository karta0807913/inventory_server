package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/karta0807913/inventory_server/model"
	"github.com/karta0807913/inventory_server/server"
	"gorm.io/gorm"
)

func createServer(t *testing.T, dbname string) (*gin.Engine, *gorm.DB) {
	serv := gin.Default()
	os.Remove(dbname)
	db, err := model.SqliteDB(dbname)
	if err != nil {
		t.Fatal("create sqlite error", err)
	}
	err = model.InitDB(db)
	if err != nil {
		t.Fatal("init sqlite error", err)
	}
	pKey, err := server.GenerateKey()
	if err != nil {
		t.Fatal("create pKey error", err)
	}
	jwt, err := server.NewJwtHelper(pKey)
	if err != nil {
		t.Fatal("create jwt error", err)
	}
	storage, err := server.NewGormStorage(db)
	if err != nil {
		t.Fatal("create storage error", err)
	}
	serv.Use(
		server.NewGinSessionFactory(jwt, storage).SessionMiddleware("session"),
	)
	return serv, db
}

func newBody(body interface{}) (io.Reader, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func TestApi(t *testing.T) {
	serv, db := createServer(t, "test.sqlite")
	defer os.Remove("test.sqlite")
	apiRouter := serv.Group("/")
	ApiRouter(RouterConfig{
		Router: apiRouter,
		DB:     db,
	})
	item := model.ItemTable{
		ItemID:   "abc",
		Name:     "HI",
		Date:     "123asc",
		AgeLimit: 3,
		Cost:     10022121212,
		Location: "e124",
		State:    model.ItemState{},
		Note:     "none",
	}
	err := db.Create(&item).Error
	if err != nil {
		t.Fatal("create item error", err)
	}

	getItem := func(itemID string) (*httptest.ResponseRecorder, *model.ItemTable) {
		request := httptest.NewRequest(
			"GET",
			fmt.Sprintf("/item?item_id=%s", item.ItemID),
			nil,
		)
		response := httptest.NewRecorder()
		serv.ServeHTTP(response, request)
		assert.Equal(t, 200, response.Code)
		var data model.ItemTable
		json.NewDecoder(response.Body).Decode(&data)
		return response, &data
	}

	// GET

	response, data := getItem(item.ItemID)
	assert.Equal(t, item, *data)

	//PUT
	item.Location = "e3224"
	body, err := newBody(gin.H{
		"item_id":  item.ItemID,
		"location": item.Location,
	})
	assert.Equal(t, err, nil)
	request := httptest.NewRequest("PUT", "/item", body)
	response = httptest.NewRecorder()
	serv.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)

	response, data = getItem(item.ItemID)
	assert.Equal(t, item, *data)

	//PUT
	item.State.Correct = true
	body, err = newBody(gin.H{
		"item_id": item.ItemID,
		"state":   item.State,
	})
	assert.Equal(t, err, nil)
	request = httptest.NewRequest("PUT", "/item", body)
	response = httptest.NewRecorder()
	serv.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)

	response, data = getItem(item.ItemID)
	assert.Equal(t, item, *data)

	//PUT
	item.Location = "123311"
	item.State.Unlabel = true
	item.State.Discard = true
	body, err = newBody(gin.H{
		"item_id":  item.ItemID,
		"state":    item.State,
		"location": item.Location,
	})
	assert.Equal(t, err, nil)
	request = httptest.NewRequest("PUT", "/item", body)
	response = httptest.NewRecorder()
	serv.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)

	response, data = getItem(item.ItemID)
	assert.Equal(t, item, *data)
}
