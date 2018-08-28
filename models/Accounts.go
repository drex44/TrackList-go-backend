package models

import "gopkg.in/mgo.v2/bson"

type UserAccount struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	FamilyName string        `bson:"familyName" json:"familyName"`
	GivenName  string        `bson:"givenName" json:"givenName"`
	Name       string        `bson:"name" json:"name"`
	Email      string        `bson:"email" json:"email"`
	Picture    string        `bson:"picture" json:"picture"`
	Lists      []CList       `bson:"lists" json:"lists"`
}

type GoogleAccountResponse struct {
	Name          string `bson:"name" json:"name"`
	GivenName     string `bson:"givenname" json:"givenname"`
	FamilyName    string `bson:"familyname" json:"familyname"`
	Email         string `bson:"email" json:"email"`
	EmailVerified string `bson:"emailverified" json:"emailverified"`
	Picture       string `bson:"picture" json:"picture"`
	Locale        string `bson:"locale" json:"locale"`
	AtHash        string `bson:"athash" json:"athash"`
	Exp           string `bson:"exp" json:"exp"`
	Iss           string `bson:"iss" json:"iss"`
	Jti           string `bson:"jti" json:"jti"`
	Iat           string `bson:"iat" json:"iat"`
	Azp           string `bson:"azp" json:"azp"`
	Aud           string `bson:"aud" json:"aud"`
	Sub           string `bson:"sub" json:"sub"`
	Alg           string `bson:"alg" json:"alg"`
	Kid           string `bson:"kid" json:"kid"`
}

type UserSession struct {
	TokenId      string `bson:"tokenId" json:"tokenId"`
	SessionToken string `bson:"sessionToken" json:"sessionToken"`
}
