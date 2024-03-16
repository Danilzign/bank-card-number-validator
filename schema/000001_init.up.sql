CREATE TABLE cards
(
    id            uuid DEFAULT gen_random_uuid(),
    number varchar(255) PRIMARY KEY
)