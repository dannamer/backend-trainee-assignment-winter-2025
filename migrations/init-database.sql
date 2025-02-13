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
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS wallet (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    balance DECIMAL(10,2) NOT NULL DEFAULT 1000.00 CHECK (balance >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    sender_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Кто отправил (NULL если пополнение)
    receiver_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Кто получил
    amount DECIMAL(10,2) NOT NULL CHECK (amount > 0), -- Сумма перевода/пополнения
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

SELECT 
    t.receiver_id, 
    u.username AS sender_username
FROM transactions t
JOIN users u ON u.id = t.sender_id
WHERE 'b7b245bd-965e-466d-9652-5fcece07f177' = t.sender_id
ORDER BY t.created_at DESC;

SELECT 
    u.username, 
    t.amount
FROM transactions t
JOIN users u ON t.receiver_id = u.id
WHERE t.receiver_id = 'b7b245bd-965e-466d-9652-5fcece07f177'
ORDER BY t.created_at DESC;


SELECT u.username, t.amount
FROM transactions t JOIN users u ON t.receiver_id = u.id
WHERE 
'b7b245bd-965e-466d-9652-5fcece07f177'
SELECT t.sender_id
    id, 
    sender_id, 
    receiver_id, 
    amount, 
    transaction_type, 
    created_at
FROM transactions
WHERE sender_id = 'user-uuid' OR receiver_id = 'user-uuid'
ORDER BY created_at DESC;

-- CREATE TABLE IF NOT EXISTS purchase_history (
--     id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
--     user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- Кто купил
--     merch_id UUID NOT NULL REFERENCES merch_store(id) ON DELETE CASCADE, -- Что купил
--     -- price DECIMAL(10,2) NOT NULL, -- Цена покупки
--     created_at TIMESTAMP DEFAULT NOW() NOT NULL
-- );

CREATE TABLE IF NOT EXISTS inventory (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    item VARCHAR(255) NOT NULL,
    quantity INTEGER DEFAULT 1 CHECK (quantity >= 1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT uq_item_user UNIQUE (user_id, item)
);

-- Индексы для повышения производительности при поиске по user_id и item
CREATE INDEX idx_inventory_user_id ON inventory(user_id);
CREATE INDEX idx_inventory_item ON inventory(item);
WITH new_user AS (
    INSERT INTO users (username, password) 
    VALUES ('khan', 'lol') 
    RETURNING id
)
INSERT INTO wallet (user_id)
VALUES ((SELECT id FROM new_user LIMIT 1))
RETURNING user_id;

-- CREATE OR REPLACE FUNCTION update_updated_at_column()
-- RETURNS TRIGGER AS $$
-- BEGIN
--   NEW.updated_at = CURRENT_TIMESTAMP;
--   RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER update_inventory_updated_at
-- BEFORE UPDATE ON inventory
-- FOR EACH ROW
-- EXECUTE FUNCTION update_updated_at_column();
-- lol
-- eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk1MDMwNDgsImlhdCI6MTczOTQxNjY0OCwic3ViIjoiODMwMmE3NmMtNDA1Mi00NzRiLWE5ODQtYjdlZTVlM2QxNWYxIn0.m_53WoYwuQ54p8Pt5Jc095C2tz1a5_k3z2flRHUJdBQ
-- string
-- eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk1MDMxMjUsImlhdCI6MTczOTQxNjcyNSwic3ViIjoiZGU3YzE2ZmMtZWE1MC00ODNjLWExNjMtNWJmMzkzN2EwMTdlIn0.um2BttkCSfKov_tImgvyyeLXrF-gqAEjCOikNLVJb6w
-- dannamer
-- eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk1MDMxNjIsImlhdCI6MTczOTQxNjc2Miwic3ViIjoiYjdiMjQ1YmQtOTY1ZS00NjZkLTk2NTItNWZjZWNlMDdmMTc3In0.onaBsviK3ThJ6VvB8QM_ex_E6SVVryfbbraDKZrudcw