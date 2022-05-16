--
-- PostgreSQL database dump
--

-- Dumped from database version 13.6 (Ubuntu 13.6-1.pgdg20.04+1)
-- Dumped by pg_dump version 13.6 (Ubuntu 13.6-1.pgdg20.04+1)

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

--
-- Name: seq_app; Type: SEQUENCE; Schema: public; Owner: app_harok
--

CREATE SEQUENCE public.seq_app
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.seq_app OWNER TO app_harok;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: apps; Type: TABLE; Schema: public; Owner: app_harok
--

CREATE TABLE public.apps (
    id integer DEFAULT nextval('public.seq_app'::regclass) NOT NULL,
    name character varying(255) NOT NULL,
    hostname character varying(50) NOT NULL,
    language character varying(15) NOT NULL,
    coderepo character varying(255) NOT NULL,
    imagerepo character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.apps OWNER TO app_harok;

--
-- Name: apps_id_seq; Type: SEQUENCE; Schema: public; Owner: app_harok
--

CREATE SEQUENCE public.apps_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.apps_id_seq OWNER TO app_harok;

--
-- Name: apps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: app_harok
--

ALTER SEQUENCE public.apps_id_seq OWNED BY public.apps.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: app_harok
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO app_harok;

--
-- Name: apps apps_pkey; Type: CONSTRAINT; Schema: public; Owner: app_harok
--

ALTER TABLE ONLY public.apps
    ADD CONSTRAINT apps_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: app_harok
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

