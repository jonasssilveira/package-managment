# 🧠 Optimal Package Service

This service receives an order quantity and calculates the optimal combination of available pack sizes to fulfill that order with minimal waste and/or fewer packs. It also supports CRUD operations on available pack sizes using a MongoDB database.

---

## 📦 Features

- ✅ Calculate optimal pack distribution (`/packs-find`)
- ✅ Add a new pack size (`/packs-create`)
- ✅ Remove a pack size (`/packs/:size`)

---

## 🚀 Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 🐳 Running with Docker

```bash
docker-compose up --build
```

This will start:
- The Go application (exposed at `http://localhost:8080`)
- A MongoDB instance with initial seed data

---

## 🧪 Endpoints

### 1. **Find Optimal Packs**

```http
POST /packs-find
```

**Request Body:**

```json
{
  "amount": 1200
}
```

**Response:**

```json
{
  "250": 1,
  "1000": 1
}
```

---

### 2. **Add a New Pack**

```http
POST /packs-create
```

**Request Body:**

```json
{
  "size": 750
}
```

**Response:** `201 Created` on success

---

### 3. **Delete a Pack by Size**

```http
DELETE /packs/:size
```

**Example:**

```http
DELETE /packs/750
```

**Response:** `200 OK` on success

---

## 📁 Project Structure

```
├── internal
│   ├── domain          # Core business logic
│   ├── infra           # Infra layer (MongoDB, server, etc.)
│   └── script          # Docker + Mongo seed
├── main.go             # App entrypoint
├── go.mod / go.sum
```

---

## 🧪 Running Tests

```bash
go test ./...
```

Unit tests include:
- Optimal pack calculation
- Pack creation and deletion logic
- Repository mocking

---

## 📄 License

MIT License. Feel free to use and modify.