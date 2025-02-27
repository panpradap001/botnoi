# Cat Cafe API

## Overview

Cat Cafe API is a RESTful API built with Golang and MongoDB, designed to manage cat-related information in a pet cafe. It allows users to create, retrieve, update, and delete cat records.

## Technologies Used

- **Golang** (Gin Framework)
- **MongoDB** (NoSQL Database)
- **Docker & Docker Compose**
- **Postman** (API Testing)

## Getting Started

### Prerequisites

Ensure you have the following installed:

- Docker & Docker Compose
- Postman (optional for testing API)

### Installation & Setup

1. **Clone the repository**

   ```sh
   git clone https://github.com/your-repo/cat_cafe.git
   cd cat_cafe
   ```

2. **Start the application using Docker Compose**

   ```sh
   docker-compose up --build
   ```

3. **Verify the running containers**
   ```sh
   docker ps
   ```
   You should see `cat_cafe_api` and `mongodb` running.

## API Endpoints

### 1. Get all cats

**Request:**

```http
GET /cats
```

**Response:**

```json
[
  {
    "name": "Mochi",
    "breed": "Scottish Fold",
    "age": 2,
    "favorite_food": "Salmon",
    "status": "Available for adoption"
  }
]
```

### 2. Add a new cat

**Request:**

```http
POST /cats
```

**Headers:**

```json
{
  "Content-Type": "application/json"
}
```

**Body:**

```json
{
  "name": "Luna",
  "breed": "Persian",
  "age": 3,
  "favorite_food": "Tuna",
  "status": "Adopted"
}
```

**Response:**

```json
{
  "message": "Cat added"
}
```

### 3. Delete a cat (Future Feature)

**Request:**

```http
DELETE /cats/:id
```

## Database Setup (MongoDB)

If you want to manually check the database, use:

```sh
docker exec -it mongodb mongosh
```

To check collections:

```sh
show dbs;
use cat_cafe;
show collections;
db.cats.find().pretty();
```

## Troubleshooting

- **Error: Cannot connect to MongoDB**
  ```sh
  docker-compose down -v
  docker-compose up --build
  ```
- **API not responding on expected port**
  ```sh
  docker ps
  ```
  Ensure the API is running on the correct port (`9090`).
