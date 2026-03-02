const express = require("express");
const { Sequelize, DataTypes } = require("sequelize");

const app = express();
const PORT = 8081;

// Middleware
app.use(express.json());

// Initialize SQLite with Sequelize
const sequelize = new Sequelize({
  dialect: "sqlite",
  storage: "./database.sqlite",
  logging: false,
});

// Define Item model
const Item = sequelize.define("Item", {
  name: {
    type: DataTypes.STRING,
    allowNull: false,
  },
});

// Initialize database and seed data
async function initializeDatabase() {
  await sequelize.sync();

  const count = await Item.count();
  if (count === 0) {
    await Item.bulkCreate([
      { name: "Book" },
      { name: "Laptop" },
      { name: "Phone" },
    ]);
  }
}

// GET /items - Return all items
app.get("/items", async (req, res) => {
  try {
    const items = await Item.findAll();
    res.json(items);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

// GET /items/:id - Return item by ID
app.get("/items/:id", async (req, res) => {
  try {
    const item = await Item.findByPk(req.params.id);

    if (!item) {
      return res.status(404).json({ error: "Item not found" });
    }

    res.json(item);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

// POST /items - Add new item
app.post("/items", async (req, res) => {
  try {
    const { name } = req.body;

    if (!name || typeof name !== "string") {
      return res.status(400).json({ error: "Invalid item name" });
    }

    const item = await Item.create({ name });
    res.status(201).json(item);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

// Start server
async function startServer() {
  await initializeDatabase();

  app.listen(PORT, () => {
    console.log(`Item Service running on port ${PORT}`);
  });
}

startServer();
