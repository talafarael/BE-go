FROM golang:1.24 

ENV APP_ENV=production
ENV PORT=8080

WORKDIR /BE-go

# install dependentcy
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o main ./cmd/main.go

EXPOSE ${PORT}

CMD ["./main"]



