--Таблица фильмов
CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    duration INTEGER NOT NULL CHECK (duration > 0),
    genre VARCHAR(255) NOT NULL
);

-- Таблица пользователей 
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    balance INTEGER NOT NULL DEFAULT 0
);  

-- Таблица сеансов
CREATE TABLE IF NOT EXISTS movie_screenings (
    id SERIAL PRIMARY KEY,
    movie_id INT NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    show_time TIMESTAMP NOT NULL,
    hall INT NOT NULL CHECK (hall > 0),
    available_seats INT NOT NULL CHECK (available_seats >= 0),
    ticket_price INTEGER NOT NULL CHECK (ticket_price >= 0)
);


-- Таблица билетов
CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    movie_screening_id INT NOT NULL REFERENCES movie_screenings(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    price INTEGER NOT NULL CHECK (price >= 0),
    seat VARCHAR(50) NOT NULL 
);

INSERT INTO movies (name, description, duration, genre) VALUES
('Inception', 'A mind-bending sci-fi thriller', 148, 'Sci-Fi'),
('The Dark Knight', 'A gripping crime drama', 152, 'Action'),
('Pulp Fiction', 'A neo-noir crime thriller', 154, 'Crime'),
('The Matrix', 'A groundbreaking sci-fi action', 136, 'Sci-Fi'),
('The Lord of the Rings', 'An epic fantasy adventure', 228, 'Fantasy');

INSERT INTO movie_screenings (movie_id, show_time, hall, available_seats, ticket_price) VALUES
(1, '2025-01-01 10:00:00', 1, 10, 350),
(2, '2025-01-01 12:00:00', 2, 10, 200),
(3, '2025-01-01 14:00:00', 3, 10, 300),
(4, '2025-01-01 16:00:00', 4, 10, 400);


