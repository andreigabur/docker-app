const express = require('express');
const app = express();
const { Pool } = require('pg');

app.get('/', (req, res) => {
  res.send('Hello, World!');
});

// Set up a PostgreSQL connection pool
const pool = new Pool({
  user: 'myuser',
  host: 'db',
  database: 'mydatabase',
  password: 'mypassword',
  port: 5432
});

app.get('/getusers', async (req, res) => {
  try {
    // Retrieve user data from the "users" table
    const result = await pool.query('SELECT * FROM users');

    // Send the user data as a JSON response
    res.json(result.rows);
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Internal server error' });
  }
});


app.listen(3000, () => {
  console.log('Node server started on port 3000');
});
