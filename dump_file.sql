--
-- PostgreSQL database dump
--

-- Dumped from database version 15.8
-- Dumped by pg_dump version 16.4

-- Started on 2024-09-11 15:19:27

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
-- TOC entry 215 (class 1259 OID 16400)
-- Name: url_addresses; Type: TABLE; Schema: public; Owner: user123
--

CREATE TABLE public.url_addresses (
    url_id smallint NOT NULL,
    url_full text NOT NULL,
    url_short character varying(100) NOT NULL
);


ALTER TABLE public.url_addresses OWNER TO user123;

--
-- TOC entry 214 (class 1259 OID 16399)
-- Name: url_adresses_url_id_seq; Type: SEQUENCE; Schema: public; Owner: user123
--

CREATE SEQUENCE public.url_adresses_url_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.url_adresses_url_id_seq OWNER TO user123;

--
-- TOC entry 3327 (class 0 OID 0)
-- Dependencies: 214
-- Name: url_adresses_url_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user123
--

ALTER SEQUENCE public.url_adresses_url_id_seq OWNED BY public.url_addresses.url_id;


--
-- TOC entry 3173 (class 2604 OID 16403)
-- Name: url_addresses url_id; Type: DEFAULT; Schema: public; Owner: user123
--

ALTER TABLE ONLY public.url_addresses ALTER COLUMN url_id SET DEFAULT nextval('public.url_adresses_url_id_seq'::regclass);


--
-- TOC entry 3321 (class 0 OID 16400)
-- Dependencies: 215
-- Data for Name: url_addresses; Type: TABLE DATA; Schema: public; Owner: user123
--

COPY public.url_addresses (url_id, url_full, url_short) FROM stdin;
\.


--
-- TOC entry 3328 (class 0 OID 0)
-- Dependencies: 214
-- Name: url_adresses_url_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user123
--

SELECT pg_catalog.setval('public.url_adresses_url_id_seq', 32, true);


--
-- TOC entry 3175 (class 2606 OID 16412)
-- Name: url_addresses uniq_short_url; Type: CONSTRAINT; Schema: public; Owner: user123
--

ALTER TABLE ONLY public.url_addresses
    ADD CONSTRAINT uniq_short_url UNIQUE (url_short);


--
-- TOC entry 3177 (class 2606 OID 16407)
-- Name: url_addresses url_adresses_pkey; Type: CONSTRAINT; Schema: public; Owner: user123
--

ALTER TABLE ONLY public.url_addresses
    ADD CONSTRAINT url_adresses_pkey PRIMARY KEY (url_id);


-- Completed on 2024-09-11 15:19:27

--
-- PostgreSQL database dump complete
--

