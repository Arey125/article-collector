CREATE TABLE IF NOT EXISTS statuses (
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO statuses (id, name) VALUES ('unread', 'unread');
INSERT INTO statuses (id, name) VALUES ('in_progress', 'in progress');
INSERT INTO statuses (id, name) VALUES ('read', 'read');
    
ALTER TABLE articles ADD COLUMN status_id TEXT NOT NULL REFERENCES statuses(id) DEFAULT 'unread';
