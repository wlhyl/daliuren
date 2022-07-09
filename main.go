package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/wlhyl/daliuren/daliuren"
	_ "github.com/wlhyl/daliuren/docs"
	"k8s.io/klog"
)

// @title 大六壬排盘 API
// @version 1.0
// @description 大六壬排盘

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @accept json
func main() {
	var port int
	flag.IntVar(&port, "port", 80, "server port.")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/healthz", http.HandlerFunc(healthzHandler))
	mux.Handle("/api/daliuren", http.HandlerFunc(daliurenHandler))
	// mux.Handle("/swagger/index.html", httpSwagger.Handler(
	// 	httpSwagger.URL("http://localhost/swagger/doc.json"), //The url pointing to API definition
	// ))

	mux.Handle("/swagger/", http.HandlerFunc(httpSwagger.WrapHandler))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	// server.ListenAndServe()

	// start webhook server in new routine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			klog.Errorf("Failed to listen and serve server: %v", err)
		}

	}()

	klog.Info("Server started")

	// listening OS shutdown singal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	klog.Infof("Got OS shutdown signal, shutting down server gracefully...")
	server.Shutdown(context.Background())
}

// @description 健康检查
// @Tags healthz
// @summary 健康检查
// @Produce plain
// @accpet plain
// @Success 200 {string} string "ok"
// @Router /healthz [get]
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("ok\n"))
	if err != nil {
		klog.Errorf("Failed check server health: %v", err)
	}
}

// @description 六壬排盘
// @Tags daliuren
// @summary 六壬排盘
// @Produce json
// @accept json
// @Param name body DaLiuRenReust true "request body"
// @Success 200 {object} daliuren.DaliurenShiPan{} "大六壬式盘json"
// @Router /api/daliuren [post]
func daliurenHandler(w http.ResponseWriter, r *http.Request) {
	//记录日志
	var log bytes.Buffer

	// verify the content type is accurate
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		log.WriteString(fmt.Sprintf("Content-Type=%s, expect `application/json`", contentType))
		klog.Errorf(log.String())
		//如果在Apiserver调用此server返回是415，说明APIServer自己传过来的数据不是json格式，处理不了
		http.Error(w, log.String(), http.StatusUnsupportedMediaType)
		return
	}

	if method := r.Method; method != "POST" {
		log.WriteString(fmt.Sprintf("requst mehtod:%s, expect `POST`", method))
		klog.Errorf(log.String())
		// 非POST调用，返回是405
		http.Error(w, log.String(), http.StatusMethodNotAllowed)
		return
	}

	//读取从ApiServer过来的数据放到body
	var body = make([]byte, r.ContentLength)
	r.Body.Read(body)
	// if _, err := r.Body.Read(body); err != nil {
	// 	log.WriteString(fmt.Sprintf("read body err: %v", err))
	// 	klog.Errorf(log.String())
	// 	// 读取body错误，返回是500
	// 	http.Error(w, log.String(), http.StatusInternalServerError)
	// 	return
	// }
	// if r.Body != nil {
	// 	if data, err := ioutil.ReadAll(r.Body); err == nil {
	// 		body = data
	// 	}
	// }

	if len(body) == 0 {
		log.WriteString("empty body")
		klog.Info(log.String())
		//返回状态码400
		//如果调用返回是400，说明传过来的数据是空
		http.Error(w, log.String(), http.StatusBadRequest)
		return
	}

	var data DaLiuRenReust
	if err := json.Unmarshal(body, &data); err != nil {
		log.WriteString(fmt.Sprintf("解析body成json错误，%s", err.Error()))
		klog.Info(log.String())
		//返回状态码400
		//如果调用返回是400，解析json错误
		http.Error(w, log.String(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		log.WriteString(fmt.Sprintf("failed to validate request，%s", err.Error()))
		klog.Info(log.String())
		http.Error(w, log.String(), http.StatusBadRequest)
		return
	}

	ephePath := os.Getenv("EPHE_PATH")
	if ephePath == "" {
		log.WriteString("EPHE_PATH must be specified")
		klog.Info(log.String())
		http.Error(w, log.String(), http.StatusInternalServerError)
		return
	}

	p, err := daliuren.NewDaliurenShiPan(data.Year, data.Month, data.Day, data.Hour, data.Minute, data.Second, data.YueJiang, data.DivinationTime, data.Diurnal, data.YearOfBirth, data.Masculine, ephePath)
	if err != nil {
		log.WriteString(fmt.Sprintf("计算六壬式盘失败，%s", err.Error()))
		klog.Info(log.String())
		http.Error(w, log.String(), http.StatusBadRequest)
		return
	}
	jsonPan, err := json.Marshal(p)
	if err != nil {
		log.WriteString(fmt.Sprintf("序列化式盘失败，%s", err.Error()))
		klog.Info(log.String())
		http.Error(w, log.String(), http.StatusBadRequest)
		return
	}

	// // fmt.Println(string(jsonPan))
	// pan := make(map[string]interface{})
	// err = json.Unmarshal(jsonPan, &pan)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// lunarCalendar := pan["lunarCalendar"].(map[string]interface{})
	// fmt.Printf("%v", lunarCalendar["year"])

	if _, err := w.Write(jsonPan); err != nil {
		log.WriteString(fmt.Sprintf("Failed send response: %v", err.Error()))
		klog.Info(log.String())
		http.Error(w, log.String(), http.StatusInternalServerError)
	}
}

type DaLiuRenReust struct {
	Year           int    `json:"year" validate:"required"`
	Month          int    `json:"month" validate:"min=1,max=12"`
	Day            int    `json:"day" validate:"min=1,max=31"`
	Hour           int    `json:"hour" validate:"min=0,max=23"`
	Minute         int    `json:"minute" validate:"min=0,max=59"`
	Second         int    `json:"second" validate:"min=0,max=59"`
	YueJiang       string `json:"yue_jiang"`                           // 月将，如：寅，空，表示自动计算
	DivinationTime string `json:"divination_time" validate:"required"` // 占时，如：寅
	Diurnal        bool   `json:"diurnal"`                             // 昼占:true，夜占:false
	YearOfBirth    int    `json:"year_of_birth" validate:"required"`   // 出生年
	Masculine      bool   `json:"masculine"`                           // 性别，男：true，女：false
}
