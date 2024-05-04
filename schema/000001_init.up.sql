CREATE TABLE users
(
   id varchar(255) not null PRIMARY KEY,
   name varchar(255) not null,
   username varchar(255) not null unique,
   password_hash varchar(255) not null
);