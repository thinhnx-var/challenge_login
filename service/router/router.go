package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhnx-var/challenge_login/service/handlers"
)

type Routers struct {
	Pattern  string
	Method   string
	HandlerF func(*gin.Context)
	Name     string
}

// define the route list
var routerList = []Routers{
	{
		Pattern:  "/challenge",
		Method:   "GET",
		HandlerF: handlers.GetChallengeHandler, // should be handlers
		Name:     "Get Challenge API",
	},
}

func NewHandler() http.Handler {
	feature := gin.New()
	gin.SetMode(gin.DebugMode)
	feature.SetTrustedProxies(nil)
	featureGrpV1 := feature.Group("/v1")
	// featuresGroup.Use(middleware.VerifyToken())

	for _, router := range routerList {
		switch router.Method {
		case http.MethodGet:
			{
				featureGrpV1.GET(router.Pattern, router.HandlerF)
			}
		case http.MethodPost:
			{
				featureGrpV1.POST(router.Pattern, router.HandlerF)
			}
		case http.MethodPut:
			{
				featureGrpV1.PUT(router.Pattern, router.HandlerF)
			}
		case http.MethodDelete:
			{
				featureGrpV1.DELETE(router.Pattern, router.HandlerF)
			}
		}
	}
	return feature
}
