DROP TABLE IF EXISTS mails;
CREATE TABLE mails
(
    id              BIGINT PRIMARY KEY NOT NULL,
    sender_email    VARCHAR            NOT NULL,
    recipient_email VARCHAR            NOT NULL,
    title           VARCHAR,
    body            VARCHAR,
    is_bookmark     BOOLEAN            NOT NULL,
    is_read         BOOLEAN            NOT NULL,
    time            BIGINT             NOT NULL
);