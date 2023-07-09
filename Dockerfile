FROM golang:latest

WORKDIR /app

# COPY api .
# COPY DB .
# COPY nginx .
# COPY .env .
# COPY main.go .
# COPY go.mod .
# COPY go.sum .

COPY . .

RUN go build -o myapp

ENTRYPOINT ["./myapp"]