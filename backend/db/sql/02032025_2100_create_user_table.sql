CREATE TABLE public.user (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    device_id VARCHAR(255) NOT NULL
);