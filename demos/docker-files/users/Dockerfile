FROM golang:1.17.5-alpine AS build

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod tidy -compat=1.17
RUN go build -o main ./cmd/main.go

FROM alpine
WORKDIR /app
RUN mkdir /configs
COPY --from=build /app/configs/app.env /configs
RUN mkdir /keys
COPY --from=build /app/keys/. /keys
COPY --from=build /app/main /app/main
CMD ["/app/main"]