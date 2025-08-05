Harry is a personal project exploring a microservices architecture to process returned items in a warehouse. It uses `NATS` for `messaging` and `Centrifugo` with `WebSockets` to notify fulfillment clients in real time, which helps unblock fulfillment items.

checkout `Makefile` at the root and see how to run/stop all infra and services. as well as run/stop specific.

use `docker exec -it <centrifugo_container>`

run `centrifugo gentoken -u 123722`

copy the generated token and use it for the `CENTRIFUGO_TOKEN` in related `.env` files and postman.

in postman harry websocket after `connect` don't forge to run `send` after to listen to topics.

in fulfillment service:

for testing the following creates some item in fulfillment db to find some item to unblock:

1. `make run-all`

2.

```
INSERT INTO dbo.fulfillment
(id, order_id, sku, quantity, status, fulfilled_at)
VALUES(nextval('dbo.fulfillment_id_seq'::regclass), 10, 'sku-1234', 2, 'blocked', CURRENT_TIMESTAMP);

SELECT id, order_id, sku, quantity, status, fulfilled_at
FROM dbo.fulfillment;
```

3. use postman to connect to `ws://localhost:8000/connection/websocket?cf_ws_frame_ping_pong=true`
message:

```
{"id": 1, "connect": { "token": "<the_generated_token>"}}
{"id": 2, "subscribe": {"channel": "fulfillment.order_unblocked"}}
```

4.

```
curl --location 'http://localhost:8080/return-item' \
--header 'Content-Type: application/json' \
--data '{
    "sku": "sku-5234",
    "quantity": 2
}'
```
