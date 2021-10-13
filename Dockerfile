# 빌드할 환경
FROM golang:1.15-alpine3.12 as build
WORKDIR /src/

# gomod 의존성 패키지 있다면 설치
COPY src/go.mod src/go.sum ./
RUN go mod download

# source 카피 및 빌드
COPY ./src /src/
RUN CGO_ENABLED=0 go build -o /bin/main

# 도커 빌더 패턴, 무거운 환경에서 빌드하고, 실제 실행시킬 컨테이너는 가벼운 컨테이너로 구성(scratch/alpine중 택)
FROM alpine:3.12
#COPY ./src /src/ 소스나 컨텐츠들 카피 필요할 경우 /src 디렉터리에 옮길 것 (비권장)
COPY --from=build /bin/main /src/main
ENTRYPOINT ["/src/main"]