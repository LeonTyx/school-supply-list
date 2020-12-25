package authentication

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"school-supply-list/auth/authorization"
	"school-supply-list/database"
	"strconv"
	"time"
)

var (
	GoogleOauthConfig *oauth2.Config
)

func ConfigOauth() {
	if os.Getenv("ENV") == "DEV" {
		GoogleOauthConfig = &oauth2.Config{
			RedirectURL:  "http://localhost:5000/oauth/v1/callback",
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		}
	} else {
		GoogleOauthConfig = &oauth2.Config{
			RedirectURL:  "https://" + os.Getenv("HOST") + "/oauth/v1/callback",
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		}
	}
}

type Error struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_msg"`
}

//All the routes created by the package nested in
// oauth/v1/*
func Routes(r *gin.RouterGroup, db *database.DB) {
	r.GET("/login", HandleGoogleLogin(db))
	r.GET("/callback", HandleGoogleCallback(db))
	r.GET("/logout", HandleGoogleLogout(db))
	r.GET("/profile", GetProfile(db))
	r.GET("/refresh", RefreshSession(db))
}

func GetSeed() int64 {
	seed := time.Now().UnixNano() // A new random seed (independent from state)
	rand.Seed(seed)
	return seed
}

func HandleGoogleLogin(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		state, err := db.SessionStore.Get(c.Request, "state")
		if err != nil {
			c.AbortWithStatusJSON(500, "Server was unable to connect to session database")
			return
		}

		stateString := strconv.FormatInt(GetSeed(), 10)
		state.Values["state"] = stateString
		err = state.Save(c.Request, c.Writer)

		if err != nil {
			print("Unable to store state data")
			c.AbortWithStatusJSON(500, "Unable to store state data")
		}

		url := GoogleOauthConfig.AuthCodeURL(stateString)
		c.Redirect(301, url)
	}

}

func HandleGoogleCallback(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		stateSession, err := db.SessionStore.Get(c.Request, "state")
		if err != nil {
			c.AbortWithStatusJSON(500, "The server was unable to retrieve session state")
			return
		}

		userData, err := GetUserInfo(c.Request.FormValue("state"), c.Request.FormValue("code"), c.Request, db)
		if err != nil {
			fmt.Println("Error getting content: " + err.Error())
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		stateSession.Options.MaxAge = -1
		_ = stateSession.Save(c.Request, c.Writer)
		// Add a user to user database if they don't exist
		// otherwise replace the previous access token field
		// with the new one

		if !UserExists(userData.Email, db) {
			err = CreateUser(userData, db)
			if err != nil {
				database.CheckDBErr(err.(*pq.Error), c)
				return
			}
		} else {
			ReplaceAccessToken(userData, db)
		}

		// set the user information
		session, err := db.SessionStore.Get(c.Request, "session")
		if err != nil {
			c.AbortWithStatusJSON(500, "Server was unable to connect to session database")
		}

		session.Values["GoogleId"] = userData.GoogleId
		session.Values["Email"] = userData.Email
		session.Values["Name"] = userData.Name
		session.Values["Picture"] = userData.Picture

		err = session.Save(c.Request, c.Writer)
		if err != nil {
			fmt.Print("Unable to store session data")
			c.AbortWithStatusJSON(500, "Unable to store session data")
		}

		c.Redirect(http.StatusPermanentRedirect, "/")
	}
}

func getRoleFromGoogleID(c *gin.Context, db *database.DB, googleID string) (authorization.Role, int, error) {
	var role authorization.Role
	var userID int
	roleRows, err := db.Db.Query(`SELECT role.role_id, role.role_name, role.role_desc, user_id from role 
											INNER JOIN account a on role.role_id = a.role_id
											where a.google_id=$1`, googleID)
	if err != nil {
		return role, userID, err
	}
	for roleRows.Next() {
		err = roleRows.Scan(&role.ID, &role.Name, &role.Desc, &userID)
		if err != nil {
			return role, userID, err
		}
		role.Resources, err = getPolicyFromRoleID(c, role.ID, db)
		if err != nil {
			return role, userID, err
		}
	}

	return role, userID, nil
}

func getPolicyFromRoleID(c *gin.Context, roleID string, db *database.DB) ([]authorization.Resource, error) {
	var resources []authorization.Resource
	resourcesRows, err := db.Db.Query(`SELECT resc.resource_id, resc.resource_name,rrb.can_add,
						rrb.can_delete, rrb.can_edit, rrb.can_view from resource resc
						INNER JOIN role_resource_bridge rrb on resc.resource_id = rrb.resource_id
						WHERE rrb.role_id=$1`, roleID)

	if err != nil {
		return resources, err
	}

	for resourcesRows.Next() {
		var resource authorization.Resource
		err := resourcesRows.Scan(&resource.ResourceID, &resource.Resource, &resource.Policy.CanAdd, &resource.Policy.CanDelete,
			&resource.Policy.CanEdit, &resource.Policy.CanView)
		if err != nil {
			c.AbortWithStatusJSON(500, "The server was unable to retrieve permission")
		}
		resources = append(resources, resource)
	}
	_ = resourcesRows.Close()

	return resources, nil
}

