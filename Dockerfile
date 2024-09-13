FROM golang:1.23.1-alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o fbpushd ./

FROM alpine:3.18

WORKDIR /app

COPY --from=build /app/fbpushd /app/fbpushd

EXPOSE 8080

CMD ["./fbpushd"]
