CREATE ROLE admin_casts_service WITH
    LOGIN
    ENCRYPTED PASSWORD 'SCRAM-SHA-256$4096:R9TMUdvkUG5yxu0rJlO+hA==$E/WRNMfl6SWK9xreXN8rfIkJjpQhWO8pd+8t2kx12D0=:sCS47DCNVIZYhoue/BReTE0ZhVRXzMGszsnnHexVwOU=';

CREATE TABLE professions (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE casts (
    movie_id INT NOT NULL,
    actor_id INT NOT NULL,
    profession_id INT REFERENCES professions(id) ON DELETE SET NULL ON UPDATE CASCADE,
    PRIMARY KEY(movie_id,actor_id,profession_id)
);

CREATE TABLE casts_labels (
    movie_id INT PRIMARY KEY,
    label TEXT NOT NULL 
);


CREATE OR REPLACE FUNCTION remove_cast()
RETURNS TRIGGER
AS $$
BEGIN
    IF NOT EXISTS (SELECT * FROM casts WHERE movie_id = OLD.movie_id) THEN
        DELETE FROM casts_labels WHERE movie_id = OLD.movie_id;
    END IF;
    RETURN NEW;
END; $$
LANGUAGE PLPGSQL;


CREATE TRIGGER remove_cast_trigger
    AFTER DELETE ON casts
    FOR EACH ROW
    EXECUTE PROCEDURE remove_cast();


GRANT SELECT, UPDATE, INSERT, DELETE ON casts TO admin_casts_service;
GRANT SELECT, UPDATE, INSERT, DELETE ON casts_labels TO admin_casts_service;
GRANT USAGE, SELECT ON SEQUENCE  professions_id_seq TO admin_casts_service;
GRANT SELECT, UPDATE, INSERT, DELETE ON professions TO admin_casts_service;

