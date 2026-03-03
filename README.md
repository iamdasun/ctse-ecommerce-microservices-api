# CTSE E-Commerce Microservices API

A modern microservices-based e-commerce platform built with Kong API Gateway for routing and load balancing. This project demonstrates microservices architecture patterns using different technology stacks for each service.

## 🏗️ System Architecture

```
                    ┌─────────────────┐
                    │   Kong Gateway  │
                    │   (Port 8080)   │
                    └────────┬────────┘
                             │
           ┌─────────────────┼─────────────────┐
           │                 │                 │
           ▼                 ▼                 ▼
    ┌──────────────┐  ┌──────────────┐  ┌──────────────┐
    │ Item Service │  │ Order Service│  │Payment Service│
    │   Node.js    │  │     Go       │  │    Python    │
    │   Port 8081  │  │   Port 8082  │  │   Port 8083  │
    │   SQLite     │  │    SQLite    │  │    SQLite    │
    └──────────────┘  └──────────────┘  └──────────────┘
```

## 📦 Services

### 1. Kong API Gateway
- **Technology**: Kong 3.4
- **Port**: 8080 (proxy), 8001 (admin)
- **Role**: API routing, load balancing, and request management
- **Configuration**: Declarative config via `kong/kong.yml`

### 2. Item Service
- **Technology**: Node.js + Express + Sequelize
- **Port**: 8081
- **Database**: SQLite
- **Endpoints**:
  - `GET /items` - List all items
  - `GET /items/:id` - Get item by ID
  - `POST /items` - Create new item

### 3. Order Service
- **Technology**: Go + Gin + GORM
- **Port**: 8082
- **Database**: SQLite
- **Endpoints**:
  - `GET /orders` - List all orders
  - `GET /orders/:id` - Get order by ID
  - `POST /orders` - Create new order

### 4. Payment Service
- **Technology**: Python + Flask + Flask-SQLAlchemy
- **Port**: 8083
- **Database**: SQLite
- **Endpoints**:
  - `GET /payments` - List all payments
  - `GET /payments/:id` - Get payment by ID
  - `POST /payments/process` - Process new payment

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd ctse-ecommerce-microservices-api
   ```

2. **Start all services**
   ```bash
   docker compose up --build
   ```

3. **Verify services are running**
   ```bash
   # Test Kong Gateway
   curl http://localhost:8080/items
   
   # Test Item Service directly
   curl http://localhost:8081/items
   
   # Test Order Service directly
   curl http://localhost:8082/orders
   
   # Test Payment Service directly
   curl http://localhost:8083/payments
   ```

### Stopping Services
```bash
docker compose down
```

## 📡 API Usage

### Via Kong Gateway (Recommended)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `http://localhost:8080/items` | Get all items |
| GET | `http://localhost:8080/items/1` | Get item by ID |
| POST | `http://localhost:8080/items` | Create item |
| GET | `http://localhost:8080/orders` | Get all orders |
| GET | `http://localhost:8080/orders/1` | Get order by ID |
| POST | `http://localhost:8080/orders` | Create order |
| GET | `http://localhost:8080/payments` | Get all payments |
| GET | `http://localhost:8080/payments/1` | Get payment by ID |
| POST | `http://localhost:8080/payments/process` | Process payment |

### Example Requests

**Create Item:**
```bash
curl -X POST http://localhost:8080/items \
  -H "Content-Type: application/json" \
  -d '{"name": "Wireless Headphones", "price": 79.99}'
```

**Create Order:**
```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"itemId": 1, "quantity": 2, "userId": "user123"}'
```

**Process Payment:**
```bash
curl -X POST http://localhost:8080/payments/process \
  -H "Content-Type: application/json" \
  -d '{"amount": 199.99, "orderId": 1}'
```

## 🧪 Testing with Bruno/Postman

The project includes API collections for testing:

### Bruno Collections
Located in `api-collection/` directory:
- **Item Service**: CRUD operations for items
- **Order Service**: CRUD operations for orders
- **Payment Service**: Payment processing operations

### Environment Variables
The `api-collection/environments/local.yml` contains:
- `base_url`: `http://localhost:8080` (Kong Gateway)
- `direct_base_url`: `http://localhost`
- `item_port`: `8081`
- `order_port`: `8082`
- `payment_port`: `8083`
- `kong_admin_port`: `8001`

## 🔧 Configuration

### Kong Routing Rules
Kong is configured via `kong/kong.yml` with the following routing:

| Incoming Path | Target Service | Internal URL |
|---------------|----------------|--------------|
| `/items/*` | Item Service | `http://item-service:8081` |
| `/orders/*` | Order Service | `http://order-service:8082` |
| `/payments/*` | Payment Service | `http://payment-service:8083` |

### Docker Network
All services communicate via the `ecommerce-network` bridge network.

### Data Persistence
Each service has a dedicated Docker volume for data persistence:
- `item-data`: Item Service database
- `order-data`: Order Service database
- `payment-data`: Payment Service database

## 🏥 Health Checks

Each service includes health check configuration:

| Service | Health Check Endpoint | Interval | Timeout |
|---------|----------------------|----------|---------|
| Item Service | `GET /items` | 10s | 5s |
| Order Service | `GET /orders` | 10s | 5s |
| Payment Service | `GET /payments` | 10s | 5s |

## 📁 Project Structure

```
ctse-ecommerce-microservices-api/
├── api-collection/           # Bruno API collections
│   ├── environments/
│   │   └── local.yml
│   ├── Item Service/
│   ├── Order Service/
│   └── Payment Service/
├── kong/
│   └── kong.yml              # Kong declarative config
├── item-service/
│   ├── Dockerfile
│   ├── index.js
│   └── package.json
├── order-service/
│   ├── Dockerfile
│   ├── main.go
│   └── go.mod
├── payment-service/
│   ├── Dockerfile
│   ├── app.py
│   └── requirements.txt
├── docker-compose.yml
├── README.md
└── LICENSE
```

## 🛠️ Development

### Building Individual Services

```bash
# Build Item Service
docker build -t item-service ./item-service

# Build Order Service
docker build -t order-service ./order-service

# Build Payment Service
docker build -t payment-service ./payment-service
```

### Running Individual Services

```bash
# Run Item Service
docker run -p 8081:8081 item-service

# Run Order Service
docker run -p 8082:8082 order-service

# Run Payment Service
docker run -p 8083:8083 payment-service
```

## 🔍 Kong Admin API

Access Kong's admin API for management:

```bash
# Get all routes
curl http://localhost:8001/routes

# Get all services
curl http://localhost:8001/services

# Get service configuration
curl http://localhost:8001/services/order-service
```

## 📝 Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.

## 📄 License

This project is licensed under the GNU General Public License v3.0 (GPL-3.0). See the [LICENSE](LICENSE) file for details.

## 🤝 Authors

- CTSE Lab 06 - E-Commerce Microservices

## 📚 Technologies Used

| Service | Framework | Database | Language |
|---------|-----------|----------|----------|
| Kong Gateway | Kong 3.4 | - | Lua/Nginx |
| Item Service | Express.js | SQLite | Node.js |
| Order Service | Gin | SQLite | Go |
| Payment Service | Flask | SQLite | Python |

---

**Note**: This is a demonstration project for educational purposes. For production use, consider implementing authentication, proper error handling, logging, monitoring, and using production-grade databases.
