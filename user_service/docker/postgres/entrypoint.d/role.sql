CREATE TABLE IF NOT EXISTS roles (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR UNIQUE NOT NULL,
  description VARCHAR NOT NULL
);
