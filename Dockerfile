FROM node:16-alpine as frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ .
RUN npm run build

FROM golang:1.19-alpine as backend-builder
WORKDIR /app/webserver
COPY webserver/go.mod ./
RUN go mod download
COPY webserver/*.go ./
RUN mkdir -p /bin && CGO_ENABLED=0 GOOS=linux go build -o /bin/Opal

FROM alpine:3.14
WORKDIR /app
COPY --from=frontend-builder /app/frontend/build ./frontend/build
COPY --from=backend-builder /bin/Opal ./webserver/bin/
EXPOSE 8080
CMD ["./webserver/bin/Opal"]