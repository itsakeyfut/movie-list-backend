SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;


CREATE TABLE public.genres (
    id integer NOT NULL,
    genre character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.genres ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.genres_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE public.movies (
    id integer NOT NULL,
    title VARCHAR(512) NOT NULL,
    description TEXT,
    image character varying(255),
    release_date DATE,
    rating DOUBLE PRECISION,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.movies_genres (
    id integer NOT NULL,
    movie_id integer,
    genre_id integer
);

ALTER TABLE public.movies_genres ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.movies_genres_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

ALTER TABLE public.movies ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.movies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE public.users (
    id integer NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255),
    password character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


SELECT pg_catalog.setval('public.genres_id_seq', 13, true);
SELECT pg_catalog.setval('public.movies_genres_id_seq', 6, true);
SELECT pg_catalog.setval('public.movies_id_seq', 3, true);
SELECT pg_catalog.setval('public.users_id_seq', 1, true);

ALTER TABLE ONLY public.genres
    ADD CONSTRAINT genres_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT movies_genres_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT movies_genres_genre_id_fkey FOREIGN KEY (genre_id) REFERENCES public.genres(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT movies_genres_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id) ON UPDATE CASCADE ON DELETE CASCADE;

CREATE INDEX idx_movies_genres_movie_id ON public.movies_genres(movie_id);
CREATE INDEX idx_movies_genres_genre_id ON public.movies_genres(genre_id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_unique UNIQUE (email);
CREATE INDEX idx_users_email ON public.users(email);