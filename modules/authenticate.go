package modules

import (
	. "checklist/configs"
	. "checklist/dao"
	. "checklist/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mongoConfig = MongoConfig{}
var jwtConfig = JWTConfig{}
var dao = TrackListDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	mongoConfig.Read()
	jwtConfig.Read()

	dao.Server = mongoConfig.Server
	dao.Database = mongoConfig.Database
	dao.Connect()
}

func VerifyGoogleTokenID(token string) (bool, GoogleAccountResponse) {
	var profile GoogleAccountResponse
	response, err := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=" + token)
	if err != nil {
		return false, profile
	} else {
		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&profile)
		if response.StatusCode != 200 || err != nil {
			return false, profile
		}
	}
	return true, profile
}

func IsUserExists(email string) (bool, string) {
	profile, err := dao.FindUserByEmail(email)
	if err != nil {
		return false, ""
	}
	return true, profile.ID.Hex()
}

func CreateNewUser(profile UserAccount) (bool, string) {
	userAccount, err := dao.InsertNewUser(profile)
	if err != nil {
		return false, ""
	}
	return true, userAccount.ID.Hex()
}

func ConverToUserAccount(profile GoogleAccountResponse) UserAccount {
	userAccount := UserAccount{Name: profile.Name, FamilyName: profile.FamilyName, GivenName: profile.GivenName, Email: profile.Email, Picture: profile.Picture}
	return userAccount
}

type jwtCustomClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func CreateSessionToken(id string) UserSession {
	var session UserSession

	// Set custom claims
	claims := &jwtCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(jwtConfig.EncryptionKey))
	if err != nil {
		return session
	}

	session.SessionToken = tokenString

	return session
}
