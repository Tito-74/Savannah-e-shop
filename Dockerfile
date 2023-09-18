FROM golang:1.21.1
WORKDIR /app
# RUN go install github.com/cosmtrek/air@latest
COPY . .
RUN go mod tidy
RUN go build -o savannah .
EXPOSE 3000
CMD ["./savannah"]