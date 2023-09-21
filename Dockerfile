# use official Golang image
FROM golang:1.19

# RUN apk update
# RUN apk upgrade
# RUN apk add --no-cache libc6-compat

# set working directory
WORKDIR /app

COPY go.mod go.sum ./
# Copy the source code
COPY . ./

# Download and install the dependencies
RUN go mod tidy
RUN ls
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /savannah
# RUN go build -o savannah .

#EXPOSE the port
EXPOSE 8181

# Run the executable
CMD ["./Savannah-e-shop"]