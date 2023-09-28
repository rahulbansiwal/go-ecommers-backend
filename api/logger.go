package api

import (
	"ecom/db/util"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type logFormat struct {
	ClientIP      string `json:"client_ip"`
	Method        string `json:"method"`
	Path          string `json:"path"`
	Timestamp     string `json:"timestamp"`
	Protocol      string `json:"protocol"`
	StatusCode    int    `json:"status_code"`
	UserAgent     string `json:"useragent"`
	ErrorMessage  string `json:"error_message"`
	Latency       string `json:"latency"`
	ClientAddress string `json:"client_address"`
}

func jsonLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			var log logFormat
			log.ClientAddress = params.Request.RemoteAddr
			log.ClientIP = params.ClientIP
			log.ErrorMessage = params.ErrorMessage
			log.Latency = fmt.Sprint(params.Latency.Milliseconds())
			log.Timestamp = params.TimeStamp.Format(time.RFC1123)
			log.Method = params.Method
			log.StatusCode = params.StatusCode
			log.UserAgent = params.Request.UserAgent()
			log.Path = params.Path
			log.Protocol = params.Request.Proto
			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}

// func loggingInFile(config util.Config) (*os.File, error) {
//
// 	logilename := fmt.Sprint(config.LOGFILEPREFIX)
// 	fileexisit := false
// 	for _, f := range files {
// 		if f.Name() == logilename {
// 			fileexisit = true
// 		}
// 	}
// 	if !fileexisit {
// 		file, err := os.Create(config.LOGPATH + "/" + logilename)
// 		if err != nil {
// 			fmt.Println("err2", err)
// 			return nil, err
// 		}
// 		return file, nil
// 	}
// 	return nil, fmt.Errorf("file is not opend or created")
// }

func loggingInFile(config util.Config) (*os.File, error) {
	_, err := os.ReadDir(config.LOGPATH)
	if err != nil {
		if !os.IsExist(err) {
			if err = os.Mkdir(config.LOGPATH, fs.ModeAppend); err != nil {
				return nil, err
			}
		}
	}
	f, err := os.Create(config.LOGPATH + "/" + config.LOGFILEPREFIX)
	return f, err
}
