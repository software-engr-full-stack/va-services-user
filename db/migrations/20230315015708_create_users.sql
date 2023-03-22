-- migrate:up
CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  email TEXT
);

-- migrate:down
DROP TABLE IF EXISTS users;
