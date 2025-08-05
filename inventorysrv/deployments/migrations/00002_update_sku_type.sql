-- +goose Up
alter table dbo.inventory
    alter column sku type varchar(50) using sku::varchar(50);

-- +goose Down
alter table dbo.inventory
    alter column sku type int using sku::int;