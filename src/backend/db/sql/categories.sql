CREATE TABLE IF NOT EXISTS public.categories (
  id bigserial not null,
  name varchar(45) default null,
  description text default null,
  status varchar(1) default 'O',
  flag varchar(4) default 'dev',
  active boolean default true,
  uuid char(36) default uuid_generate_v4(),
  created_at timestamptz default now(),
  updated_at timestamptz default now(),
  deleted_at timestamptz null,
  CONSTRAINT categories_pkey PRIMARY KEY (id)
);

ALTER SEQUENCE categories_id_seq MINVALUE 1000 START 1000 RESTART WITH 1000;
CREATE UNIQUE INDEX category_name_unique_idx ON categories (name);

