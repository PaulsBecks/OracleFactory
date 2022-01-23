package main

import (
	"fmt"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/docs"
	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/routes"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin" // swagger embed files
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func auth(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, isvalid := t.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", t.Header["alg"])
		}
		return []byte(utils.JWT_SECRET), nil
	})
	if err == nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["id"].(float64))
		var user models.User
		db.First(&user, userID)
		ctx.Set("user", user)
	} else {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}

// @title           Subscription Factory API
// @version         2.0
// @description     This is the Subscription Factory server.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	app := gin.Default()
	middleware(app)
	models.InitDB()

	docs.SwaggerInfo.BasePath = "/"
	app.POST("/users/login", routes.Login)
	app.POST("/users/signup", routes.Register)

	app.POST("/outboundSubscriptions/:outboundSubscriptionId/events", routes.PostOutboundSubscriptionEvent)
	app.POST("/inboundSubscriptions/:inboundSubscriptionID/events", routes.PostInboundSubscriptionEvent)
	app.POST("/webServiceProviders/:webServiceProviderID/events", routes.HandleWebServiceProviderEvent)
	app.POST("/smartContractProviders/:smartContractProviderID/events", routes.HandleSmartContractEvent)

	authorized := app.Group("/", auth)
	{
		authorized.GET("/outboundSubscriptions", routes.GetOutboundSubscriptions)
		authorized.GET("/outboundSubscriptions/:outboundSubscriptionId", routes.GetOutboundSubscription)
		authorized.PUT("/outboundSubscriptions/:outboundSubscriptionId", routes.UpdateOutboundSubscription)
		authorized.DELETE("/outboundSubscriptions/:outboundSubscriptionId", routes.DeleteOutboundSubscription)
		authorized.POST("/outboundSubscriptions/:outboundSubscriptionId/start", routes.StartOutboundSubscription)
		authorized.POST("/outboundSubscriptions/:outboundSubscriptionId/stop", routes.StopOutboundSubscription)
		authorized.POST("/outboundSubscriptions", routes.PostOutboundSubscription)

		authorized.GET("/smartContractProviders", routes.GetSmartContractProviders)
		authorized.POST("/smartContractProviders", routes.PostSmartContractProvider)
		authorized.GET("/smartContractProviders/:smartContractProviderID", routes.GetSmartContractProvider)
		authorized.POST("/smartContractProviders/:smartContractProviderID/eventParameters", routes.PostOutboundEventParameters)

		authorized.GET("/inboundSubscriptions/:inboundSubscriptionId", routes.GetInboundSubscription)
		authorized.PUT("/inboundSubscriptions/:inboundSubscriptionId", routes.UpdateInboundSubscription)
		authorized.GET("/smartContractConsumers/:smartContractConsumerID", routes.GetSmartContractConsumer)
		authorized.GET("/inboundSubscriptions", routes.GetInboundSubscriptions)
		authorized.POST("/inboundSubscriptions/:inboundSubscriptionID/start", routes.StartInboundSubscription)
		authorized.POST("/inboundSubscriptions/:inboundSubscriptionID/stop", routes.StopInboundSubscription)
		authorized.POST("/inboundSubscriptions", routes.PostInboundSubscription)

		authorized.POST("/smartContractConsumers/:smartContractConsumerID/inboundSubscriptions", routes.PostInboundSubscription)
		authorized.POST("/smartContractConsumers/:smartContractConsumerID/eventParameters", routes.PostInboundEventParameters)
		authorized.GET("/smartContractConsumers", routes.GetSmartContractConsumers)
		authorized.POST("/smartContractConsumers", routes.PostSmartContractConsumer)

		authorized.GET("/filters", routes.GetFilters)

		authorized.GET("/subscriptions/:subscriptionID/parameterFilters", routes.GetSubscriptionParameterFilters)
		authorized.POST("/subscriptions/:subscriptionID/parameterFilters", routes.PostSubscriptionParameterFilters)
		authorized.DELETE("/subscriptions/:subscriptionID/parameterFilters/:parameterFilterID", routes.DeleteSubscriptionParameterFilter)

		authorized.GET("/webServiceProviders", routes.GetWebServiceProviders)
		authorized.GET("/webServiceProviders/:webServiceProviderID", routes.GetWebServiceProvider)
		authorized.POST("/webServiceProviders", routes.PostWebServiceProvider)

		authorized.GET("/webServiceConsumers", routes.GetWebServiceConsumers)
		authorized.GET("/webServiceConsumers/:webServiceConsumerID", routes.GetWebServiceConsumer)
		authorized.POST("/webServiceConsumers", routes.PostWebServiceConsumer)

		authorized.GET("/providerConsumers/:providerConsumerID/eventParameters", routes.GetProviderConsumerEventParameters)

		authorized.GET("/user", routes.GetCurrentUserDetail)
		authorized.PUT("/user", routes.UpdateCurrentUser)
	}
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.Run()
}
