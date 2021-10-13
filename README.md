## template-golang-echo
> Creator: 한승진
> Date: 2021/10/13

## Description
* Kubernetes-Istio 환경인 리얼패킹 클라우드에 배포하기 위한 Golang Echo 서버 기초 템플릿입니다.
* 배포 타겟 기술스텍 : Kubernetes 1.12.7 / Istio 1.5.4 / Jaeger 1.8.0  / Enable Horizental Pod AutoScaler

## 기본 구성 라이브러리
* golang 1.15
* require github.com/labstack/echo/v4 v4.1.16
* require github.com/labstack/echo-contrib v0.9.0
* require github.com/opentracing/opentracing-go v1.1.0
* require github.com/uber/jaeger-client-go v2.19.1-0.20191002155754-0be28c34dabf

## 해당 템플릿을 기반으로 개발하고자 한다면
* /k8s 하위의 values.yaml, prod.values.yaml을 수정해주세요.
* skaffold.yaml에 예시로 되어있는 벨류들을 변경해주세요.
* /src/go.mod의 첫번째 라인 모듈명을 변경해주세요.
* /src/main.go 상단 import 부분에 jaegertracing "example.template.com/rp_echo_tracer"를 go.mod에 변경된 모듈명로 맞추어 변경해주세요
* 필요에 따라 Dockerfile 수정이 필요할 수 있으니 Dockerfile의 내용도 확인해주세요.

## Echo framework
* Echo 프레임워크는 상당히 유용한 기능을 쉽게 제공합니다. 
* 공식 가이드 문서를 확인하며 개발해주세요.
* https://echo.labstack.com/guide