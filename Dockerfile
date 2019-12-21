FROM golang as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build

FROM scratch

COPY --from=builder /app/go-supervise /

CMD [ "/go-supervise" ]