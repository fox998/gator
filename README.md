# 🐊 gator

`gator` is a powerful command-line tool written in Go for managing and interacting with your PostgreSQL database. It simplifies common tasks like migrations, seeding, and database status checks.

***

## 🚨 Prerequisites

To successfully run and build **gator**, you need to have the following prerequisites installed and configured on your system:

1.  **Go:** The Go programming language (version 1.18 or later is recommended).
2.  **PostgreSQL:** A running and accessible PostgreSQL instance, as `gator` is designed to interact with it.

***

## 📦 Installation

You can easily install the `gator` CLI tool using the standard Go installation command:

```bash
go install [github.com/fox998/gator@latest](https://github.com/fox998/gator@latest)

```

You need to create a .gatorconfig.json in the ~/ directory with the following content:

```json
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable","current_user_name":""}
```
