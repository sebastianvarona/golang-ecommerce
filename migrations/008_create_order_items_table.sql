CREATE TABLE order_items (
  id SERIAL PRIMARY KEY,
  order_id INT REFERENCES orders(id) ON DELETE CASCADE,
  product_id INT REFERENCES products(id) ON DELETE SET NULL,
  variant_id INT REFERENCES variants(id) ON DELETE SET NULL,
  size_id INT REFERENCES sizes(id) ON DELETE SET NULL,
  quantity INT NOT NULL,
  price NUMERIC(10, 2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
