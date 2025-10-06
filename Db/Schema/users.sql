CREATE TABLE
    users (
        id BIGSERIAL PRIMARY KEY,
        email VARCHAR UNIQUE NOT NULL,
        name VARCHAR NOT NULL,
        password VARCHAR NOT NULL,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW ()
    );

CREATE TABLE
    profiles (
        id BIGSERIAL PRIMARY KEY,
        user_id BIGINT FOREIGN KEY,
        avatar_url TEXT,
        bio TEXT,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW (),
    );