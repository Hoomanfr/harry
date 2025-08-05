-- +goose Up
create table if not exists dbo.fulfillment(
    id serial primary key,
    order_id int not null,
    sku int not null,
    quantity int not null,
    status varchar(50) not null,
    fulfilled_at timestamp default current_timestamp
);

-- +goose Down
drop table if exists dbo.fulfillment;