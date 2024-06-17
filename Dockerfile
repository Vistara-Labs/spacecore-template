# Starter template
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the application files to the container
COPY . .

CMD ["ls", "-lhtr"]