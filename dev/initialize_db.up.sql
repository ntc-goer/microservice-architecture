-------------------------------------- Order -----------------------------
-- Create the database
CREATE DATABASE orderdb;
-- Create the user with the specified password
CREATE USER orderuser WITH PASSWORD 'orderpwd';
-- Grant all privileges on the database to the user
ALTER USER orderuser WITH SUPERUSER