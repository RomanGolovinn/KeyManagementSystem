create table client (
    id uuid primary key default gen_random_uuid(),
    name varchar(100),
    tls_pem Text not null,
    valid_to timestamp with time zone default now() + interval '1 day'
);

create table server (
    id uuid primary key default gen_random_uuid(),
    name varchar(100),
    address varchar(50) -- ip:port or domain:port
);

create table permissions (
    id serial primary key,
    client_id integer references client(id) on delete cascade,
    server_id integer references server(id) on delete cascade,
    unique (client_id, server_id)
)
