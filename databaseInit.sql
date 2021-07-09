CREATE DATABASE IF NOT EXISTS tekabTest;
use tekabTest;
CREATE TABLE users (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    email varchar(255),
    password varchar(255)
);
INSERT INTO users (email, password) values (
    "admin@admin.com","$2a$10$bmhOR2r3iYBknefZ.Qto2efB.or6w/YTHgkv1IrUL7e663/IymEVu"
);

