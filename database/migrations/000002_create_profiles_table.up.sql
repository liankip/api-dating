CREATE TABLE profiles (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  username VARCHAR(50) NOT NULL,
  bio TEXT,
  photo_url VARCHAR(255)
);