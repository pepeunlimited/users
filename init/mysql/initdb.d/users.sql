CREATE DATABASE IF NOT EXISTS users CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE users;

CREATE TABLE users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    username VARCHAR (320) UNIQUE NOT NULL,
    email VARCHAR(320) UNIQUE NOT NULL,
    password CHAR(72) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    is_banned BOOLEAN NOT NULL DEFAULT FALSE,
    is_locked BOOLEAN NOT NULL DEFAULT FALSE,
    last_modified DATETIME(3) NOT NULL,
    profile_picture_id BIGINT NULL, -- reference from the files-service
    PRIMARY KEY (id)
);

CREATE TABLE tickets (
    id            BIGINT      NOT NULL AUTO_INCREMENT,
    token         CHAR(72)    UNIQUE NOT NULL,
    created_at    DATETIME(3) NOT NULL,
    expires_at    DATETIME(3) NOT NULL,
    user_tickets  BIGINT      NULL,
    FOREIGN KEY (user_tickets) REFERENCES users (id),
    PRIMARY KEY (id)
);

CREATE TABLE roles (
    id            BIGINT       NOT NULL AUTO_INCREMENT,
    role          VARCHAR(200) NOT NULL DEFAULT 'user', -- user, admin, reviewer
    user_roles    BIGINT NULL,
    FOREIGN KEY (user_roles) REFERENCES users (id),
    PRIMARY KEY (id)
);