# syntax=docker/dockerfile:1

# Compile Go binary
FROM golang:1.17-buster as gobuild

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

# make ent is checked in, so no need to run here
RUN make graphql
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /webserver ./cmd/site/

# Compile frontend
FROM node:16 as jsbuild
ENV NODE_ENV=production

WORKDIR /app/frontend

COPY ["frontend/package.json", "frontend/yarn.lock", "./"]

RUN yarn

COPY frontend/ ./
COPY schema/schema.graphql ../schema/schema.graphql

RUN yarn relay
RUN yarn build

# Start webserver
FROM alpine
RUN apk add --no-cache ca-certificates

WORKDIR /

COPY --from=gobuild /webserver /webserver
COPY --from=jsbuild /app/frontend/build/ /frontend/build/

EXPOSE 8080

# RUN adduser -D nonroot
# USER nonroot

ENTRYPOINT ["/webserver"]
