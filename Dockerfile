FROM golang:1.24.4-alpine as builder

WORKDIR /app

COPY go.mod go.mod
#COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o bin/password-generator-server ./cmd/app/

FROM alpine

COPY --from=builder /app/bin/password-generator-server /password-generator-server

CMD [ "/password-generator-server" ]
