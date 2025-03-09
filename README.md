# Chat System README

## Overview

This project is a **chat system** built using **Ruby on Rails** for the backend API, **Golang** for additional RESTful services, **Redis** for caching, **Sidekiq** for background job processing, and **Elasticsearch** for message searching. The system is containerized using **Docker Compose**.

## Tech Stack

- **Ruby on Rails**: Backend API
- **Golang**: RESTful API layer
- **MySQL**: Relational database
- **Redis**: Caching and queue management
- **Sidekiq**: Background job processing
- **Elasticsearch**: Full-text search engine
- **Docker Compose**: Containerized environment

---

## Setup and Installation

### **1. Build and Start the Project**

To build and run all services, use the following command:

```sh
docker-compose up --build
```

This will start:

- **MySQL Database** (Port: `3306`)
- **Redis** (Port: `6379`)
- **Rails Backend API** (Port: `3000`)
- **Golang API** (Port: `8080`)
- **Elasticsearch** (Port: `9200`)
- **Sidekiq Worker** (Processes background jobs)

If you want to run it in detached mode:

```sh
docker-compose up -d --build
```

---

## Available Endpoints

### **Applications**

- `POST /applications` - Create a new application
- `GET /applications/:token` - Retrieve application details

### **Chats**

- `POST /applications/:token/chats` - Create a chat within an application
- `GET /applications/:token/chats/:chat_number` - Retrieve chat details

### **Messages**

- `POST /applications/:token/chats/:chat_number/messages` - Send a message
- `POST /applications/:token/chats/:chat_number/messages/search` - Search for messages

---

## **Monitoring Sidekiq**

Sidekiq can be monitored via a web UI at:

```
http://localhost:3000/sidekiq
```

If CSRF issues arise, ensure session middleware is enabled for API mode.

To check Sidekiq logs:

```sh
docker-compose logs -f sidekiq
```

---

## **Debugging and Logs**

To view logs for a specific service:

```sh
docker-compose logs -f backend
```

```sh
docker-compose logs -f golang_api
```

To access the Rails backend container shell:

```sh
docker-compose exec backend sh
```

To access the Golang API container shell:

```sh
docker-compose exec golang_api sh
```

To restart a specific service:

```sh
docker-compose restart backend
```

---

## **Project Structure**

### **Where can we find your database schema?**

The database schema can be found in the **Rails project** under:

```
backend/db/schema.rb
```

### **Where can we find each API controller you implemented?**

All API controllers are located in the **Rails project** under:

```
backend/app/controllers/
```

### **Where can we find workers you implemented (if any)?**

Sidekiq jobs can be found in:

```
backend/app/jobs/
```

---

## **Database Indexes**

Indexes that were added for performance optimization:

- **Applications**:
  - `token (unique)`
- **Chats**:
  - `number (unique per application)`
- **Messages**:
  - `chat_id`
  - `number (unique per chat)`
  - `body (full-text index for search using Elasticsearch)`

### **What does "full index for search" mean?**

The **Messages** table uses **Elasticsearch** for full-text search indexing. Unlike traditional database indexes, Elasticsearch provides:

- **Fuzzy matching** for handling typos and variations.
- **Relevance scoring** to prioritize better results.
- **Fast text-based searching** across large datasets.

This improves the speed and accuracy of search queries over the **body** column in messages.

---

## **Things We Could Improve On**

- Implement authentication for API endpoints
- Improve error handling for concurrent requests
- Enhance search capabilities with Elasticsearch tuning
- Optimize database indexing for better performance
- Introduce rate limiting to prevent spam requests

