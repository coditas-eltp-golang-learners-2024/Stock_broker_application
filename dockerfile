# Use an official Golang runtime as a parent image
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

RUN go mod tidy





# Copy go.mod and go.sum files to the workspace
# COPY src/app/authentication/go.mod src/app/authentication/
# COPY src/app/watchlist/go.mod src/app/watchlist/

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# # Copy the entire source code into the container
# COPY . .

# # Build the Go apps separately
# RUN go build -o authentication ./src/app/authentication
# RUN go build -o watchlist ./src/app/watchlist

# # Expose port 8080 to the outside world (assuming authentication uses 8080)
# EXPOSE 8080

# # Command to run the authentication service
# CMD ["./authentication"]
