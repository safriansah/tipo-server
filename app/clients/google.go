package clients

import (
	"log"
	"net/url"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthState = ""
)

/*
InitializeOAuthGoogle Function
*/
func InitializeOAuthGoogle() {
	oauthConfig.ClientID = viper.GetString("GOOGLE_CLIENT_ID")
	oauthConfig.ClientSecret = viper.GetString("GOOGLE_CLIENT_SECRET")
	oauthState = viper.GetString("GOOGLE_OAUTH_STATE_STRING")
	oauthConfig.RedirectURL = viper.GetString("GOOGLE_CALLBACK_URL")
}

func GetGoogleLoginUrl() (loginUrl string, err error) {
	URL, err := url.Parse(oauthConfig.Endpoint.AuthURL)
	if err != nil {
		log.Println("Parse: " + err.Error())
		return loginUrl, err
	}
	log.Println(URL.String())
	parameters := url.Values{}
	parameters.Add("client_id", oauthConfig.ClientID)
	parameters.Add("scope", strings.Join(oauthConfig.Scopes, " "))
	parameters.Add("redirect_uri", oauthConfig.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthState)
	URL.RawQuery = parameters.Encode()
	loginUrl = URL.String()
	return loginUrl, err
}
