package routers

import (
	"apigw/src/utils/cryptox"
	"apigw/src/utils/kafkax"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Request struct {
	Header    map[string][]string `json:"header"`
	ServiceID string              `json:"service_id"`
	Body      []byte              `json:"body"`
}
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (h *Handler) Messenger(c *gin.Context) {
	request := &Request{}
	crypto := viper.GetString(`apigw.crypto`)
	pathWithoutQuery := c.Request.URL.Path
	request.Header = c.Request.Header.Clone()
	request.ServiceID = h.HttpRoutes[pathWithoutQuery].ServiceID
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}
	request.Body = requestBody

	//encrypt
	reqEncrypt, err := cryptox.Encrypt(request, crypto)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//sent request
	if err := kafkax.Producer([]byte(reqEncrypt), request.ServiceID, 0); err != nil {
		log.Fatalln(err.Error())
	}

	//recive response
	msg, err := kafkax.Consumer(request.ServiceID, 0)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//decrypt
	respDecrypt, err := cryptox.Decrypt(string(msg), crypto)
	if err != nil {
		log.Fatalln(err.Error())
	}

	response := Response{}
	if err := json.Unmarshal(respDecrypt, &response); err != nil {
		log.Fatalln(err.Error())
	}

	c.JSON(http.StatusOK, response)
}

// func (h *Handler) Endpoints(c *gin.Context) {

// 	method := c.Request.Method
// 	pathWithoutQuery := c.Request.URL.Path
// 	url := h.HttpRoutes[pathWithoutQuery].Url
// 	port := h.HttpRoutes[pathWithoutQuery].Port
// 	path := c.Request.URL.String()

// 	if checkRouteAndMethod(h.HttpRoutes, pathWithoutQuery, method) {
// 		fullUrl := fmt.Sprintf(`%s:%v%s`, url, port, path)
// 		log.Println(`[`, method, `] Request URL : `, fullUrl)

// 		hRes, err := doRequest(c, method, fullUrl)
// 		if err != nil {
// 			log.Fatalln(err.Error())
// 		}

// 		res, err := bindResponse(c, hRes)
// 		if err != nil {
// 			log.Fatalln(err.Error())
// 		}

// 		c.JSON(hRes.StatusCode, res)
// 	} else {

// 		fullUrl := fmt.Sprintf(`%s:%v%s`, url, port, path)
// 		log.Println(`[`, method, `] Request URL : `, fullUrl)

// 		c.JSON(http.StatusNotFound, Response{
// 			Code:    http.StatusNotFound,
// 			Message: `Not found path : ` + pathWithoutQuery + ` or wrong method`,
// 		})
// 	}

// }

// func doRequest(c *gin.Context, method string, url string) (*http.Response, error) {

// 	client := http.DefaultClient
// 	c.Request.Header.Add(`Auth`, viper.GetString(`apigw.crypto`))
// 	hReq, err := http.NewRequestWithContext(c, method, url, c.Request.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("http req : %v", err.Error())
// 	}

// 	hRes, err := client.Do(hReq)
// 	if err != nil {
// 		return nil, fmt.Errorf("client do : %v", err.Error())
// 	}

// 	return hRes, nil
// }

// func bindResponse(c *gin.Context, h *http.Response) (*Response, error) {

// 	res := Response{}
// 	defer h.Body.Close()

// 	bodyRes, err := io.ReadAll(h.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("res body : %v", err.Error())
// 	}

// 	if err := json.Unmarshal(bodyRes, &res); err != nil {
// 		return nil, fmt.Errorf("unmarshal : %v", err.Error())
// 	}

// 	return &res, nil

// }

// func checkRouteAndMethod(httpRoutes map[string]HttpRoutes, pathWithoutQuery string, method string) bool {

// 	if value, exist := httpRoutes[pathWithoutQuery]; exist {
// 		if value.Method == method {
// 			return true
// 		}
// 	}
// 	return false
// }
