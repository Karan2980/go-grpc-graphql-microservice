# 🚀 Go gRPC GraphQL Microservice

A complete microservice architecture built with **Go**, **gRPC**, **GraphQL**, **PostgreSQL**, and **Elasticsearch**. This project demonstrates modern microservice patterns including service-to-service communication, database per service, and API gateway patterns.

## 🏗️ Architecture Overview

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   GraphQL       │    │   Account        │    │   Catalog       │
│   Gateway       │◄──►│   Service        │    │   Service       │
│   :8083         │    │   :8080          │    │   :8081         │
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │                       │
         │              ┌─────────────────┐              │
         └─────────────►│   Order         │◄─────────────┘
                        │   Service       │
                        │   :8082         │
                        └─────────────────┘
                                 │
                    ┌─────────────────────────────┐
                    │                             │
            ┌───────▼──────┐            ┌────────▼────────┐
            │ PostgreSQL   │            │ Elasticsearch   │
            │ (Account &   │            │ (Catalog)       │
            │  Order DBs)  │            │                 │
            └──────────────┘            └─────────────────┘
```

## 🎯 Features

### **Services**
- **GraphQL Gateway** - Single entry point for all client requests
- **Account Service** - User account management
- **Catalog Service** - Product catalog with search capabilities
- **Order Service** - Order processing with multi-service validation

### **Technologies**
- **Backend**: Go 1.21+
- **Communication**: gRPC for inter-service communication
- **API Gateway**: GraphQL for client-facing API
- **Databases**: PostgreSQL for transactional data, Elasticsearch for search
- **Containerization**: Docker & Docker Compose
- **Code Generation**: Protocol Buffers, GraphQL Code Generator

### **Patterns Implemented**
- Microservice Architecture
- Database per Service
- API Gateway Pattern
- Service Discovery
- Circuit Breaker (via gRPC)
- Distributed Transactions
- Event-Driven Architecture

## 📋 Prerequisites

- **Docker** 20.10+
- **Docker Compose** 2.0+
- **Go** 1.21+ (for development)
- **Make** (optional, for build automation)



## 🚀 Quick Start

### 1. Clone the Repository
```bash
git clone https://github.com/Karan2980/go-grpc-graphql-microservice.git
cd go-grpc-graphql-microservice
```

### 2. Start All Services
```bash
docker-compose up -d
```

### 3. Verify Services are Running
```bash
# Check all containers
docker-compose ps

# Check logs
docker-compose logs -f
```

### 4. Access GraphQL Playground
Open your browser and navigate to:
```
http://localhost:8083/graphql
```

## 🔌 Service Endpoints

| Service | Port | Protocol | Health Check |
|---------|------|----------|--------------|
| GraphQL Gateway | 8083 | HTTP/GraphQL | `http://localhost:8083/graphql` |
| Account Service | 8080 | gRPC | `grpcurl -plaintext localhost:8080 list` |
| Catalog Service | 8081 | gRPC | `grpcurl -plaintext localhost:8081 list` |
| Order Service | 8082 | gRPC | `grpcurl -plaintext localhost:8082 list` |
| PostgreSQL (Account) | 5432 | SQL | `psql -h localhost -U postgres -d account` |
| PostgreSQL (Order) | 5433 | SQL | `psql -h localhost -U postgres -d order` |
| Elasticsearch | 9200 | HTTP | `curl http://localhost:9200/_cluster/health` |

## 📖 API Documentation

### **GraphQL Schema**

#### **Queries**
```graphql
type Query {
  # Account queries
  account(id: ID!): Account
  accounts(pagination: PaginationInput): [Account!]!
  
  # Product queries
  product(id: ID!): Product
  products(pagination: PaginationInput, query: String): [Product!]!
}
```

#### **Mutations**
```graphql
type Mutation {
  # Account mutations
  createAccount(account: AccountInput!): Account!
  
  # Product mutations
  createProduct(product: ProductInput!): Product!
  
  # Order mutations
  createOrder(order: OrderInput!): Order!
}
```

#### **Types**
```graphql
type Account {
  id: ID!
  name: String!
  orders: [Order!]!
}

type Product {
  id: ID!
  name: String!
  description: String!
  price: Float!
}

type Order {
  id: ID!
  createdAt: String!
  totalPrice: Float!
  products: [OrderedProduct!]!
}

type OrderedProduct {
  id: ID!
  name: String!
  description: String!
  price: Float!
  quantity: Int!
}
```

## 🧪 API Examples

### **1. Create Account**
```bash
curl -X POST http://localhost:8083/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createAccount(account: { name: \"John Doe\" }) { id name } }"
  }'
```

### **2. Create Products**
```bash
curl -X POST http://localhost:8083/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createProduct(product: { name: \"iPhone 15\", description: \"Latest iPhone\", price: 999.99 }) { id name price } }"
  }'
```

### **3. Create Order**
```bash
curl -X POST http://localhost:8083/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createOrder(order: { accountId: \"acc_123\", products: [{ id: \"prod_456\", quantity: 2 }] }) { id totalPrice products { name quantity } } }"
  }'
```

