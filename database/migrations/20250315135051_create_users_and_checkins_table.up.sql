-- create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    did VARCHAR(100) UNIQUE NOT NULL,
    handle VARCHAR(100) UNIQUE NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    avatar TEXT,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- create checkins table
CREATE TABLE IF NOT EXISTS checkins (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user_did VARCHAR(100) NOT NULL REFERENCES users(did),
    text TEXT,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    location_name VARCHAR(100) NOT NULL,
    images JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- create index
CREATE INDEX IF NOT EXISTS idx_checkins_user_id ON checkins(user_id);
CREATE INDEX IF NOT EXISTS idx_checkins_created_at ON checkins(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_checkins_user_did ON checkins(user_did);
CREATE INDEX IF NOT EXISTS idx_checkins_location_name ON checkins(location_name);