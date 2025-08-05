-- +goose Up
create table if not exists dbo.receiving(
    id serial primary key,
    sku int not null,
    quantity int not null,
    received_at timestamp default current_timestamp
);

-- +goose Down
drop table if exists dbo.receiving;
