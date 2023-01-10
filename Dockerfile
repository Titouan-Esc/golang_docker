FROM golang:1.19

# Copy the .env file for run the api
WORKDIR /usr/src/env
COPY /Users/titouanescorneboueu/Documents/flit-env/flitSport .

# Now the api
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["go", "run", "main.go"]