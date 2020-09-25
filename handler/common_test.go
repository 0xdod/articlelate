package handler

// import (
// 	"context"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"github.com/Kamva/mgm/v3"
// 	"github.com/fibreactive/articlelate/models"
// 	"github.com/fibreactive/articlelate/service"
// 	"github.com/gin-gonic/gin"
// )

// // set up test db, handler and models
// var d = db.TestNew()
// var testus = service.NewUserStore(d)
// var testas = service.NewArticleStore(d)
// var testh = NewHandler(testus, testas)
// var testArticles = []*models.Article{
// 	models.NewArticle("Hello test", "Testing 1"),
// 	models.NewArticle("Hello test", "Testing 2"),
// }
// var testUsers = []*models.User{
// 	models.NewUser("tester 1", "tester1@testing.com", "12345abc"),
// 	models.NewUser("tester 2", "tester2@testing.com", "abc12345"),
// }

// func setUp() {
// 	//populate articles
// 	for _, a := range testArticles {
// 		_ = mgm.CollectionByName("articles").Create(a)
// 	}
// 	//populate users
// 	for _, u := range testUsers {
// 		_ = mgm.CollectionByName("users").Create(u)
// 	}
// }

// func tearDown() {
// 	//delete articles
// 	for _, a := range testArticles {
// 		_ = mgm.CollectionByName("articles").Delete(a)
// 	}
// 	//delete users
// 	for _, u := range testUsers {
// 		_ = mgm.CollectionByName("users").Delete(u)
// 	}
// }

// // This function is used for setup before executing the test functions
// func TestMain(m *testing.M) {
// 	//Set Gin to Test Mode
// 	gin.SetMode(gin.TestMode)
// 	setUp()
// 	// Run the other tests
// 	code := m.Run()
// 	tearDown()
// 	d.Db.Drop(context.Background())
// 	os.Exit(code)
// }

// // Helper function to create a router during testing
// func getRouter(withTemplates bool) *gin.Engine {
// 	r := gin.Default()
// 	if withTemplates {
// 		r.LoadHTMLGlob("../templates/*")
// 	}
// 	return r
// }

// // Helper function to process a request and test its response
// func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

// 	// Create a response recorder
// 	w := httptest.NewRecorder()

// 	// Create the service and process the above request.
// 	r.ServeHTTP(w, req)

// 	if !f(w) {
// 		t.Fail()
// 	}
// }
