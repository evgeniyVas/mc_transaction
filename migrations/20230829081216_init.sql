-- +goose Up
-- +goose StatementBegin

-- users
CREATE TABLE IF NOT EXISTS users
(
    id   BIGINT       NOT NULL UNIQUE,
    name varchar(100) NULL,

    primary key (id)
);
ALTER TABLE
    users
    OWNER TO user123;

-- transactions
CREATE TABLE IF NOT EXISTS transactions
(
    id         BIGSERIAL                NOT NULL UNIQUE,
    user_id    BIGINT                   NOT NULL,
    amount     DECIMAL(20, 2)           NOT NULL,
    status     TEXT                     NULL,
    pay_id     INT                      NULL,
    locked     BOOLEAN                  NULL,
    token      TEXT                     NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    primary key (id),
    foreign key (user_id) references users (id)
);
ALTER TABLE
    transactions
    OWNER TO user123;

-- balance
CREATE TABLE IF NOT EXISTS balance
(
    id      BIGINT         NOT NULL UNIQUE,
    user_id BIGINT         NOT NULL,
    amount  DECIMAL(20, 2) NOT NULL,

    primary key (id),
    foreign key (user_id) references users (id)
);
ALTER TABLE
    balance
    OWNER TO user123;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS balance;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
