package models

import "gopkg.in/mgo.v2/bson"

type UserAccount struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	FamilyName string        `bson:"familyName" json:"familyName"`
	GivenName  string        `bson:"givenName" json:"givenName"`
	Name       string        `bson:"name" json:"name"`
	Email      string        `bson:"email" json:"email"`
	Picture    string        `bson:"picture" json:"picture"`
}

type GoogleAccountResponse struct {
	Name          string `bson:"name" json:"name"`
	GivenName     string `bson:"givenName" json:"given_name"`
	FamilyName    string `bson:"familyName" json:"family_name"`
	Email         string `bson:"email" json:"email"`
	EmailVerified string `bson:"emailVerified" json:"email_verified"`
	Picture       string `bson:"picture" json:"picture"`
	Locale        string `bson:"locale" json:"locale"`
	AtHash        string `bson:"atHash" json:"at_hash"`
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
