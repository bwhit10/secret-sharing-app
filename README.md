# 🔐 Secret Sharing App

This is a secure "one-time secret sharing" service. Submit a secret, receive a unique one-time URL, and share it securely. Once the secret is viewed, it is permanently deleted and cannot be retrieved again.

---

## ✅ Features

- ⏳ Secrets auto-expire after a configurable TTL (time to live).
- 🔒 Secrets are encrypted in memory (not stored in plain text).
- 🔁 One-time access only — the secret self-destructs after being viewed once.
- 🐳 Dockerized for easy deployment.
- ✅ Includes test for expired secret case (can be extended).
- 💡 Clean, simple Go backend with no external dependencies except Echo.

---

## 🚀 How to Run Locally (with Docker)

1. **Clone the repo**:

   ```bash
   git clone https://github.com/YOUR-USERNAME/devex-challenge-bwhit10.git
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
├── internal/              # Internal Go packages
│   └── secret/            # Business logic: store, model, encryption, handlers
├── scripts/               # Utility shell scripts (e.g., port checks)
├── test/                  # Place for additional tests (if needed)
├── web/                   # Placeholder for frontend (not used in this version)
├── .gitignore             # Git ignored files
├── Dockerfile             # Docker build instructions
├── go.mod                 # Go module definition
├── go.sum                 # Go dependencies checksum
├── main.go                # App entry point and HTTP handlers
├── README.md              # Project documentation
└── template.go            # Optional: HTML/template helpers

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
 📫 linkedin.com/in/brittanywhitfield
DevOps | Platform Engineering | Security-Conscious
