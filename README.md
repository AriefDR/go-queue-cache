# ⚙️ Simple Queue-Based Cache in Go

This project is a simple implementation of a fixed-size queue-based cache using a doubly linked list and hash map. It simulates the basic concept of a **Least Recently Used (LRU)** cache mechanism.

## 📦 Tech Stack

- 🐹 Golang
- 🧠 Custom Queue implementation
- 🔗 Doubly Linked List
- ⚡ Hash Map (for O(1) lookup and updates)

## 🧠 How It Works

- The queue uses a **doubly linked list** to store nodes.
- A **hash map** keeps track of node references for O(1) access.
- When a new element is added:
  - If it's already in the cache, it is moved to the front (most recent).
  - If it's new and the cache is full, the **least recently used** item is removed.
- The queue has a **fixed size (5)**, configurable via the `SIZE` constant.

## 🚀 Usage

Clone the repo and run:

```bash
go run main.go
