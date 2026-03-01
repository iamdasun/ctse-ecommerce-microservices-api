# Item Service

Node.js + Express microservice for managing items in the CTSE E-Commerce platform.

## Tech Stack

- **Runtime:** Node.js 18
- **Framework:** Express.js
- **Container:** Docker (node:18-alpine)

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/items` | Return all items |
| GET | `/items/:id` | Return item by index |
| POST | `/items` | Add a new item |

## Usage

### Local Development

```bash
npm install
npm start
```

### Docker

```bash
docker build -t item-service .
docker run -p 8081:8081 item-service
```

## Example Requests

### Get All Items
```bash
curl http://localhost:8081/items
```

### Get Item by ID
```bash
curl http://localhost:8081/items/0
```

### Add New Item
```bash
curl -X POST http://localhost:8081/items \
  -H "Content-Type: application/json" \
  -d '{"name": "Tablet"}'
```
