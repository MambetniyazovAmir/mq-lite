# MQ-Lite 🚀
*A lightweight message broker written in Go*

## 📌 Overview
MQ-Lite is a lightweight **pub/sub message broker** implemented in Go.  
It is designed to demonstrate concepts of **network protocols, concurrency, and distributed systems** while being simple enough to run locally in Docker.

✅ Features:
- Topic-based **publish/subscribe** messaging  
- **TCP protocol** with custom commands (`PUB`, `SUB`)  
- **In-memory message routing** (with optional persistence layer)  
- **Go client SDK** for publishers and subscribers  
- **Monitoring with Prometheus** (messages/sec, active clients, latency)  
- Dockerized deployment + CI/CD via GitHub Actions  

---

## 🔧 Quick Start

### Run with Go
```bash
git clone https://github.com/your-username/mq-lite.git
cd mq-lite/cmd/broker
go run main.go
