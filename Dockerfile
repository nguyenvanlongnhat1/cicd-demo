FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /cicd-demo
EXPOSE 3000
CMD ["/cicd-demo"]
