FROM golang:alpine as build

RUN apk add git

WORKDIR /app
COPY . .
RUN go build


FROM alpine

COPY --from=build /app/go-supervise /
COPY server.config.yml /server.config.yml

CMD [ "/go-supervise" ]