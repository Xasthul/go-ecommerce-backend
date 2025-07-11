# Go Microservices Backend Project

![Go](https://img.shields.io/badge/Go-1.20-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14-blue)
![RabbitMQ](https://img.shields.io/badge/RabbitMQ-3.9-orange)
![Redis](https://img.shields.io/badge/Redis-7.0-red)
![Gin](https://img.shields.io/badge/Gin-Framework-green)

## Project Overview

This project is a fully-featured microservices backend ecommerce implementation written in Go. It demonstrates key backend concepts such as user authentication, role-based authorization, inter-service communication, event-driven architecture, and caching.

The system includes:

- **Auth Service:** Handles user registration, login, JWT-based authentication, and role management.
- **API Gateway:** Central entry point routing requests, validating JWTs, and injecting trusted headers.
- **Product Service:** Manages product catalog and stock.
- **Order Service:** Responsible for order processing, including validation, stock reservation via messaging, and order item management.
- **Payment Service:** Receives orders and creates payments.

---

## Key Technologies

| Technology       | Purpose                                 |
|------------------|-----------------------------------------|
| Go               | Backend language                        |
| Gin              | HTTP web framework                      |
| PostgreSQL       | Database                                |
| RabbitMQ         | Asynchronous event messaging            |
| Redis            | Caching layer and session storage       |
| Docker & Compose | Containerization and orchestration      |
| sqlc             | Type-safe SQL query generation          |

---

## Architecture Highlights

- **Microservices**: Services are loosely coupled and communicate asynchronously via RabbitMQ.
- **Authentication & Authorization**: JWT tokens include user roles; API Gateway validates tokens and enforces access control.
- **Event-Driven**: Order creation triggers events for stock reservation in Product Service.
- **Caching**: Redis is used for caching frequently accessed data to optimize performance.
- **Database Migrations**: Migrations are run at service startup ensuring schema consistency.
- **Secure Internal Communication**: Internal endpoints protected by API keys; sensitive environment variables managed via `.env` files.
