# Friends Management App (Backend)

This is the backend service for the Friends Management App, built using **Go**, **Gin**, **GORM**, and **MySQL**. This service handles the core functionalities such as friend management, friend requests, blocking users, and retrieving mutual friends.

## Getting Started

Follow these steps to set up and run the backend service locally:

### 1. Install Dependencies

```bash
go mod init | go mod tidy
```
### 2. Config Database
Rename `config.yml.example` to `config.yml` and fill the credential of your database, also make sure the connection of your db

### 3. Start Development Server

```bash
go run main.go
```