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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: apps; Type: TABLE; Schema: public; Owner: app_harok
--

CREATE TABLE public.apps (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    hostname character varying(50) NOT NULL,
    language character varying(15) NOT NULL,
    code_repo character varying(255) NOT NULL,
    image_repo character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.apps OWNER TO app_harok;

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

