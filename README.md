
# Go Application with Prometheus Metrics
## This project demonstrates a basic Go application that exposes metrics to Prometheus for monitoring. It includes a Dockerized setup using Docker Compose.

Project Structure
Go Application: A simple Go app that provides two metrics:

http_requests_total: A counter for HTTP requests.
room_temperature_celsius: A gauge for the room temperature.
Dockerfile: Containerizes the Go application.

docker-compose.yml: Defines the services for Prometheus and the Go application.

prometheus.yml: Configures Prometheus to scrape metrics from the Go application.