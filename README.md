# GoLang Photo Management API

A simple API for managing user registration, authentication, and photo uploads using Go (Golang), Gin Gonic, and Gorm.

## Table of Contents

1. [Overview](#overview)
2. [Requirements](#requirements)
3. [Getting Started](#getting-started)
4. [Project Structure](#project-structure)
5. [Contributing](#contributing)
6. [License](#license)

## Overview

This project provides a GoLang-based API for handling user registration, authentication, and photo management. It includes endpoints for registering and logging in users, uploading photos, updating user information, and deleting users.

### Features

- User registration with unique email and password validation
- User login using email and password
- Photo upload with title, caption, and URL
- Update user information
- Delete user accounts
- Authorization using JWT tokens
- Relationship between users and photos

## Requirements

Ensure you have the following installed:

- Go (Golang)
- Gin Gonic Framework: https://github.com/gin-gonic/gin
- Gorm: https://gorm.io/index.html
- JWT Go: https://github.com/golang-jwt/jwt
- Go Validator: https://github.com/asaskevich/govalidator
- MySQL, PostgreSQL, or another SQL database

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/okto16/task-5-pbi-btpns-Oktorino-Bagas-Aji-Sudarno.git
cd task-5-pbi-btpns-Oktorino-Bagas-Aji-Sudarno
```

2. Run the application:
```bash
go run main.go
