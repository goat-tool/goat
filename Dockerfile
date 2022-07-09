FROM golang:1.18.3 as build

ENV APP_NAME goat

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -v -o $APP_NAME

FROM alpine:3.14

ENV APP_NAME goat

RUN mkdir /app

COPY --from=build /app/$APP_NAME /usr/local/bin/$APP_NAME

CMD $APP_NAME serve
