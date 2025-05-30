# --- Stage 1: Builder Stage ---
# Use a Go base image for building the application.
FROM golang:1.23-alpine AS builder

# Set the working directory inside the builder container.
WORKDIR /app

# Copy go.mod and go.sum files to leverage Docker's build cache.
# This step allows downloading dependencies before copying the rest of the code,
# which speeds up builds if only source code changes, not dependencies.
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire application source code into the builder container.
# This includes main.go, internal/, and any other Go files.
COPY . .

# Build the Go application.
# CGO_ENABLED=0 creates a statically linked binary, which is important for
# running on minimal base images like alpine (which uses musl libc) or scratch.
# The output binary will be named 'package-managment' and placed at the root of /app.
RUN CGO_ENABLED=0 go build -o /app/package-managment

# --- Stage 2: Final (Runtime) Stage ---
# Use a very minimal base image for the final application.
# Alpine is a good choice for small image size.
FROM alpine:latest

# Set the working directory in the final image.
WORKDIR /app

# Copy the compiled binary from the builder stage to the final image.
# Only the executable is copied, keeping the final image small.
COPY --from=builder /app/package-managment /app/package-managment

# Copy the static assets (HTML, CSS, JS, etc.) from your project's 'static' directory
# into the '/app/static/' directory inside the final container.
# Your Go application will serve these files.
COPY static/ /app/static/

# Expose the port that your Go application will listen on.
# This informs Docker that the container listens on this port.
EXPOSE 8080

# Define the command to run when the container starts.
# This executes your compiled Go application.
CMD ["/app/package-managment"]
