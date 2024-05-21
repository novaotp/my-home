PRAGMA foreign_keys = TRUE;

CREATE TABLE food (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    quantity REAL NOT NULL DEFAULT 0,
    unit VARCHAR(255) NOT NULL
);

CREATE TABLE shopping_list (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE shopping_list_food (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    shopping_list_id INTEGER NOT NULL,
    food_id INTEGER NOT NULL,
    priority INTEGER NOT NULL DEFAULT 3,
    purchased BOOLEAN NOT NULL DEFAULT FALSE,
    quantity_to_buy REAL,
    notes TEXT,
    FOREIGN KEY (food_id) REFERENCES food (id) ON DELETE CASCADE,
    FOREIGN KEY (shopping_list_id) REFERENCES shopping_list (id) ON DELETE CASCADE
);
