CREATE TABLE IF NOT EXISTS admin_event (
    admin_id BIGINT,
    event_message_id BIGINT,
    NOTIFIED BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (admin_id, event_message_id),
    UNIQUE (admin_id, event_message_id),
    FOREIGN KEY (admin_id) REFERENCES admins (user_id),
    FOREIGN KEY (event_message_id) REFERENCES events (message_id)
);
