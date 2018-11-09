SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: users; Type: TABLE; Schema: public; Owner: drd
--

CREATE TABLE public.users (
    id character varying NOT NULL,
    tokens double precision DEFAULT 0.0 NOT NULL
);


ALTER TABLE public.users OWNER TO drd;

--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: drd
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--
