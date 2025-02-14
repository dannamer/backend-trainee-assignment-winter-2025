CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS merch_store (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    item VARCHAR(255) UNIQUE NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO merch_store (item, price) 
VALUES 
    ('t-shirt', 80.00),
    ('cup', 20.00),
    ('book', 50.00),
    ('pen', 10.00),
    ('powerbank', 200.00),
    ('hoody', 300.00),
    ('umbrella', 200.00),
    ('socks', 10.00),
    ('wallet', 50.00),
    ('pink-hoody', 500.00)
ON CONFLICT (item) DO NOTHING;

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR(32) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
## TODO: создать новую ошибку в openapi 409? Как ты думаешь, будущий я
CREATE OR REPLACE FUNCTION check_user_limit() 
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT COUNT(*) FROM users) > 100000 THEN
        RAISE EXCEPTION 'User limit of 100,000 has been reached';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_user_limit_trigger
BEFORE INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION check_user_limit();


CREATE TABLE IF NOT EXISTS wallet (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    balance DECIMAL(10,2) NOT NULL DEFAULT 1000.00 CHECK (balance >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    sender_id UUID REFERENCES users(id) ON DELETE CASCADE,
    receiver_id UUID REFERENCES users(id) ON DELETE CASCADE,
    amount DECIMAL(10,2) NOT NULL CHECK (amount > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS inventory (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    item VARCHAR(255) NOT NULL,
    quantity INTEGER DEFAULT 1 CHECK (quantity >= 1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT uq_item_user UNIQUE (user_id, item)
);

CREATE INDEX IF NOT EXISTS idx_inventory_user_id ON inventory(user_id);

CREATE INDEX IF NOT EXISTS idx_inventory_item ON inventory(item);