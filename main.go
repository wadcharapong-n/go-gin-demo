package main

import (
	"context"
	"wadcharapong-n/go-gin-demo/src/controller"
	"wadcharapong-n/go-gin-demo/src/model"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ginLambda *ginadapter.GinLambda

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		ginLambda = ginadapter.New(ginEngine())
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func ginEngine() *gin.Engine {
	app := gin.Default()
	model.ConnectDataBase()

	app.GET("/hello", controller.GetHello)
	app.POST("/user", controller.User)

	return app
}

func main() {
	if gin.Mode() == "release" {
		lambda.Start(lambdaHandler)
	} else {
		app := ginEngine()
		app.Run(":3000")
	}
}
