CREATE TABLE shipping (
  id SERIAL PRIMARY KEY,
  order_id INT REFERENCES orders(id) ON DELETE CASCADE,
  method VARCHAR(100),
  tracking_number VARCHAR(100),
  status VARCHAR(50) DEFAULT 'preparing',
  estimated_delivery TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
