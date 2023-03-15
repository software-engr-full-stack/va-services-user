-- migrate:up
CREATE TABLE IF NOT EXISTS users (
  email TEXT
);

-- migrate:down
DROP TABLE IF EXISTS users;
