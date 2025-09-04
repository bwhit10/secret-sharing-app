# 🔐 Secret Sharing App

This is a simple Go-based secret-sharing service that allows users to:
- Create encrypted secrets via an HTTP POST request.
- Retrieve secrets exactly **once** using a unique ID.
- Automatically expire secrets after a given TTL (Time-To-Live).
- Store all secrets in memory only (ephemeral storage).

---

## 🚀 Running the Application

### 🖥️ Option 1: Run Locally with Go

```bash
# Clone the repo
git clone https://github.com/yourusername/devex-challenge-bwhit10.git
cd devex-challenge-bwhit10

# Run the application
go run main.go
```

### 🐳 Option 2: Run with Docker

```bash
# Build the Docker image
docker build -t secret .

# Run the container
docker run -p 8080:8080 secret
```

The app will now be available at:  
[http://localhost:8080](http://localhost:8080)

---

## 📡 API Endpoints

### 🔸 `POST /secret`

- **Description:** Create a secret
- **Request Body (JSON):**
```json
{
  "secret": "My secret goes here",
  "ttl": 300
}
```
- **Response:**
```json
{
  "url": "http://localhost:8080/secret/abc123"
}
```

### 🔸 `GET /secret/{id}`

- **Description:** Retrieve a secret by its unique ID.
- **Behavior:** Only works **once**. After it's been retrieved or expired, you’ll get:
```json
{
  "error": "Secret not found or expired"
}
```

---

## 🧪 Running Tests

To run unit tests:

```bash
go test ./internal/secret
```

You’ll see output similar to:

```
ok  	devex-challenge-bwhit10/internal/secret	0.386s
```

---

## 📁 Project Structure

```plaintext
devex-challenge-bwhit10/
├── Dockerfile
├── README.md
├── go.mod
├── go.sum
├── main.go
├── internal/
│   └── secret/
│       ├── handler.go
│       ├── handler_test.go
│       ├── model.go
│       └── store.go
├── test/
├── scripts/
└── web/
```

---

## 🧠 Key Features

- 🔒 AES-256 encrypted secrets
- 🕓 Secrets expire after TTL (seconds)
- 🔁 Secrets can only be retrieved **once**
- 🐳 Dockerized for easy deployment
- ✅ Unit tests included

---

## 🙋🏽‍♀️ Author

**Brittany Whitfield**  
DevOps | Platform Engineering | Security-Conscious
