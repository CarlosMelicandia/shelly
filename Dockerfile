FROM node:16-alpine as frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ .
RUN npm run build

FROM golang:1.19-alpine as backend-builder
WORKDIR /app
COPY webserver/go.mod webserver/go.sum ./
RUN go mod download
COPY webserver/*.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:3.14
WORKDIR /app
COPY --from=frontend-builder /app/frontend/dist ./dist
COPY --from=backend-builder /app/server .
EXPOSE 8080
CMD ["./server"]
