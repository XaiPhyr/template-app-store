CREATE TABLE IF NOT EXISTS public.products (
  id bigserial not null,
  name varchar(45) default null,
  description text default null,
  price numeric(11, 2) default 0.00,
  category varchar(45) default null,
  status varchar(1) default 'O',
  flag varchar(4) default 'dev',
  active boolean default true,
  uuid char(36) default uuid_generate_v4(),
  created_at timestamptz default now(),
  updated_at timestamptz default now(),
  deleted_at timestamptz null,
  CONSTRAINT products_pkey PRIMARY KEY (id),
  CONSTRAINT fk_categories FOREIGN KEY (category) REFERENCES public.categories (name) 
    ON DELETE SET NULL 
    ON UPDATE CASCADE
);

ALTER SEQUENCE products_id_seq MINVALUE 1000 START 1000 RESTART WITH 1000;
CREATE UNIQUE INDEX product_name_unique_idx ON products (name);