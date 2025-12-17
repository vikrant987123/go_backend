
## Overview

This project is a **Go-based backend REST API** for managing users.  
The API accepts a user’s **name** and **date of birth (DOB)**, stores them in a PostgreSQL database, and **calculates the user’s age dynamically** when returning responses.

The application is designed using **clean architecture principles**, focusing on clarity, maintainability, and production-ready practices.

---

## Tech Stack

- **Go** (Fiber framework)
- **PostgreSQL** with **SQLC** for type-safe database access
- **go-playground/validator** for request validation
- **Uber Zap** for structured logging

---

## Problem Understanding

The core requirements were:
- Store user information (`name`, `dob`)
- Expose RESTful CRUD APIs
- Calculate and return user age dynamically
- Use SQL with PostgreSQL and SQLC
- Validate incoming requests
- Provide proper logging and HTTP status codes

A key challenge was ensuring that **derived data (age)** is handled correctly without introducing data inconsistency.

---

## Functionality

The API supports the following operations:

- Create a user
- Get user by ID
- Update user by ID
- Delete user by ID
- List all users (with pagination)

### Key Behaviors
- Age is calculated dynamically at runtime using the DOB.
- Pagination is supported via `page` and `limit` query parameters.
- Proper HTTP status codes are returned for all responses.
- Requests are logged with tracing information.

---

## Request Flow
Client
→ Server
→ Request ID Middleware
→ Logging Middleware (Zap)
→ Route
→ Handler
→ Service
→ SQLC (PostgreSQL)
→ Response sent back to client


---

## Design Decisions

### 1️⃣ Dynamic Age Calculation

- The client provides DOB in `YYYY-MM-DD` format.
- DOB is stored in the database.
- Age is calculated dynamically using Go’s `time` package.

---

### 2️⃣ Pagination Strategy

- Pagination is implemented using `page` and `limit`.
- If invalid values are provided, defaults are applied:
  - `page = 1`
  - `limit = 10`

**Reasoning:**
- Prevents misuse of the API

---

### 3️⃣  Middleware

Two middlewares are implemented:
- **Request ID Middleware**: assigns a unique ID to each request and exposes it via headers.
- **Logging Middleware**: logs method, path, status, duration, and request ID using Zap.

**Reasoning:**
- Enables request tracing
- Simplifies debugging

---




