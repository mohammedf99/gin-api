# 🎉 gin-api

A clean, modular starter **REST API** built with **Gin (Go)** — designed to showcase backend development skills.

---

## 📂 Project Structure

/
├─ db/          — database setup & connections


├─ middlewares/ — request logging, auth, CORS, etc.


├─ models/      — data models + DB logic


├─ routes/      — route definitions by resource


├─ utils/       — helpers (error responses, validation)


├─ main.go      — entrypoint: sets up Gin, middleware, routes


├─ go.mod       — Go module info

---

## 🚀 Quick Start

1. **Clone & Install Dependencies**  
   ```bash
   git clone https://github.com/mohammedf99/gin-api.git
   cd gin-api
   go mod tidy

2. **Configure Database**
   
Set your DB credentials (PostgreSQL/MySQL) in db/.

4. **Run the Server**
   
go run main.go


