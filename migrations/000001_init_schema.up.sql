create type permission_type as enum (
    'none',
    'payment:initiate', 
    'payment:approve', 
    'payment:execute', 
    'payment:reverse',
    'account:read:balance', 
    'account:read:history', 
    'account:read:pii', 
    'config:limits:write', 
    'product:manage', 
    'system:override'
);

create table clients (
    id uuid primary key default gen_random_uuid(),
    name varchar(100) unique not null,
    tls_pem Text not null,
    valid_to timestamp with time zone default now() + interval '1 day'
);

create table servers(
    id uuid primary key default gen_random_uuid(),
    name varchar(100) unique not null,
    address varchar(50) -- ip:port or domain:port
);

create table permissions (
    id serial primary key,
    client_id uuid references clients(id) on delete cascade,
    server_id uuid references servers(id) on delete cascade,
    permission permission_type default 'none' not null,
    unique (client_id, server_id)
);

create table issued_credentials (
    id uuid primary key default gen_random_uuid(),
    client_id uuid references clients(id) on delete cascade,
    server_id uuid references servers(id) on delete cascade,
    permission permission_type not null,
    temp_username varchar(100) not null,
    temp_password_hash varchar(255) not null,
    created_at timestamp with time zone default now(),
    valid_to timestamp with time zone not null,
    revoked_at timestamp with time zone
);

create index idx_issued_credentials_valid_to on issued_credentials(valid_to);

create table audit_log (
    id uuid primary key default gen_random_uuid(),
    timestamp timestamp with time zone default now() not null,
    client_id uuid references clients(id) on delete set null,
    server_id uuid references servers(id) on delete set null,
    event_type varchar(100) not null,
    ip_address varchar(50) not null,
    details jsonb
);