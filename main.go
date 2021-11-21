package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/huangzhuo492008824/go-gin-example/models"
	"github.com/huangzhuo492008824/go-gin-example/pkg/gredis"
	"github.com/huangzhuo492008824/go-gin-example/pkg/logging"
	"github.com/huangzhuo492008824/go-gin-example/pkg/setting"
	"github.com/huangzhuo492008824/go-gin-example/routers"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath
// @query.collection.format multi

/*
// @securityDefinitions.basic BasicAuth
*/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token

/*
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
*/

// @x-extension-openapi {"example": "value on a json format"}

// func newTransport() *customTransport {
// 	return &customTransport{
// 		originalTransport: http.DefaultTransport,
// 	}
// }

// type customTransport struct {
// 	originalTransport http.RoundTripper
// }

// func (c *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {

// 	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

// 	resp, err := c.originalTransport.RoundTrip(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }

func main() {

	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Server err: %v", err)
		}

	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}

	log.Println("server exiting")
	/*
		router := routers.InitRouter()

		s := &http.Server{
			Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
			Handler:        router,
			ReadTimeout:    setting.ReadTimeout,
			WriteTimeout:   setting.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		}

		s.ListenAndServe()
	*/
}
