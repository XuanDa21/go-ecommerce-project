FROM golang:1.20 as build-stage

# Set destination for COPY
WORKDIR /src/

COPY . .

# Download Go modules
RUN go mod download

# Build
# WORKDIR /src/ecommerce/

RUN go build -ldflags="-s -w" --trimpath -o ecommerce ./

FROM debian:stable as server

RUN apt-get update
RUN apt-get install ca-certificates -y

WORKDIR /app

COPY --from=build-stage /src/ecommerce /app/ecommerce

ENTRYPOINT ["./ecommerce"]