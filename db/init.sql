CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE CHECK (username <> ''),
  email VARCHAR(100) NOT NULL UNIQUE CHECK (email <> ''),
  password_hash CHAR(60) NOT NULL CHECK (password_hash <> '')
);

CREATE TABLE login_history (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  login_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE follower (
  follower_id INTEGER NOT NULL,
  followee_id INTEGER NOT NULL,
  PRIMARY KEY (follower_id, followee_id),
  FOREIGN KEY (follower_id) REFERENCES users (id) ON DELETE CASCADE,
  FOREIGN KEY (followee_id) REFERENCES users (id) ON DELETE CASCADE
);

INSERT INTO users (username, email, password_hash) VALUES ('testuser', 'testuser@example.com', '$2a$10$b6Nk3aziRYO9mVnMGnSSQuI6cj1hyYezlHce1RRbNYY4Tmjlw//Cy');
--psql -U golang_usr -d golang_db -h localhost -p 5432
