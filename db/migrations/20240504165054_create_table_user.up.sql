CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TYPE cat_race AS ENUM (
    'Persian',
    'Maine Coon',
    'Siamese',
    'Ragdoll',
    'Bengal',
    'Sphynx',
    'British Shorthair',
    'Abyssinian',
    'Scottish Fold',
    'Birman'
);

CREATE TYPE cat_sex AS ENUM (
    'male',
    'female'
);

CREATE TABLE cats (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    has_matched BOOLEAN DEFAULT FALSE,
    name VARCHAR(30) NOT NULL,
    race cat_race NOT NULL,  
    sex cat_sex NOT NULL, 
    age_in_month INTEGER, 
    description VARCHAR(200) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE cat_image (
    id SERIAL PRIMARY KEY,
    cat_id INTEGER,
    image_url TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (cat_id) REFERENCES cats(id)
);

CREATE TABLE match_cats (
    id SERIAL PRIMARY KEY,
    match_cat_id INTEGER,
    user_cat_id INTEGER,
    message VARCHAR(120) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (match_cat_id) REFERENCES cats(id),
    FOREIGN KEY (user_cat_id) REFERENCES cats(id)
);