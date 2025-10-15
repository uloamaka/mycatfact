# 🐾 MyCatFact API

MyCatFact is a simple Go API that returns random cat facts by calling the external endpoint [`https://catfact.ninja/fact`](https://catfact.ninja/fact).  
It demonstrates clean Go code, HTTP handling, and JSON responses — perfect for learning or small demo projects.

---

## 🚀 Features

- Fetches random cat facts from [catfact.ninja](https://catfact.ninja/fact)
- Exposes a simple `/me` endpoint returning user info and a cat fact
- Uses Go’s native `net/http` package
- Includes timeout handling and JSON formatting
- Lightweight and deployment-ready (e.g., PXXL app, etc.)

---

## 🛠 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/uloamaka/mycatfact.git
cd mycatfact
```

### 2. Install Dependency & Build
```
go mod tidy
go install
go build
```

### 3. Run the Api
``` ./mycatfact ```
### 4. Test Endpoint
``` curl http://localhost:8080/me ```
## You should see a json response similar to
```
{
  "status": "success",
  "user": {
    "email": "ebitegift235@gmail.com",
    "name": "Godsgift Uloamaka Ebite",
    "stack": "Go/Gin"
  },
  "timestamp": "2025-10-15T18:10:21.123Z",
  "facts": "Cats have five toes on their front paws but only four on the back ones."
}
```
