CREATE TABLE IF NOT EXISTS participants (
    user_id BIGINT,
    event_message_id BIGINT,
    NOTIFIED BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (user_id, event_message_id),
    UNIQUE (user_id, event_message_id),
    FOREIGN KEY (event_message_id) REFERENCES events (message_id)
);
