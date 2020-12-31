package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"school-supply-list/database"
	"testing"
)

func createRouter() *gin.Engine {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env file.",err)
	}
	db := database.InitDBConnection()
	SStore := database.InitOauthStore()

	dbConnection := &database.DB{Db: db, SessionStore: SStore}
	r.GET("/login", handleGoogleLogin(dbConnection))
	r.GET("/callback", handleGoogleCallback(dbConnection))
	r.GET("/logout", handleGoogleLogout(dbConnection))
	r.GET("/profile", getProfile(dbConnection))
	r.GET("/refresh", refreshSession(dbConnection))

	return r
}

func createSession(r *http.Request, w *httptest.ResponseRecorder, db *database.DB) {
	session, err := db.SessionStore.Get(r, "session")
	if err != nil {
		log.Fatal("Could not create session")
	}
	session.Values["GoogleId"] = "111644517051019423711"
	session.Values["Email"] = "leontyx@gmail.com"
	session.Values["Name"] = "leon"
	session.Values["Picture"] = "img.png"

	err = session.Save(r, w)

	if err != nil {
		fmt.Print("Unable to store session data")
	}
}

func TestProfile(t *testing.T) {
	r := createRouter()
	req, err := http.NewRequest("GET", "/profile", nil)

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := httptest.NewRecorder()

	createSession(req, w, &database.DB{SessionStore: database.InitOauthStore()})

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.FailNow()
	}
	var profile Profile
	contents, _ := ioutil.ReadAll(w.Body)
	err = json.Unmarshal(contents, &profile)

	fmt.Println("email: ", profile.Email)
	fmt.Println("Name: ", profile.Name)
	fmt.Println("Profile picture URL: ", profile.Picture)
	fmt.Println("UUID: ", profile.ID)
}