### **4. Get Account with Orders**
```bash
curl -X POST http://localhost:8083/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "{ account(id: \"acc_123\") { id name orders { id totalPrice products { name quantity } } } }"
  }'
```

### **5. Search Products**
```bash
curl -X POST http://localhost:8083/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "{ products(query: \"iPhone\") { id name description price } }"
  }'
```

## 🏗️ Development Setup

### **1. Install Dependencies**
```bash
# Install Go dependencies
go mod download

# Install development tools
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### **2. Generate Code**
```bash
# Generate protobuf code
make proto

# Generate GraphQL code
make graphql
```

### **3. Run Individual Services**
```bash
# Start databases first
docker-compose up -d account_db order_db elasticsearch

# Run Account Service
cd account && go run cmd/account/main.go

# Run Catalog Service
cd catalog && go run cmd/catalog/main.go

# Run Order Service
cd order && go run cmd/order/main.go

# Run GraphQL Gateway
cd graphql && go run cmd/graphql/main.go
```

## 📁 Project Structure

```
├── account/                 # Account microservice
│   ├── cmd/account/        # Service entry point
│   ├── pb/                 # Generated protobuf code
│   ├── client.go           # gRPC client
│   ├── server.go           # gRPC server
│   ├── service.go          # Business logic
│   ├── repository.go       # Data access layer
│   ├── Dockerfile          # Container definition
│   └── up.sql             # Database schema
├── catalog/                # Catalog microservice
│   ├── cmd/catalog/       # Service entry point
│   ├── pb/                # Generated protobuf code
│   ├── client.go          # gRPC client
│   ├── server.go          # gRPC server
│   ├── service.go         # Business logic
│   ├── repository.go      # Data access layer
│   └── Dockerfile         # Container definition
├── order/                  # Order microservice
│   ├── cmd/order/         # Service entry point
│   ├── pb/                # Generated protobuf code
│   ├── client.go          # gRPC client
│   ├── server.go          # gRPC server
│   ├── service.go         # Business logic
│   ├── repository.go      # Data access layer
│   ├── Dockerfile         # Container definition
│   └── up.sql            # Database schema
├── graphql/               # GraphQL gateway
│   ├── cmd/graphql/      # Gateway entry point
│   ├── schema.graphql    # GraphQL schema
│   ├── resolver.go       # GraphQL resolvers
│   └── Dockerfile        # Container definition
├── docker-compose.yaml   # Multi-container setup
├── go.mod               # Go module definition
└── README.md           # This file
```

## 🔧 Configuration

### **Environment Variables**

#### **Account Service**
```env
DATABASE_URL=postgres://postgres:password@account_db:5432/account?sslmode=disable
```

#### **Catalog Service**
```env
DATABASE_URL=http://elasticsearch:9200
```

#### **Order Service**
```env
DATABASE_URL=postgres://postgres:password@order_db:5432/order?sslmode=disable
ACCOUNT_SERVICE_URL=account:8080
CATALOG_SERVICE_URL=catalog:8080
```

#### **GraphQL Gateway**
```env
ACCOUNT_SERVICE_URL=account:8080
CATALOG_SERVICE_URL=catalog:8080
ORDER_SERVICE_URL=order:8080
```

## 🐳 Docker Configuration

### **Build Individual Services**
```bash
# Build Account Service
docker build -f account/Dockerfile -t account-service .

# Build Catalog Service
docker build -f catalog/Dockerfile -t catalog-service .

# Build Order Service
docker build -f order/Dockerfile -t order-service .

# Build GraphQL Gateway
docker build -f graphql/Dockerfile -t graphql-gateway .
```

### **Docker Compose Commands**
```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f [service-name]

# Stop all services
docker-compose down

# Rebuild and restart
docker-compose up -d --build

# Remove all data
docker-compose down -v
```

## 🔍 Monitoring & Debugging

### **Health Checks**
```bash
# Check GraphQL Gateway
curl http://localhost:8083/graphql

# Check Elasticsearch
curl http://localhost:9200/_cluster/health

# Check PostgreSQL (Account)
docker exec -it go-grpc-graphql-micro-account_db-1 psql -U postgres -d account -c "SELECT COUNT(*) FROM accounts;"

# Check PostgreSQL (Order)
docker exec -it go-grpc-graphql-micro-order_db-1 psql -U postgres -d order -c "SELECT COUNT(*) FROM orders;"
```

### **View Logs**
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f account
docker-compose logs -f catalog
docker-compose logs -f order
docker-compose logs -f graphql
```

### **Database Access**
```bash
# Account Database
docker exec -it go-grpc-graphql-micro-account_db-1 psql -U postgres -d account

# Order Database
docker exec -it go-grpc-graphql-micro-order_db-1 psql -U postgres -d order

# Elasticsearch
curl -X GET "localhost:9200/catalog/_search?pretty"
```
