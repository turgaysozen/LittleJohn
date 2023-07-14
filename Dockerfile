FROM golang:1.20
WORKDIR /app
COPY . .
EXPOSE 8080
RUN go test ./api/handlers
RUN go build -o littlejohn .
ENTRYPOINT ["./littlejohn"]
CMD ["-p", "8080:8080", "--network", "host"]
