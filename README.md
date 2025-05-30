# 🧮 Order Pack Calculator

This project provides a backend service written in Go that calculates optimal packaging combinations based on available pack sizes. It also comes with a simple HTML frontend interface for testing.

---

## 🚀 Features

- Create, retrieve, and delete pack sizes
- Calculate the optimal number of packs for a given quantity
- REST API with JSON payloads
- In-memory data storage for easy deployment
- HTML frontend for interactive testing

---

## 📦 Endpoints

| Method | Endpoint         | Description                      |
|--------|------------------|----------------------------------|
| POST   | `/packs-create`  | Create multiple pack sizes       |
| POST   | `/packs-find`    | Calculate optimal pack combo     |
| DELETE | `/packs/:size`   | Delete a pack by its size        |

---

## 🌐 Frontend

Open `index.html` in your browser to test the backend interactively.

---

## 🛠️ Requirements

- Go 1.20+
- Make
- Docker (optional for containerization)

---

## 🧪 Running Locally

### 🔹 Clone the Repository

```bash
git clone  https://github.com/jonasssilveira/package-managment.git
cd order-pack-calculator
```

### 🔹 Run the App

```bash
make run
```

The server will start at [http://localhost:8080](http://localhost:8080)

### 🔹 Run Tests

```bash
make test
```

### 🔹 Run Tests with Coverage

```bash
make coverage
```

### 🔹 Clean Up

```bash
make clean
```

---

## 🐳 Docker

Build and run using Docker:

```bash
docker build -t order-pack .
docker run -p 8080:8080 order-pack
```

---

## 📂 Project Structure

```
├── internal/
│   └── domain/
│       └── optimalpackage/
│           ├── adapters/
│           ├── dto/
│           ├── entity/
│           ├── mock/
│           ├── *.go
│   └── infra/
│       ├── config/
│       ├── repository/
│       ├── *.go
├── script/
│   └── docker-compose.yaml
├── Makefile
├── main.go
├── README.md
```

---

## 👨‍💻 Author

Made with 💚 by Jonas Silveira
