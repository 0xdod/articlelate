package handler

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/Kamva/mgm/v3"
// 	"github.com/fibreactive/articlelate/models"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// func TestRegister(t *testing.T) {
// 	router := getRouter(false)
// 	router.POST("/u/register", testh.Register)
// 	reader := strings.NewReader(`{"username": "test user","email":"test@testing.com", "password":"12345abc"}`)
// 	req, _ := http.NewRequest("POST", "/u/register", reader)
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("Content-type", "application/json")

// 	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
// 		// Test that the http status code is 200
// 		statusOK := w.Code == http.StatusSeeOther

// 		//Test models is created and stored in db
// 		var users []models.User
// 		err := mgm.CollectionByName("users").SimpleFind(&users, bson.M{"username": "test user"})
// 		pageOK := err == nil && users[0].Username == "test user"
// 		return statusOK && pageOK
// 	})
// }

// func TestLogin(t *testing.T) {
// 	router := getRouter(false)
// 	router.POST("/u/login", testh.Login)
// 	reader := strings.NewReader(`{"login":"test user","password":"12345abc"}`)
// 	req, _ := http.NewRequest("POST", "/u/login", reader)
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("Content-type", "application/json")

// 	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
// 		statusOK := w.Code == http.StatusSeeOther
// 		return statusOK
// 	})
// }
