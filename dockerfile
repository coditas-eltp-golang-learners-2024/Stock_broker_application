# Use an official Golang runtime as a parent image
FROM golang:1.20

# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY ./ /app

# Set the environment variable
ARG MICROSERVICE=authentication
ARG PORT= 8080

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
WORKDIR /app/src/app/$MICROSERVICE

# Copy go mod and sum files
RUN go mod tidy

#Testing the microservice path
RUN pwd
RUN ls -a /app/src/app/$MICROSERVICE

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go install .

# SWAG installation
RUN curl -o /app/src/app/$MICROSERVICE/Swag https://github.com/linuxserver/docker-Swag/releases/latest/download/Swag && chmod +x /app/src/app/$MICROSERVICE/Swag

#Expose port to outside world
EXPOSE $PORT

# Build the Go app
RUN go build -o main .

# Swagger update
RUN /app/src/app/$MICROSERVICE/Swag init -g main.go

# Command to run the executable
CMD ["./main"]
