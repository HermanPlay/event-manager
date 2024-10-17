INSERT INTO public.users (name, email, created_at, updated_at, deleted_at, password, role)
VALUES 
    ('John Doe', 'johndoe@example.com', NOW(), NOW(), NULL, 'password123', 'admin'),
    ('Jane Smith', 'janesmith@example.com', NOW(), NOW(), NULL, 'password123', 'user'),
    ('Mike Johnson', 'mikejohnson@example.com', NOW(), NOW(), NULL, 'password123', 'user'),
    ('Alice Brown', 'alicebrown@example.com', NOW(), NOW(), NULL, 'password123', 'user');

INSERT INTO public.events (title, description, location, date, "time", is_featured, created_by, created_at, updated_at, deleted_at, short_description)
VALUES 
    ('Tech Conference 2024', 'A conference for tech enthusiasts.', 'New York', '2024-11-15', '09:00 AM', TRUE, 1, NOW(), NOW(), NULL, 'Tech event for 2024'),
    ('Music Festival', 'An outdoor music festival.', 'Los Angeles', '2024-12-05', '04:00 PM', TRUE, 2, NOW(), NOW(), NULL, 'Enjoy live music all day'),
    ('Art Expo', 'An exhibition of modern art.', 'San Francisco', '2024-10-25', '11:00 AM', FALSE, 3, NOW(), NOW(), NULL, 'Explore modern art'),
    ('Startup Pitch', 'Pitch your startup ideas to investors.', 'Boston', '2024-11-20', '10:30 AM', FALSE, 1, NOW(), NOW(), NULL, 'Pitch event for startups');

INSERT INTO public.event_users (event_id, user_id, created_at, updated_at, deleted_at)
VALUES
    (1, 2, NOW(), NOW(), NULL),  
    (1, 3, NOW(), NOW(), NULL),  
    (2, 4, NOW(), NOW(), NULL),  
    (3, 2, NOW(), NOW(), NULL),  
    (4, 3, NOW(), NOW(), NULL);
