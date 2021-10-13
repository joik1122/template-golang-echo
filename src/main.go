package main

import (
	jaegertracing "example.template.com/rp_echo_tracer"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
	"time"
)

// 헬스체크 url에 대한 스키퍼 (로그/트레이싱 스킵함)
func urlSkipper(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/healthz") {
		return true
	}
	return false
}

func main() {
	e := echo.New()

	// Jaeger tracing 미들웨어 활성화
	c := jaegertracing.New(e, urlSkipper)
	if c != nil {
		defer c.Close()
	}

	// Request에 대한 표준출력 콘솔 로그
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: urlSkipper,
		Format:  "method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Golang echo framework template!")
	})

	e.GET("/healthz", func(c echo.Context) error {
		/**
		 * 필수입니다. 쿠버네티스 probe가 healthz url에 200을 성공적으로 받지 못하면 POD이 죽은걸로 판단합니다.
		 * 관련설정 /k8s/templates/deployment.yaml에서 확인
		 */
		return c.String(http.StatusOK, "Health checked....")
	})

	e.GET("/matthew", func(c echo.Context) error {
		/**
		 * 필수입니다. 쿠버네티스 probe가 healthz url에 200을 성공적으로 받지 못하면 POD이 죽은걸로 판단합니다.
		 * 관련설정 /k8s/templates/deployment.yaml에서 확인
		 */
		return c.String(http.StatusOK, "안녕 조나스")
	})

	// Jaeger tracing 예제입니다
	e.GET("/jaeger-example", func(c echo.Context) error {

		// 특정 펑션 자동 트레이싱
		jaegertracing.TraceFunction(c, slowFunc, "Test function tracer")

		// 직접 차일드 span 만들기
		sp := jaegertracing.CreateChildSpan(c, "Child span for additional processing")
		defer sp.Finish()

		// 각종 트레이싱 정보들 추가하기 (jaeger ui에서 더 디테일한 확인이 가능합니다)
		sp.LogEvent("Test log")
		sp.SetBaggageItem("Test baggage", "baggage")
		sp.SetTag("Test tag", "New Tag")
		time.Sleep(1000 * time.Millisecond)

		return c.String(http.StatusOK, "Health jaeger")
	})

	e.Logger.Fatal(e.Start(":80"))
}

// jaeger 예시를 위한 슬로우 펑션
func slowFunc(s string) {
	time.Sleep(200 * time.Millisecond)
	return
}
