-- +goose Up
ALTER TABLE users
ADD CONSTRAINT users_name_unique UNIQUE (name);