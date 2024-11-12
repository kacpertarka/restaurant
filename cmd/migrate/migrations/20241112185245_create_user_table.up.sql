-- create first version of user table (for employees and boss?)


-- table for emplyess and boss? - store names
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,

    -- user_id - string with user login identifier, required to login with password
    user_id VARCHAR(255) UNIQUE NOT NULL,

    -- info about user
    first_name VARCHAR(128) NOT NULL,
    last_name VARCHAR(128) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(128),

    password BYTEA NOT NULL,
    
    -- simple adres - only city
    address VARCHAR(255),

    -- roles like: worker, admin 
    role VARCHAR(10) NOT NULL, 

    -- info about when emplyee was hired
    created_at DATE NOT NULL,

    -- info about account activation
    is_active BOOLEAN DEFAULT FALSE
)