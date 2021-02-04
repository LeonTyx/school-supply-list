package schools

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/antonlindstrom/pgstore"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"school-supply-list/auth/authorization"
	"school-supply-list/database"
	"testing"
)

var r *gin.Engine
var SessionStore *pgstore.PGStore
var db *sql.DB

func init() {
	r = gin.Default()
	if _, err := os.Stat("../../projectvars.env"); err == nil {
		err := godotenv.Load("../../projectvars.env")
		if err != nil {
			fmt.Println("Error loading environment.env")
		}
		fmt.Println("Current environment:", os.Getenv("ENV"))
	}
	database.PerformMigrations("file://../../database/migrations")

	db = database.InitDBConnection()
	SessionStore = database.InitOauthStore()

	dbConnection := &database.DB{Db: db, SessionStore: SessionStore}

	rgroup := r.Group("api/v1")
	rgroup.PUT("/school", authorization.ValidSession(dbConnection),
		authorization.LoadPolicy(dbConnection, "school"),
		authorization.CanCreate(),
		CreateSchool(dbConnection))
	rgroup.GET("/school/:id",
		GetSchool(dbConnection))
	rgroup.GET("/schools",
		GetSchools(dbConnection))
	rgroup.POST("/school/:id", authorization.ValidSession(dbConnection),
		authorization.LoadPolicy(dbConnection, "school"),
		authorization.CanEdit(),
		UpdateSchool(dbConnection))
	rgroup.DELETE("/school/:id", authorization.ValidSession(dbConnection),
		authorization.LoadPolicy(dbConnection, "school"),
		authorization.CanDelete(),
		DeleteSchool(dbConnection))
}

func createTestUser() {
	row := db.QueryRow(`INSERT INTO account (user_id, google_id, email, name, google_picture, expires_in, access_token)  
								VALUES ('00000000-0000-11eb-a029-00ff282e905c', '00000000000000', 'testuser@example.com', 
								'Johnny Test', 'img.png', '2080-12-29 03:56:43.854138' , '0000000') on conflict do nothing returning user_id`)
	if row.Err() != nil {
		log.Fatal("Unable to create test user to be deleted. Error: ", row.Err())
	}
}

func cleanupDatabase() {
	row := db.QueryRow(`DELETE from account where user_id=$1`, "00000000-0000-11eb-a029-00ff282e905c")
	if row.Err() != nil {
		log.Fatal("Unable to remove test user. Error: ", row.Err())
	}

	row = db.QueryRow(`DELETE from role where role_id=-1`)
	if row.Err() != nil {
		log.Fatal("Unable to remove test role. Error: ", row.Err())
	}
}

func createDefaultUser(r *http.Request, w *httptest.ResponseRecorder) {
	createTestUser()

	session, err := SessionStore.Get(r, "session")
	fmt.Println("...............................................................................")
	if err != nil {
		log.Fatal("Could not create session")
	}

	session.Values["GoogleId"] = "00000000000000"
	session.Values["Email"] = "testuser@example.com"
	session.Values["Name"] = "Johnny Test"
	session.Values["Picture"] = "img.png"

	err = session.Save(r, w)

	if err != nil {
		fmt.Print("Unable to store session data")
	}
}

func addValidRole() int {
	row := db.QueryRow(`INSERT INTO role (role_id, role_name, role_desc) 
								VALUES (-1, 'test', 'Temporary test role. Delete if present outside of testing.') ON CONFLICT do nothing returning role_id`)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Fatal("Unable to create test role to be deleted. Error: ", err)
	}
	row = db.QueryRow(`INSERT INTO role_resource_bridge (can_add, can_view, can_edit, can_delete, resource_id, role_id) 
								VALUES (true, true, true, true, $1 , $2) returning rrb_id`, 1, id)
	var rrbID int
	err = row.Scan(&rrbID)
	if err != nil {
		log.Fatal("Unable to add test permissions to be deleted. Error: ", err)
	}

	db.QueryRow(`INSERT INTO user_role_bridge (user_uuid, role_id)
								VALUES ($1 , $2)`, "00000000-0000-11eb-a029-00ff282e905c", id)
	return id
}

func TestCreateSchool(t *testing.T) {
	school := school{
		SchoolName: "Little Test Elementary",
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
	createDefaultUser(req, w)
	addValidRole()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}

	contents, _ := ioutil.ReadAll(w.Body)
	err = json.Unmarshal(contents, &school)

	rowCount := db.QueryRow("SELECT count(*) from school WHERE school_id=$1", school.SchoolID)
	var count int
	err = rowCount.Scan(&count)
	if err != nil {
		t.Fail()
	}

	if count == 0 {
		fmt.Println("School is still present in database")
		t.Fail()
	}

	db.QueryRow("DELETE from school where school_id=$1", school.SchoolID)
	cleanupDatabase()
}

func TestGetSchools(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/schools", nil)

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := httptest.NewRecorder()

	createDefaultUser(req, w)
	addValidRole()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	var school school
	contents, _ := ioutil.ReadAll(w.Body)
	err = json.Unmarshal(contents, &school)
	cleanupDatabase()
}

func TestGetSchool(t *testing.T) {
	row := db.QueryRow("INSERT INTO school (school_name, school_id) VALUES ('Test', default) RETURNING school_id")
	var id string
	err := row.Scan(&id)
	if err != nil{
		t.Fail()
	}
	req, err := http.NewRequest("GET", "/api/v1/school/"+id, nil)

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := httptest.NewRecorder()

	createDefaultUser(req, w)
	addValidRole()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	var school school
	contents, _ := ioutil.ReadAll(w.Body)
	err = json.Unmarshal(contents, &school)
	cleanupDatabase()
}

func TestUpdateSchool(t *testing.T) {
	row := db.QueryRow("INSERT INTO school (school_name, school_id) VALUES ('Test', default) RETURNING school_id")
	var id string
	err := row.Scan(&id)
	if err != nil {
		log.Fatal("Unable to create test school to be updated. Error: ", err)
	}
	school := school{
		SchoolName: "New Little Test Elementary",
	}
	schoolJson, err := json.Marshal(school)
	if err != nil {
		fmt.Println("Unable to marshall provides test school into JSON")
		t.Fail()
	}

	req, err := http.NewRequest("POST", "/api/v1/school/"+id, bytes.NewBuffer(schoolJson))

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := httptest.NewRecorder()

	createDefaultUser(req, w)
	addValidRole()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}

	db.QueryRow("DELETE from school where school_id=$1", id)
	cleanupDatabase()
}

func TestDeleteSchool(t *testing.T) {
	row := db.QueryRow("INSERT INTO school (school_name, school_id) VALUES ('Test', default) RETURNING school_id")
	var id string
	err := row.Scan(&id)
	if err != nil {
		log.Fatal("Unable to create test school to be deleted. Error: ", err)
	}
	req, err := http.NewRequest("DELETE", "/api/v1/school/"+id, nil)

	if err != nil {
		t.Fatalf(err.Error())
	}

	w := httptest.NewRecorder()

	createDefaultUser(req, w)
	addValidRole()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}

	rowCount := db.QueryRow("SELECT count(*) from school WHERE school_id=$1", id)
	var count int
	err = rowCount.Scan(&count)
	if err != nil {
		t.Fail()
	}

	if count > 0 {
		fmt.Println("School is still present in database")
		t.Fail()
	}

	cleanupDatabase()
}
