package handlers

import (
	"github.com/go-chi/render"
	"log"
	"net/http"
	"robin/api/src/utils"
	"strconv"
	"time"
)

func (rd GetTokenResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type jsonEpochTime time.Time

func (t jsonEpochTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

type GetTokenResponse struct {
	Token   string        `json:"token"`
	Expires jsonEpochTime `json:"expires"`
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	claims := utils.JwtClaims{UserId: "cool-new-user"}
	token, tokenString, _ := utils.GetTokenAuth().Encode(claims.ToClaimsMap(time.Hour * 1))

	render.Status(r, http.StatusOK)
	err := render.Render(w, r, GetTokenResponse{Token: tokenString, Expires: jsonEpochTime(token.Expiration())})

	if nil != err {
		log.Println(err)
	}
}
