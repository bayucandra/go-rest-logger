FROM golang:1.13-alpine AS builder
ENV APP_NAME go-rest-logger

WORKDIR "/go/src/github.com/bayucandra/${APP_NAME}"


RUN apk update && apk add --no-cache git openssh-client
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /usr/bin/${APP_NAME} ./cmd/$APP_NAME/...



FROM alpine
#COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/bin/${APP_NAME} /usr/bin/${APP_NAME}
RUN adduser -D -g '' biqdev
USER biqdev
ENTRYPOINT ["/usr/bin/go-rest-logger"]
