
# Go Application with Prometheus Metrics
### This project demonstrates a basic Go application that exposes metrics to Prometheus for monitoring. It includes a Dockerized setup using Docker Compose.

Project Structure:
- `Go Application`: A simple Go app that provides two metrics:

1. http_requests_total: A counter for HTTP requests.
2. room_temperature_celsius: A gauge for the room temperature.
3. Dockerfile: Containerizes the Go application.

- `docker-compose.yml`: Defines the services for Prometheus and the Go application.

- `prometheus.yml`: Configures Prometheus to scrape metrics from the Go application.
