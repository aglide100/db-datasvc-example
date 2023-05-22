FROM golang:1.20 AS builder

RUN apt update && apt install -y ca-certificates upx

WORKDIR "/go/src/github.com/aglide100/db-datasvc-snippet/"

COPY go.mod .
COPY go.sum .

RUN mkdir -p /opt/bin/apid/

RUN go mod tidy \
    && go get -u -d -v ./...

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w' -o /opt/bin/apid/main cmd/apid/main.go \    
    && upx --lzma /opt/bin/apid/main

FROM debian:stretch-slim AS runtime
COPY --from=builder /opt/bin/apid /opt/bin/apid
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD [ "/opt/bin/apid/main" ]
