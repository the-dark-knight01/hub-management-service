-- Insert sample data into hubs table
INSERT INTO hubs (name, location) VALUES
                                      ('Hub A', 'New York'),
                                      ('Hub B', 'San Francisco'),
                                      ('Hub C', 'Chicago');

-- Insert sample data into teams table
INSERT INTO teams (name, hub_id) VALUES
                                     ('Team Alpha', 1), -- Associated with Hub A
                                     ('Team Beta', 2),  -- Associated with Hub B
                                     ('Team Gamma', 3); -- Associated with Hub C

-- Insert sample data into users table
INSERT INTO users (name, email, team_id) VALUES
                                             ('John Doe', 'john.doe@example.com', 1), -- Associated with Team Alpha
                                             ('Jane Smith', 'jane.smith@example.com', 2), -- Associated with Team Beta
                                             ('Mike Johnson', 'mike.johnson@example.com', 3); -- Associated with Team Gamma