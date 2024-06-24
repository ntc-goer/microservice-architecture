-------------------------------------- Order -----------------------------
-- Create the user with the specified password
CREATE USER orderuser WITH PASSWORD 'orderpwd';
-- Grant all privileges on the database to the user
ALTER USER orderuser WITH SUPERUSER;

-------------------------------------- Kitchen -----------------------------
-- Create the user with the specified password
CREATE USER kitchenuser WITH PASSWORD 'kitchenpwd';
-- Grant all privileges on the database to the user
ALTER USER kitchenuser WITH SUPERUSER;