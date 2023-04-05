CREATE TABLE IF NOT EXISTS genesis.exchanges (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  amount DECIMAL(15, 2) NOT NULL,
  from_currency INTEGER NOT NULL,
  to_currency INTEGER NOT NULL,
  rate DECIMAL(15, 2) NOT NULL,
  FOREIGN KEY (from_currency) REFERENCES currency(id),
  FOREIGN KEY (to_currency) REFERENCES currency(id)
)