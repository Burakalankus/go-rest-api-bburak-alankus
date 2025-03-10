# Go REST API with JWT Authentication ⚡️

## 🚀 Project Overview

This project is a **RESTful API** built using **Go (Golang)** and the **Gin-Gonic** framework. It features **JWT authentication**, **PostgreSQL** integration, and **Redis** caching for better performance. The project is designed to be easily extendable and scalable, allowing you to build secure web applications.

This API implements a **Book Library Management System** with the following resources:

- **Authors**
- **Books**
- **Reviews**

---

## ✨ Features

- **JWT Authentication**: Secure API endpoints with token-based authentication.
- **PostgreSQL Integration**: Seamlessly integrated with PostgreSQL using GORM for data storage and management.
- **Redis Caching**: Optimized in-memory caching for faster response times.
- **Swagger Documentation**: Automatically generated and user-friendly API docs for easy usage.
- **Database Migrations**: Automatically applied using GORM to handle schema changes.

---

## ⚙️ Technologies Used

- **Go (Golang)**: The backend programming language.
- **Gin-Gonic**: Fast and lightweight web framework for building the REST API.
- **GORM**: ORM library for PostgreSQL database interaction.
- **PostgreSQL**: The relational database used for storing data.
- **Redis**: In-memory key-value store for caching.
- **JWT (JSON Web Tokens)**: Secure authentication method using tokens.
- **Swagger**: API documentation tool for ease of use.

---

## 💻 Installation

To get started with the project locally, follow the steps below:

### Prerequisites

Ensure you have the following installed on your machine:

- **Go (1.20 or higher)**
- **Docker** (Optional, but recommended for containerized services like PostgreSQL and Redis)
- **PostgreSQL** (For database setup)
- **Redis** (For caching)

### Clone the Repository

Clone the repository to your local machine:


git clone https://github.com/Burakalankus/go-rest-api-bburak-alankus.git
cd go-rest-api-bburak-alankus


## 📑 API Endpoints

### 🔐 Authentication

- **POST `/api/v1/auth/register`**: Register a new user and receive a JWT token.
  - 🔑 *Create a new user and generate a JWT token for secure access.*
  
- **POST `/api/v1/auth/login`**: Authenticate and get a JWT token.
  - 🔒 *Login and receive a JWT token for authentication.*

### 📚 Books

- **GET `/api/v1/books`**: Fetch all books.
  - 📝 *Get a list of all books available in the system.*
  
- **GET `/api/v1/books/{id}`**: Fetch a specific book by ID.
  - 🔍 *Get details of a specific book using its ID.*

- **POST `/api/v1/books`**: Create a new book.
  - ➕ *Add a new book to the system.*
  
- **PUT `/api/v1/books/{id}`**: Update an existing book.
  - ✏️ *Update the details of an existing book using its ID.*
  
- **DELETE `/api/v1/books/{id}`**: Delete a book.
  - 🗑️ *Remove a specific book from the system.*

### ✍️ Authors

- **GET `/api/v1/authors`**: Fetch all authors.
  - 🖋️ *Get a list of all authors in the system.*

- **GET `/api/v1/authors/{id}`**: Fetch a specific author by ID.
  - 🔍 *Get details of a specific author using their ID.*
  
- **POST `/api/v1/authors`**: Create a new author.
  - ➕ *Add a new author to the system.*

- **PUT `/api/v1/authors/{id}`**: Update an existing author.
  - ✏️ *Update the details of an existing author using their ID.*
  
- **DELETE `/api/v1/authors/{id}`**: Delete an author.
  - 🗑️ *Remove an author from the system.*

### 💬 Reviews

- **GET `/api/v1/reviews/book/{id}`**: Get all reviews for a specific book.
  - 📝 *Get all reviews related to a specific book.*
  
- **POST `/api/v1/reviews`**: Create a new review.
  - ➕ *Add a review for a specific book.*
  
- **PUT `/api/v1/reviews/{id}`**: Update an existing review.
  - ✏️ *Edit an existing review by its ID.*

- **DELETE `/api/v1/reviews/{id}`**: Delete a review.
  - 🗑️ *Remove a review from the system.*

### 🔒 Protected Routes (JWT Authentication Required)

- **GET `/api/v1/protected/books`**: Get all books (protected).
  - 🔑 *Fetch books only with valid authentication (JWT required).*

- **GET `/api/v1/protected/authors`**: Get all authors (protected).
  - 🔑 *Fetch authors only with valid authentication (JWT required).*

---

## 🔧 Database & Environment Setup

### 🗄️ PostgreSQL Database Configuration

Make sure to set up your PostgreSQL database connection using the **.env** file. Ensure the following fields are configured correctly:

```ini
# PostgreSQL Settings
DB_HOST=localhost
DB_PORT=5432
DB_USER=burakalankus  
DB_PASSWORD=bbaa
DB_NAME=go_rest_api
DB_SSLMODE=disable
REDIS_HOST=localhost
REDIS_PORT=6379
```


## 🚨 Error Handling

The error handling system is currently a work in progress. Some basic error handling has been implemented, but further improvements are necessary to standardize error messages and better handle edge cases.

### Known Issues:
- 🔴 **Inconsistent error responses** across some endpoints.
- ❗ **Error messages** can be more descriptive in certain cases.

These issues are actively being worked on, and improvements will be implemented in future releases.

## 🚀 Future Improvements

- **Comprehensive Error Handling**: 
  - A standardized error handling system with clear and actionable error messages to ensure consistency across the API.
  
- **API Monitoring & Metrics**: 
  - Integration with monitoring tools like **Prometheus** and **Grafana** for real-time performance tracking and observability.

- **Rate Limiting**: 
  - Implement rate limiting to prevent abuse and improve fairness, especially for high-frequency requests.

- **Unit Testing**: 
  - Add unit tests for core functionalities to ensure the stability, reliability, and performance of the application.

## ⚡ PostgreSQL Error Handling

PostgreSQL error handling is a key part of this API's development. At the moment, some PostgreSQL-related errors (e.g., database connectivity, query issues) are not handled gracefully. Work is being done to improve this, including more descriptive error messages when interacting with the database. 

Some issues include:
- 🚫 **Database connection errors** (e.g., incorrect database credentials or missing database) are not handled consistently.
- ⚠️ **Query errors** (such as invalid data formats or data type mismatches) need better error reporting.

Future updates will focus on improving error messages for PostgreSQL-related issues, ensuring that all errors are caught, and providing clearer feedback to the users.



