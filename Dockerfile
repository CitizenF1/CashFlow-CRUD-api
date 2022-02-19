FROM golang:1.13.4-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN mkdir /app

# Copy files
COPY . /app

# Set the Current Working Directory inside the container
WORKDIR /app

# Build the Go app
RUN go build -o api /main.go

# Run the executable
CMD ["./app/api"]