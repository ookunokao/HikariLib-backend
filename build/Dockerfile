FROM golang:1.22.4-alpine

RUN go version

RUN go install -tags '${DB_SERVICE}' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN ln -s /go/bin/linux_amd64/migrate /usr/local/bin/migrate

COPY . .

RUN go mod download
RUN  go build -o hl-backend ./cmd/main.go
CMD ["./hl-backend"]

#FROM alpine:latest
#WORKDIR /root
#COPY --from=builder /app /app
