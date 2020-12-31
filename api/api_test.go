package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"school-supply-list/database"
	"testing"
)

func createRouter() *gin.Engine {
	r := gin.Default()
	if _, err := os.Stat("../projectvars.env"); err == nil {
		err := godotenv.Load("../projectvars.env")
		if err != nil {
			fmt.Println("Error loading environment.env")
		}
		fmt.Println("Current environment:", os.Getenv("ENV"))
	}

	db := database.InitDBConnection()
	SStore := database.InitOauthStore()

	dbConnection := &database.DB{Db: db, SessionStore: SStore}
	Routes(r.Group("api/v1"), dbConnection)

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

func TestCreateSchool(t *testing.T) {
	r := createRouter()
	database.PerformMigrations("file://../database/migrations")
	school := school{
		SchoolName:  "Little Test Elementary",
	}
	schoolJson, err := json.Marshal(school)
	if err != nil {
		fmt.Println("Unable to marshall provides test school into JSON")
		t.Fail()
	}

	req, err := http.NewRequest("PUT", "/api/v1/school", bytes.NewBuffer(schoolJson))

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := httptest.NewRecorder()

	createSession(req, w, &database.DB{SessionStore: database.InitOauthStore()})

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}

	contents, _ := ioutil.ReadAll(w.Body)
	err = json.Unmarshal(contents, &school)
}

func TestGetSchool(t *testing.T) {
	r := createRouter()
	req, err := http.NewRequest("GET", "/api/v1/school", nil)

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := httptest.NewRecorder()

	createSession(req, w, &database.DB{SessionStore: database.InitOauthStore()})

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	var school school
	contents, _ := ioutil.ReadAll(w.Body)
	err = json.Unmarshal(contents, &school)
}

func TestUpdateSchool(t *testing.T) {
	r := createRouter()
	req, err := http.NewRequest("PUT", "/api/v1/school", nil)

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := httptest.NewRecorder()

	createSession(req, w, &database.DB{SessionStore: database.InitOauthStore()})

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	var school school
	contents, _ := ioutil.ReadAll(w.Body)
	err = json.Unmarshal(contents, &school)
}

func TestDeleteSchool(t *testing.T) {
	r := createRouter()
	req, err := http.NewRequest("DELETE", "/api/v1/school", nil)

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := httptest.NewRecorder()

	createSession(req, w, &database.DB{SessionStore: database.InitOauthStore()})

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	var school school
	contents, _ := ioutil.ReadAll(w.Body)
	err = json.Unmarshal(contents, &school)
}