func CreateUser(userData User, db *database.DB) error {
	// Prepare the sql query for later
	insert, err := db.Db.Prepare(`INSERT INTO account (email, access_token, google_id, expires_in, google_picture, name) VALUES ($1, $2, $3, $4, $5, $6)`)
	if err != nil {
		return err
	}

	//Execute the previous sql query using data from the
	// userData struct being passed into the function
	_, err = insert.Exec(userData.Email, userData.AccessToken, userData.GoogleId, userData.ExpiresIn, userData.Picture, userData.Name)

	if err != nil {
		return err
	}
	return nil
}

func UserExists(email string, db *database.DB) bool {
	fmt.Println("Checking if user exist: ", email)

	// Prepare the sql query for later
	rows, err := db.Db.Query("SELECT COUNT(*) as count FROM account WHERE email = $1", email)
	PanicOnErr(err)

	return CheckCount(rows) > 0
}

func CheckCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		PanicOnErr(err)
	}
	return count
}

func ReplaceAccessToken(userData User, db *database.DB) {
	_, err := db.Db.Query("UPDATE account SET access_token=$1, expires_in=$2, google_picture=$3, name=$4 WHERE email = $5",
		userData.AccessToken, userData.ExpiresIn, userData.Picture, userData.Name, userData.Email)
	if err != nil {
		fmt.Println("Unable to update access token", err)
	}
}

type User struct {
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Picture     string    `json:"picture"`
	GoogleId    string    `json:"id"`
	ExpiresIn   time.Time `json:"expires_in"`
	AccessToken string
}

func GetUserInfo(state string, code string, r *http.Request, db *database.DB) (User, error) {
	var userData User
	stateSession, err := db.SessionStore.Get(r, "state")
	if err != nil {
		return userData, err
	}

	//Check if the oauth state google returned matches the one saved

	if state != stateSession.Values["state"] {
		return userData, fmt.Errorf("invalid oauth state")
	}

	token, err := GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return userData, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	//Send access token to google's user api in return for a users data!
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		return userData, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return userData, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	err = json.Unmarshal(contents, &userData)
	if err != nil {
		log.Println(err)
	}

	userData.ExpiresIn = token.Expiry
	userData.AccessToken = token.AccessToken

	return userData, nil
}

func HandleGoogleLogout(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Attempting to expire session")

		session, err := db.SessionStore.Get(c.Request, "session")
		if err != nil {
			c.AbortWithStatusJSON(500, "The server was unable to retrieve this session")
			return
		}

		fmt.Println("current session: ", session)
		fmt.Println("Is session new? ", session.IsNew)

		if session.ID != "" {
			session.Options.MaxAge = -1

			err = session.Save(c.Request, c.Writer)

			if err != nil {
				c.AbortWithStatusJSON(500, "The server was unable to expire this session")
			} else {
				c.JSON(200, `{"successful logout"}`)
			}

		} else {
			c.Redirect(http.StatusTemporaryRedirect, "./")
		}
	}
}

type Profile struct {
	Email   string             `json:"email"`
	Name    string             `json:"name"`
	Picture string             `json:"picture"`
	Role    authorization.Role `json:"role"`
	ID      int                `json:"user_id"`
}

func GetProfile(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := db.SessionStore.Get(c.Request, "session")
		if err != nil {
			c.AbortWithStatusJSON(500, "The server was unable to retrieve this session")
			return
		}

		if session.ID != "" {
			fmt.Println("Getting cookies for profile")
			// get some session values
			Email := session.Values["Email"]
			EmailStr := fmt.Sprintf("%v", Email)
			Name := session.Values["Name"]
			NameStr := fmt.Sprintf("%v", Name)
			PictureUrl := session.Values["Picture"]
			PictureUrlStr := fmt.Sprintf("%v", PictureUrl)
			GoogleID := session.Values["GoogleId"]
			GoogleIDStr := fmt.Sprintf("%v", GoogleID)
			role, userID, err := getRoleFromGoogleID(c, db, GoogleIDStr)
			if err != nil {
				database.CheckDBErr(err.(*pq.Error), c)
				return
			}

			userData := Profile{EmailStr, NameStr, PictureUrlStr, role, userID}

			c.JSON(200, userData)
		} else {
			c.AbortWithStatusJSON(401, "Session not found. Session may be expired or non-existent")
		}
	}
}

func RefreshSession(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := db.SessionStore.Get(c.Request, "session")
		if err != nil {
			c.AbortWithStatusJSON(500, "The server was unable to retrieve this session")
			return
		}

		fmt.Println("Current session: ", session)
		fmt.Println("Is session new? ", session.IsNew)

		if session.ID != "" {
			session.Options.MaxAge = 3600

			err = session.Save(c.Request, c.Writer)
			if err != nil {
				c.AbortWithStatusJSON(500, "The server was unable to refresh this session")
			} else {
				c.JSON(200, "successful refresh")
			}
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "./login")
		}
	}
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
