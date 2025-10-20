FROM golang:1.24 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /iris-api

FROM golang:1.24.0-alpine as run

WORKDIR /go/bin

COPY --from=build /iris-api /go/bin/iris-api

EXPOSE 8080

CMD ["iris-api"]