FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY service/api/wxapp/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/wxApp service/api/wxapp/wxApp.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/wxApp /app/wxApp
COPY --from=builder /app/etc /app/etc

CMD ["./wxApp", "-f", "etc/wxapp.yaml"]
