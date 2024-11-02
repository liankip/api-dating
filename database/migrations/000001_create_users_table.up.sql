CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   username VARCHAR(50) NOT NULL UNIQUE,
   password_hash VARCHAR(255) NOT NULL,
   is_premium BOOLEAN DEFAULT FALSE,
   verified_badge BOOLEAN DEFAULT FALSE,
   swipe_quota INT DEFAULT 10,
   quota_reset_time TIMESTAMP
);