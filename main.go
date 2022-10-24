package main

import (
	"os"
	"regexp"
	"strings"
	"time"
	"webapi/exception"
	"webapi/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// set timezone
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst

	echo := echo.New()

	echo.Use(middleware.LoggerWithConfig(customLogger()))
	echo.Use(middleware.Recover())
	echo.Use(checkRequestMiddleware)

	// logger setting
	file, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	echo.Logger.SetLevel(log.DEBUG)
	echo.Logger.SetOutput(file)

	router(echo)

	if handler.IsDocker() {
		echo.Start(":8080")
	} else {
		echo.Start("")
	}
}

func customLogger() middleware.LoggerConfig {
	file, err := os.OpenFile("logs/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	config := middleware.DefaultLoggerConfig

	config.Output = file
	config.Skipper = customSkipper
	config.Format = "time:${time_custom}, " +
		"id:${id}, " +
		"remote_ip:${remote_ip}, " +
		"host:${host}, " +
		"uri:${uri}, " +
		"user_agent:${user_agent}, " +
		"status:${status}, " +
		"error:${error}, " +
		"latency_human:${latency_human}, " +
		"size:${bytes_out}, " +
		"x-forwarded-for:${header:x-forwarded-for}, " +
		"x-request-id:${header:X-Request-Id}"
	config.CustomTimeFormat = "2006/01/02 15:04:05.00000"

	return config
}

func customSkipper(context echo.Context) bool {
	// docker or health checker ?
	if handler.IsDocker() {
		return true
	}
	regex := regexp.MustCompile("(?i)ELB-HealthChecker")
	return regex.MatchString(context.Request().UserAgent())
}

func checkRequestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		header := context.Request().Header
		contentType := header.Get("Content-Type")
		if !strings.Contains(contentType, "application/json") {
			exception.NotAllowed(context)
			os.Exit(1)
		}

	labelDoWhile:
		for {
			if handler.IsDocker() {
				break labelDoWhile
			}

			myHeader := header.Get("X-MyHeader")
			if strings.Compare(myHeader, "MyHeaderCode") == 0 {
				break labelDoWhile
			}

			// arrow IP list
			allow := [...]string{"0.0.0.0"}
			from := header.Get("X-FORWARDED-FOR")
			for _, val := range allow {
				if strings.Compare(val, from) == 0 {
					break labelDoWhile
				}
			}

			exception.NotAllowed(context)
			os.Exit(1)
		}

		if err := next(context); err != nil {
			context.Error(err)
		}
		return nil
	}
}
