FROM alpine:3

COPY --from=golang:1.21-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]