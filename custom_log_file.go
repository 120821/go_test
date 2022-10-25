package main

import (
  "log"
  "github.com/gin-gonic/gin"
  "os"
  "fmt"
  "time"
)
func main() {
  LOG_FILE := "/tmp/myapp_log"
  // open log file
  logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
  if err != nil {
    log.Panic(err)
  }
  defer logFile.Close()

  // Set log out put and enjoy :)
  log.SetOutput(logFile)

  // optional: log date-time, filename, and line number
  log.SetFlags(log.Lshortfile | log.LstdFlags)

  log.Println("Logging to custom file")
	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run(":9888")
}

