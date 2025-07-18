# compile phase

FROM docker.io/golang:latest AS build

WORKDIR /app

COPY main.go go.mod .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server .

# build phase

FROM scratch

WORKDIR /

COPY --from=build /app/server /server

EXPOSE 8080

CMD ["/server"]
