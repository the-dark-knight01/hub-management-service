-- Up: Create hubs table
CREATE TABLE hubs (
                      id SERIAL PRIMARY KEY,
                      name VARCHAR(255) NOT NULL,
                      location VARCHAR(255),
                      created_at TIMESTAMP DEFAULT NOW(),
                      updated_at TIMESTAMP DEFAULT NOW()
);

-- Up: Create teams table
CREATE TABLE teams (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       hub_id INT NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW(),
                       FOREIGN KEY (hub_id) REFERENCES hubs (id) ON DELETE CASCADE
);

-- Up: Create users table
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       team_id INT NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW(),
                       FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE CASCADE
);