FROM golang:1.22.4-alpine

WORKDIR ${GOPATH}/src/github.com/yuminekosan/hikariLibBackend
RUN go version

RUN go install -tags '${DB_SERVICE}' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN ln -s /go/bin/linux_amd64/migrate /usr/local/bin/migrate

COPY . .

RUN go mod download
RUN go build -o hl-backend ./cmd/main.go


#FROM alpine:latest as dev
#RUN apk add gcompat
#WORKDIR /app
#COPY --from=builder hl-backend ./
#
ENTRYPOINT []
