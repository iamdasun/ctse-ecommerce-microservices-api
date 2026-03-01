const express = require('express');

const app = express();
const PORT = 8081;

// Middleware
app.use(express.json());

// In-memory data store
let items = ["Book", "Laptop", "Phone"];

// GET /items - Return all items
app.get('/items', (req, res) => {
  res.json(items);
});

// GET /items/:id - Return item by index
app.get('/items/:id', (req, res) => {
  const id = parseInt(req.params.id);

  if (id < 0 || id >= items.length) {
    return res.status(404).json({ error: 'Item not found' });
  }

  res.json(items[id]);
});

// POST /items - Add new item
app.post('/items', (req, res) => {
  const { name } = req.body;

  if (!name || typeof name !== 'string') {
    return res.status(400).json({ error: 'Invalid item name' });
  }

  items.push(name);
  res.status(201).json({ id: items.length - 1, name });
});

// Start server
app.listen(PORT, () => {
  console.log(`Item Service running on port ${PORT}`);
});
