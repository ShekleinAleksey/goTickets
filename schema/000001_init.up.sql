--Таблица фильмов
CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Таблица пользователей 
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    balance DECIMAL(10, 2)
);

-- Таблица сеансов
CREATE TABLE IF NOT EXISTS movie_screenings (
    id SERIAL PRIMARY KEY,
    movie_id INT NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    show_time TIMESTAMP NOT NULL,
    hall INT NOT NULL CHECK (hall > 0)
);

-- Таблица билетов
CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    movie_screening_id INT NOT NULL REFERENCES movie_screenings(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    price DECIMAL(10, 2) NOT NULL CHECK (price >= 0),
    seat VARCHAR(50) NOT NULL 
);