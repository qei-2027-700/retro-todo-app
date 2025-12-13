CREATE TABLE IF NOT EXISTS sprints (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(50) DEFAULT 'bg-purple-500',
    is_favorite BOOLEAN DEFAULT false,
    is_deleted BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_sprints_is_favorite ON sprints(is_favorite);

CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT false,
    sprint_id INTEGER REFERENCES sprints(id) ON DELETE SET NULL,
    is_deleted BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_todos_completed ON todos(completed);
CREATE INDEX IF NOT EXISTS idx_todos_sprint_id ON todos(sprint_id);

-- サンプルデータ
INSERT INTO sprints (name, color, is_favorite) VALUES
('バックログ', 'bg-blue-500', false),
('2510-3', 'bg-purple-500', false),
('2510-4', 'bg-purple-500', false),
('Personal Sprint', 'bg-purple-500', true)
ON CONFLICT DO NOTHING;
