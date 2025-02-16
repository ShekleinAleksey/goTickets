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
    available_seats INT NOT NULL CHECK (available_seats >= 0)
);

-- Таблица билетов
CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    movie_screening_id INT NOT NULL REFERENCES movie_screenings(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    price INTEGER NOT NULL CHECK (price >= 0),
    seat VARCHAR(50) NOT NULL 
);