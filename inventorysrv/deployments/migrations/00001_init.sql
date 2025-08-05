-- +goose Up
create table if not exists dbo.inventory(
    id serial primary key,
    sku int not null,
    quantity int not null,
    updated_at timestamp default current_timestamp
);

-- +goose Down
drop table if exists dbo.inventory;
