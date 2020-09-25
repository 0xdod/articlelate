package handler

// import (
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/Kamva/mgm/v3"
// 	"github.com/fibreactive/articlelate/models"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// // Test that a GET request to the home page returns the home page with
// // the HTTP code 200 for an unauthenticated user
// func TestShowIndexPageUnauthenticated(t *testing.T) {
// 	r := getRouter(true)

// 	r.GET("/", testh.ShowIndexPage)

// 	// Create a request to send to the above route
// 	req, _ := http.NewRequest("GET", "/", nil)

// 	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
// 		// Test that the http status code is 200
// 		statusOK := w.Code == http.StatusOK

// 		// Test that the page title is "Home Page"
// 		// You can carry out a lot more detailed tests using libraries that can
// 		// parse and process HTML pages
// 		p, err := ioutil.ReadAll(w.Body)
// 		pageOK := err == nil && strings.Index(string(p), `<h2>ArticleLate.com</h2>`) > 0

// 		return statusOK && pageOK
// 	})
// }

// func TestGetArticleUnAuth(t *testing.T) {
// 	router := getRouter(true)
// 	router.GET("/article/view/:article_id", testh.GetArticle)
// 	req, _ := http.NewRequest("GET", "/article/view/"+testArticles[0].ID.Hex(), nil)

// 	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
// 		// Test that the http status code is 200
// 		statusOK := w.Code == http.StatusOK

// 		// Test that the page title is "Hello test"
// 		// You can carry out a lot more detailed tests using libraries that can
// 		// parse and process HTML pages
// 		p, err := ioutil.ReadAll(w.Body)
// 		pageOK := err == nil && strings.Index(string(p), `<title>Hello test</title>`) > 0

// 		// Test the article returned from the db
// 		var articles []models.Article
// 		err = mgm.CollectionByName("articles").SimpleFind(&articles, bson.M{"_id": testArticles[0].ID})
// 		articleOK := err == nil && articles[0].ID == testArticles[0].ID
// 		return statusOK && pageOK && articleOK
// 	})
// }

// func TestGetArticleUnAuthJSON(t *testing.T) {
// 	router := getRouter(false)
// 	router.GET("/article/view/:article_id", testh.GetArticle)
// 	req, _ := http.NewRequest("GET", "/article/view/"+testArticles[0].ID.Hex(), nil)
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("Content-type", "application/json")

// 	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
// 		// Test that the http status code is 200
// 		statusOK := w.Code == http.StatusOK

// 		json := `"title":"Hello test","content":"Testing 1"`
// 		p, err := ioutil.ReadAll(w.Body)
// 		pageOK := err == nil && strings.Index(string(p), json) > 0

// 		var articles []models.Article
// 		err = mgm.CollectionByName("articles").SimpleFind(&articles, bson.M{"_id": testArticles[0].ID})
// 		articleOK := err == nil && articles[0].ID == testArticles[0].ID
// 		return statusOK && pageOK && articleOK
// 	})
// }

// func TestCreateArticle(t *testing.T) {
// 	router := getRouter(false)
// 	router.POST("/article/create", testh.CreateArticle)
// 	reader := strings.NewReader(`{"title": "test post","content":"Testing 3"}`)
// 	req, _ := http.NewRequest("POST", "/article/create", reader)
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("Content-type", "application/json")

// 	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
// 		// Test that the http status code is 200
// 		statusOK := w.Code == http.StatusSeeOther

// 		//Test models is created and stored in db
// 		var articles []models.Article
// 		err := mgm.CollectionByName("articles").SimpleFind(&articles, bson.M{"content": "Testing 3"})
// 		pageOK := err == nil && articles[0].Title == "test post"
// 		return statusOK && pageOK
// 	})
// }
