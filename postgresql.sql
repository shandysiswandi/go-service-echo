-- CREATED_AT = 17-01-2021 12:12:01
-- create table `users` with id using uuid
DROP TABLE IF EXISTS public.users;
CREATE TABLE public.users
(
    id char(36) not null constraint users_pkey primary key,
    name varchar(100) not null,
    email varchar(100) not null constraint users_email_key unique,
    password varchar(60)  not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp
);

-- add one data to login // password=password
INSERT INTO users(id, name, email, password)
VALUES ('190e9c30-c119-40a2-a554-85b6fb8cffaf', 'admin', 'admin@admin.com',
        '$2y$10$M.VrmGnEdrJyYyYTGmkynexSg9MjHDhPckyyVJvxoZGRAOMKDL3fq ');
