# URL Shortener API

A production-ready **URL shortening service** built with **Go**, featuring clean architecture, Redis caching, MySQL persistence, and deployment on Railway.

This project was built to demonstrate **backend engineering fundamentals**, including API design, database modeling, caching, logging, and deployment.

---

## 🌍 Base URL

- https://ul.up.railway.app
  All endpoints below are relative to this base URL.

---

## ✨ Features

- Shorten long URLs into compact, shareable links
- Redirect short URLs to original destinations
- Redis caching for fast lookups
- MySQL for persistent storage
- Structured logging
- Environment-based configuration
- Production deployment on Railway

---

## 🛠 Tech Stack

- **Language:** Go
- **HTTP Framework:** Gin
- **Database:** MySQL
- **Cache:** Redis
- **Deployment:** Railway
- **Formatting:** gofmt

---

## 📐 Architecture

The project follows a clean, layered structure:

```
cmd/
  server/        # Application entry point
  handler/       # HTTP handlers (controllers)
  service/       # Business logic
  repository/    # Database access
  middleware/    # authmiddleware / request id / rate limiting
  helpers/       # helper functions
config/          # Environment & configuration & database set up
```

---

## 🚀 API Endpoints

````

### Create Short URL

```http
POST /
````

**Request body**

```json
{
  "long_url": "https://example.com",
  "alias": "omar" // optional
}
```

**Response**

```json
{
  "short_url": "https://ul.up.railway.app/omar",
  "success": true,
  "message": "URL created successfully"
}
```

---

### Redirect

```http
GET /{short_url}
```

Redirects to the original URL and increments click count.

---

### Create User

```http
POST /auth/signup
```

**Request body**

```json
{
  "email": "example1@email.com",
  "password": "password",
  "first_name": "John",
  "last_name": "Doe"
}
```

**Response**

```json
{
  "message": "Signup successful"
}
```

### Sign in

```http
POST /auth/signin
```

**Request body**

```json
{
  "email": "example@email.com",
  "password": "password"
}
```

**Response**

```json
{
  "access_token": "JWT Token",
  "message": "Login successful"
  // set cookie for refresh token
}
```

### Generate new access token

```http
POST /auth/refresh
```

**header**
Authorization: Bearer <access_token>
**Response**

```json
{
  "token": "JWT TOKEN"
}
```

### Sign out

```http
POST /auth/signout
```

**header**
Authorization: Bearer <access_token>
**Response**

```json
{
  "message": "Logout successful"
}
```

## 🌍 Deployment

The application is deployed on **Railway** with managed MySQL and Redis services.
