FROM golang:1.22.3-alpine

WORKDIR /app

# install dependencies
RUN apk add --no-cache gcc musl-dev

# copy go mod and sum files
COPY go.mod go.sum* ./

# download dependencies
RUN go mod download && go mod verify

# copy source code
COPY . .

# build the application
RUN go build -o checkin-app .

# expose API port
EXPOSE 8080

# run migrations and then start the application
CMD ./checkin-app -migrate && ./checkin-app -api