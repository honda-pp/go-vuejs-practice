CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(100) NOT NULL UNIQUE,
  password_hash CHAR(60) NOT NULL
);

INSERT INTO users (username, email, password_hash) VALUES ('testuser', 'testuser@example.com', 'password123');
