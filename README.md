# Go Grafana Logs & Metrics

This project provides a comprehensive observability solution for standalone Go applications. It implements a centralized telemetry pipeline to manage logs and metrics using the Promtail, Loki, Prometheus, and Grafana stack.

## 🌟 Key Features

* **File-Based Logging**: The application writes structured JSON logs directly into a physical file located at `logs/app.log`.
* **Promtail Log Scraping**: Uses Promtail as an agent to tail the log file in real-time and forward it to Grafana Loki.
* **Automated Metrics**: Exposes application performance metrics via the `/metrics` endpoint, which is scraped periodically by Prometheus.
* **Unified Visualization**: View and correlate logs and metrics within a single dashboard in Grafana.

## 🛠 Tech Stack

* **Language**: Go (Golang)
* **Log Shipper**: Promtail
* **Log Storage**: Grafana Loki
* **Metrics Database**: Prometheus
* **Visualization**: Grafana
* **Containerization**: Docker & Docker Compose

## How to run

1. Copy file `.env.example` to `.env`.
  ```sh
  cp .env.example .env
  ```

2. Run the Project.
  ```sh
  docker compose up -d --build
  ```

3. Access Dashboards:
  * Grafana: http://localhost:3000
  * Prometheus: http://localhost:9090

## ⚙️ How It Works
1. Go Application: Generates events and writes them in JSON format to the /app/logs directory.

2. Promtail: Mounts the same directory, monitors for new entries (tailing), and ships them to Loki.

3. Prometheus: Scrapes numerical data from the application's /metrics endpoint.

4. Grafana: Aggregates data from both Loki (Logs) and Prometheus (Metrics) to provide actionable insights via graphs and tables.