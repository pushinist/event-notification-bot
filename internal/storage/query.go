package storage

const (
	createEventsTable = `
		CREATE TABLE IF NOT EXISTS events(
		message_id BIGINT PRIMARY KEY,
		date TIMESTAMP NOT NULL
		);
	`

	createParticipantsTable = `
		CREATE TABLE IF NOT EXISTS participants(
		user_id BIGINT,
		event_message_id BIGINT,
		NOTIFIED BOOLEAN NOT NULL DEFAULT FALSE,
		PRIMARY KEY (user_id, event_message_id),
		UNIQUE (user_id, event_message_id),

		FOREIGN KEY (event_message_id) REFERENCES events(message_id)
		);
	`

	createAdminsTable = `
		CREATE TABLE IF NOT EXISTS admins(
		user_id BIGINT PRIMARY KEY
		);
	`

	createAdminEventTable = `
		CREATE TABLE IF NOT EXISTS admin_event(
		admin_id BIGINT,
		event_message_id BIGINT,
		NOTIFIED BOOLEAN NOT NULL DEFAULT FALSE,
		PRIMARY KEY (admin_id, event_message_id),
		UNIQUE (admin_id, event_message_id),

		FOREIGN KEY (admin_id) REFERENCES admins(user_id),
		FOREIGN KEY (event_message_id) REFERENCES events(message_id)
		)
		`
)
