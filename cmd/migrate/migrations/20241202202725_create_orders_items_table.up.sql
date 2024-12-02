-- create table for orders, items (food) and connection table
-- many to  many relationship

CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY UNIQUE NOT NULL, -- UUID - postgres type - instead of varchar(16)

    -- orders info such as date, full_price (that enough?)
    full_price NUMERIC(7, 2) NOT NULL,
    currency VARCHAR(10) NOT NULL,

    -- order count
    count SMALLINT NOT NULL,
    
    -- date with time
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create table for items (food)
CREATE TABLE IF NOT EXISTS items (
    id UUID PRIMARY KEY UNIQUE NOT NULL,

    -- item info such as name, price, category, is_available, currency
    name VARCHAR(100) NOT NULL,
    price NUMERIC(7, 2) NOT NULL,
    category VARCHAR(50) NOT NULL,
    
    currency VARCHAR(10) NOT NULL,
    
    count SMALLINT NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT TRUE
);

-- create connection table for many to many relationship
CREATE TABLE IF NOT EXISTS orders_items (
    id UUID PRIMARY KEY UNIQUE NOT NULL,

    order_id UUID NOT NULL REFERENCES orders(id),
    item_id UUID NOT NULL REFERENCES items(id)
);