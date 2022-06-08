# syntax=docker/dockerfile:1

FROM golang:1.18

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
copy . .
RUN go mod download

# Copy the source code.
ADD /cmd ./
ADD /pkg ./
COPY *.go ./

# Build

RUN go build -o /dna-app/cmd/gopherapi/ ./cmd/gopherapi
RUN chmod +x /app
RUN chmod +x ./cmd/gopherapi

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Run
CMD [ "/app" ]