--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

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
-- Name: reservations; Type: TABLE; Schema: public; Owner: jinyanomura
--

CREATE TABLE public.reservations (
    id integer NOT NULL,
    first_name character varying(255) DEFAULT ''::character varying NOT NULL,
    last_name character varying(255) DEFAULT ''::character varying NOT NULL,
    email character varying(255) NOT NULL,
    phone character varying(255),
    table_id integer NOT NULL,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.reservations OWNER TO jinyanomura;

--
-- Name: reservations_id_seq; Type: SEQUENCE; Schema: public; Owner: jinyanomura
--

CREATE SEQUENCE public.reservations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reservations_id_seq OWNER TO jinyanomura;

--
-- Name: reservations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jinyanomura
--

ALTER SEQUENCE public.reservations_id_seq OWNED BY public.reservations.id;


--
-- Name: restriction; Type: TABLE; Schema: public; Owner: jinyanomura
--

CREATE TABLE public.restriction (
    id integer NOT NULL,
    restriction_type character varying(50) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.restriction OWNER TO jinyanomura;

--
-- Name: restriction_id_seq; Type: SEQUENCE; Schema: public; Owner: jinyanomura
--

CREATE SEQUENCE public.restriction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.restriction_id_seq OWNER TO jinyanomura;

--
-- Name: restriction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jinyanomura
--

ALTER SEQUENCE public.restriction_id_seq OWNED BY public.restriction.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: jinyanomura
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO jinyanomura;

--
-- Name: table_restrictions; Type: TABLE; Schema: public; Owner: jinyanomura
--

CREATE TABLE public.table_restrictions (
    id integer NOT NULL,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL,
    table_id integer NOT NULL,
    restriction_id integer NOT NULL,
    reservation_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.table_restrictions OWNER TO jinyanomura;

--
-- Name: table_restrictions_id_seq; Type: SEQUENCE; Schema: public; Owner: jinyanomura
--

CREATE SEQUENCE public.table_restrictions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.table_restrictions_id_seq OWNER TO jinyanomura;

--
-- Name: table_restrictions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jinyanomura
--

ALTER SEQUENCE public.table_restrictions_id_seq OWNED BY public.table_restrictions.id;


--
-- Name: tables; Type: TABLE; Schema: public; Owner: jinyanomura
--

CREATE TABLE public.tables (
    id integer NOT NULL,
    capacity integer DEFAULT 1 NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tables OWNER TO jinyanomura;

--
-- Name: tables_id_seq; Type: SEQUENCE; Schema: public; Owner: jinyanomura
--

CREATE SEQUENCE public.tables_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tables_id_seq OWNER TO jinyanomura;

--
-- Name: tables_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jinyanomura
--

ALTER SEQUENCE public.tables_id_seq OWNED BY public.tables.id;


--
-- Name: reservations id; Type: DEFAULT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.reservations ALTER COLUMN id SET DEFAULT nextval('public.reservations_id_seq'::regclass);


--
-- Name: restriction id; Type: DEFAULT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.restriction ALTER COLUMN id SET DEFAULT nextval('public.restriction_id_seq'::regclass);


--
-- Name: table_restrictions id; Type: DEFAULT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.table_restrictions ALTER COLUMN id SET DEFAULT nextval('public.table_restrictions_id_seq'::regclass);


--
-- Name: tables id; Type: DEFAULT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.tables ALTER COLUMN id SET DEFAULT nextval('public.tables_id_seq'::regclass);


--
-- Name: reservations reservations_pkey; Type: CONSTRAINT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_pkey PRIMARY KEY (id);


--
-- Name: restriction restriction_pkey; Type: CONSTRAINT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.restriction
    ADD CONSTRAINT restriction_pkey PRIMARY KEY (id);


--
-- Name: table_restrictions table_restrictions_pkey; Type: CONSTRAINT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.table_restrictions
    ADD CONSTRAINT table_restrictions_pkey PRIMARY KEY (id);


--
-- Name: tables tables_pkey; Type: CONSTRAINT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: jinyanomura
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: reservations reservations_table_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_table_id_fkey FOREIGN KEY (table_id) REFERENCES public.tables(id);


--
-- Name: table_restrictions table_restrictions_reservation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.table_restrictions
    ADD CONSTRAINT table_restrictions_reservation_id_fkey FOREIGN KEY (reservation_id) REFERENCES public.reservations(id) ON DELETE CASCADE;


--
-- Name: table_restrictions table_restrictions_restriction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.table_restrictions
    ADD CONSTRAINT table_restrictions_restriction_id_fkey FOREIGN KEY (restriction_id) REFERENCES public.restriction(id);


--
-- Name: table_restrictions table_restrictions_table_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: jinyanomura
--

ALTER TABLE ONLY public.table_restrictions
    ADD CONSTRAINT table_restrictions_table_id_fkey FOREIGN KEY (table_id) REFERENCES public.tables(id);


--
-- PostgreSQL database dump complete
--

