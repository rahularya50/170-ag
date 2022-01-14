# syntax=docker/dockerfile:1

# Compile Go binary
FROM golang:1.16-alpine as gobuild

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go run github.com/99designs/gqlgen generate
RUN go build -o /webserver ./cmd/site/

# Compile frontend
FROM node:12.18.1 as jsbuild
ENV NODE_ENV=production

WORKDIR /app/frontend

COPY ["frontend/package.json", "frontend/yarn.lock", "./"]

RUN yarn

COPY frontend/* ./

RUN yarn build

# Start webserver
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=gobuild /webserver /webserver
COPY --from=jsbuild /app/frontend/build/* /frontend/build/*

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/webserver"]
