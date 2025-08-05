-- +goose Up
ALTER TABLE dbo.inventory ADD CONSTRAINT unique_sku UNIQUE (sku);

-- +goose Down
ALTER TABLE dbo.inventory DROP CONSTRAINT unique_sku;