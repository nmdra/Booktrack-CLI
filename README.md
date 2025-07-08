# Booktrack CLI

A simple command-line tool to track your books using SQLite and Go.

---

## Features

- Add books with title, author, and optional year
- Stores data in a local SQLite database (`books.db`)
- Built using Go, Cobra (CLI), and sqlc (type-safe SQL)

---

## Installation

```bash
git clone https://github.com/nmdra/Booktrack-CLI.git
cd Booktrack-CLI
make build
````
This generates the binary at `./bin/booktrack-cli`.


## 🛠️ Usage

### 1. Initialize the database (run once)

```bash
./bin/booktrack-cli initdb
```

This creates the required table in `books.db`.

---

### 2. Add a new book

```bash
./bin/booktrack-cli add --title "Morning Star" --author "Pierce Brown" --year 2016
```

