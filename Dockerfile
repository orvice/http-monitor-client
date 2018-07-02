FROM golang:1.10 as builder

## Create a directory and Add Code
RUN mkdir -p /go/src/github.com/orvice/http-monitor-client
WORKDIR /go/src/github.com/orvice/http-monitor-client
ADD .  /go/src/github.com/orvice/http-monitor-client

# Download and install any required third party dependencies into the container.
RUN go get
# RUN go-wrapper install
RUN CGO_ENABLED=0 go build


FROM alpine

COPY --from=builder /go/src/github.com/orvice/http-monitor-client/http-monitor-client .

RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
# Change TimeZone
RUN apk add --update tzdata
ENV TZ=Asia/Shanghai
# Clean APK cache
RUN rm -rf /var/cache/apk/*

ENTRYPOINT [ "./http-monitor-client" ]