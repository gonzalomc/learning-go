FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN apk add git 
RUN go get -d github.com/gorilla/mux
RUN go build -o app .

FROM alpine
WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT ["/app/app"]