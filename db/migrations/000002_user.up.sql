CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    name VARCHAR NOT NULL,
    lastname VARCHAR NOT NULL, 
    password VARCHAR NOT NULL,
    phone_number VARCHAR DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_tokens(
    user_id UUID, --UNIQUE 
    refresh_token VARCHAR NOT NULL,
    access_token VARCHAR NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);