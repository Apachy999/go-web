FROM golang:latest

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/go-web
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
EXPOSE 8080
CMD ["bee", "run"]
