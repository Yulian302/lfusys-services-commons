package oauth

import (
	"github.com/Yulian302/lfusys-services-commons/oauth/github"
	"github.com/Yulian302/lfusys-services-commons/oauth/google"
)

type OAuthConfig struct {
	*github.GithubConfig
	*google.GoogleConfig
}
