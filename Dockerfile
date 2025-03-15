FROM golang:1.22.3-alpine

WORKDIR /app

# install dependency
RUN apk add --no-cache gcc musl-dev

# copy go mod
COPY go.mod go.sum* ./

# download dependency
RUN go mod download && go mod verify

# copy programming codes
COPY . .

# complie
RUN go build -o checkin-app .

# execution
CMD ["./checkin-app"]