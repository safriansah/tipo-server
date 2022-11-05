package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"tipo-server/app/models"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	conf = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	oauthState = ""
)

/*
InitializeOAuthGoogle Function
*/
func InitializeOAuthGoogle() {
	conf.ClientID = viper.GetString("GOOGLE_CLIENT_ID")
	conf.ClientSecret = viper.GetString("GOOGLE_CLIENT_SECRET")
	oauthState = viper.GetString("GOOGLE_OAUTH_STATE")
	conf.RedirectURL = viper.GetString("GOOGLE_CALLBACK_URL")
}

func GetGoogleOauthState() (state string) {
	return oauthState
}

func GetGoogleLoginUrl() (loginUrl string) {
	loginUrl = conf.AuthCodeURL(oauthState)
	return loginUrl
}

func HandleGoogleCodeExcange(code string) (result *models.GoogleProfile, err error) {
	code, err = url.QueryUnescape(code)
	if err != nil {
		log.Printf("HandleGoogleCodeExcange::url.QueryUnescape::%v", err)
		return nil, err
	}
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("HandleGoogleCodeExcange::conf.Exchange::%v", err)
		return nil, err
	}

	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		log.Printf("Get: " + err.Error() + "\n")
		return nil, err
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		fmt.Printf("FetchCheckTypo::json.NewDecoder::%v", err)
		return nil, err
	}

	result.AccessToken = token.AccessToken
	result.TokenType = token.TokenType
	result.RefreshToken = token.RefreshToken
	result.Expiry = token.Expiry

	return result, nil
}
