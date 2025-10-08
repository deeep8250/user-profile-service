CREATE TABLE
    profiles (
        id BIGSERIAL PRIMARY KEY,
        user_id BIGINT REFERENCES users(id),
        avatar_url TEXT,
        bio TEXT,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW ()
    );