package routers

import (
	"apigw/src/utils/kafkax"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/ratelimit"
)

var (
	limit ratelimit.Limiter
)

type HttpRoutes struct {
	Url       string
	Port      int
	ServiceID string
	Path      string
	Method    string
}

type Handler struct {
	*kafkax.KafkaClient
	HttpRoutes map[string]HttpRoutes
}

func Init() {

	fmt.Fprintln(os.Stdout, `[GIN-debug] Start Initial Gin Engine`)

	rps := viper.GetInt(`apigw.rateLimit`)
	limit = ratelimit.New(rps)

	pro, con := kafkax.InitKafka()
	defer pro.Close()
	defer con.Close()

	h := &Handler{
		KafkaClient: &kafkax.KafkaClient{
			P: pro,
			C: con,
		},
	}

	h.WatcherProducer()

	h.HttpRoutes = initRoutes()

	r := gin.Default()

	r.Use(leakBucket())

	r.NoRoute(h.Messenger)

	port := `:` + viper.GetString(`apigw.serverPort`)
	if err := r.Run(port); err != nil {
		log.Println(err.Error())
	}

}

func initRoutes() map[string]HttpRoutes {

	httpRoutes := []HttpRoutes{}
	httpRoutesMap := map[string]HttpRoutes{}

	viper.UnmarshalKey(`httpRoutes`, &httpRoutes)
	for _, v := range httpRoutes {
		httpRoutesMap[v.Path] = HttpRoutes{
			Path:      v.Path,
			Method:    v.Method,
			Url:       v.Url,
			Port:      v.Port,
			ServiceID: v.ServiceID,
		}
	}
	return httpRoutesMap
}

func leakBucket() gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		log.Print(color.CyanString("Last Request period : [ %v ] ", now.Sub(prev)))
		prev = now
	}
}
