# 📘 Scalable Leaderboard System with Search & Tie-Aware Ranking

This project implements a scalable leaderboard system with accurate ranking, tie handling, and fast user search.  
It was built as part of the **Matiks – Full-Stack Engineer Intern Assignment**.

---

## 🚀 Features

- Handles **10,000+ users** (architecture designed to scale to millions)
- **Dense ranking** (users with the same rating share the same rank)
- Real-time **random score updates**
- Fast **username-based search** with global rank lookup
- Single-server setup (backend + UI) for reliability and simplicity

---

## 🛠 Tech Stack

- **Backend:** Golang, Gin Framework  
- **Frontend:** HTML + Vanilla JavaScript (served by backend)  
- **Data Store:** In-memory (extendable to Redis Sorted Sets)

> **Note:** Expo Web had runtime issues on Windows environments.  
> To ensure a stable and demonstrable solution, the UI is served directly from the backend while keeping all core logic intact.

---

## 📂 Project Structure

matiks-leaderboard/
│
├── backend/
│ ├── main.go
│ ├── handlers/
│ │ ├── leaderboard.go
│ │ └── search.go
│ ├── services/
│ │ └── leaderboard_service.go
│ ├── models/
│ │ └── user.go
│ ├── seed/
│ │ └── seed.go
│ └── static/
│ └── index.html
│
└── README.md

yaml
Copy code

---

## ▶️ How to Run Locally

### Prerequisites
- Go **1.20+**
- Any modern browser (Chrome / Edge)

### Run the Application

```bash
cd backend
go run main.go
You should see:

nginx
Copy code
Listening and serving HTTP on :9090
Open in Browser
arduino
Copy code
http://localhost:9090
🔍 Usage
Leaderboard
Displays Rank | Username | Rating

Uses tie-aware dense ranking

Search
Type a username (e.g. rahul)

Instantly displays:

Global Rank

Username

Rating

Example Output
yaml
Copy code
Rank | Username     | Rating
1    | rahul_3528   | 5004
2    | rahul_8573   | 5000
5    | rahul_7105   | 4997
🧠 Ranking Logic
Users are sorted by rating in descending order

Dense ranking is applied:

Same rating → same rank

Rank increases only when rating decreases

Example
yaml
Copy code
Rating 5000 → Rank 1
Rating 5000 → Rank 1
Rating 4999 → Rank 3
📈 Scalability Design
Current Implementation
In-memory slice with mutex locking

Efficient and responsive for 10k+ users

Future Enhancements
Redis Sorted Sets (ZADD, ZRANK) for millions of users

Stateless backend for horizontal scaling