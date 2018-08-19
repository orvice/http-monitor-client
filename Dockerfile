FROM golang:1.10 as builder

## Create a directory and Add Code
RUN mkdir -p /go/src/github.com/orvice/http-monitor-client
WORKDIR /go/src/github.com/orvice/http-monitor-client
ADD .  /go/src/github.com/orvice/http-monitor-client

# Download and install any required third party dependencies into the container.
RUN go get
# RUN go-wrapper install
RUN CGO_ENABLED=0 go build


FROM orvice/go-runtime

COPY --from=builder /go/src/github.com/orvice/http-monitor-client/http-monitor-client .

ENTRYPOINT [ "./http-monitor-client" ]