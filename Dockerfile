FROM golang:1.19

# LABEL stage=builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app

ADD go.mod ./
ADD go.sum ./
RUN go mod download
COPY ./ ./

# RUN go install recipeApi


# RUN apk add --no-cache ca-certificates
# RUN apk add --no-cache tzdata
# Build the Go application
RUN go build -ldflags="-s -w" -o ./recipeAPI cmd/main.go

CMD /recipeAPI

# WORKDIR /app
# COPY --from=build /app .
# COPY --from=builder /app /app

CMD ["./recipeAPI"]
# Set the entry point for the container
EXPOSE 8080