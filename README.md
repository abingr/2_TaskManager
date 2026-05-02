# Go Task Manager API

A production-style backend Task Manager REST API built with Go, Gin, PostgreSQL, Redis, JWT Authentication, and Docker.

This project demonstrates how to build a secure, multi-user, optimized backend service with:

- User registration and login
- JWT protected routes
- PostgreSQL persistent storage
- Redis cache-aside optimization
- Cache invalidation after writes
- Rate limiting middleware
- Dockerized infrastructure

---

## Tech Stack

- Golang
- Gin Web Framework
- PostgreSQL
- Redis
- Docker / Docker Compose
- JWT Authentication

---

## Features

### Authentication
- Register new users
- Login existing users
- JWT token generation
- Protected API endpoints

### Task Management
- Create user-specific tasks
- Retrieve only authenticated user's tasks

### Performance Optimization
- Redis caching for GET /tasks
- 5-minute TTL cache storage
- Automatic cache invalidation on POST /tasks

### Security / Stability
- Rate limit middleware (5 requests / 10 sec)
- Protected middleware chaining

---

## API Endpoints

### Public
- POST /register
- POST /login

### Protected
- GET /tasks
- POST /tasks

---

## Run Locally

### Start PostgreSQL + Redis

```bash
docker compose up -d