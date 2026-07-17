create table clients (
    id uuid primary key default gen_random_uuid(),
    name varchar(100),
    tls_pem Text not null,
    valid_to timestamp with time zone default now() + interval '1 day'
);

create table servers(
    id uuid primary key default gen_random_uuid(),
    name varchar(100),
    address varchar(50) -- ip:port or domain:port
);

create table permissions (
    id serial primary key,
    client_id uuid references clients(id) on delete cascade,
    server_id uuid references servers(id) on delete cascade,
    permission integer default 0,
    unique (client_id, server_id)
);
