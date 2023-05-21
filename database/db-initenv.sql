--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

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
-- Name: diffs; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.diffs (
    id bigint NOT NULL,
    gtm_env character varying(10),
    gtm_table character varying(20),
    key character varying(50),
    value text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at time with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.diffs OWNER TO admin;

--
-- Name: diffs_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.diffs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.diffs_id_seq OWNER TO admin;

--
-- Name: diffs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.diffs_id_seq OWNED BY public.diffs.id;


--
-- Name: masters; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.masters (
    id bigint NOT NULL,
    key character varying(50),
    value text,
    gtm_table character varying(20),
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.masters OWNER TO admin;

--
-- Name: sames_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.sames_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sames_id_seq OWNER TO admin;

--
-- Name: sames_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.sames_id_seq OWNED BY public.masters.id;


--
-- Name: diffs id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.diffs ALTER COLUMN id SET DEFAULT nextval('public.diffs_id_seq'::regclass);


--
-- Name: masters id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.masters ALTER COLUMN id SET DEFAULT nextval('public.sames_id_seq'::regclass);


--
-- Data for Name: diffs; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.diffs (id, gtm_env, gtm_table, key, value, created_at, updated_at, deleted_at) FROM stdin;
1	iuat	utbl	aaa-bbb	^UTBL("aaa","bbb")=1|2|3|4	2023-05-21 07:28:27.620886+07	07:28:27.620886+07	\N
\.


--
-- Data for Name: masters; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.masters (id, key, value, gtm_table, created_at, updated_at, deleted_at) FROM stdin;
1	aaa-bbb	^UTBL("aaa","bbb")=1|2|3|4	utbl	2023-05-21 07:25:50.853921+07	2023-05-21 07:25:50.853921+07	\N
2	ZTLRPMTH-AN0816-6	^UTBL("ZTLRPMTH","AN0816",6)=61005|86002|BATCH|[ZUTBLTLRPMT]COMPTITLE1|BATCH|61005	UTBL	2023-05-21 08:18:19.311935+07	2023-05-21 08:18:19.311935+07	\N
3	ZTLRPMTH-AN0816-7	^UTBL("ZTLRPMTH","AN0816",7)=61005|86002|BATCH|[ZUTBLTLRPMT]COMPTITLE1|BATCH|61005	UTBL	2023-05-21 08:26:49.062473+07	2023-05-21 08:26:49.062473+07	\N
\.


--
-- Name: diffs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.diffs_id_seq', 1, true);


--
-- Name: sames_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.sames_id_seq', 3, true);


--
-- Name: diffs diffs_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.diffs
    ADD CONSTRAINT diffs_pkey PRIMARY KEY (id);


--
-- Name: masters sames_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.masters
    ADD CONSTRAINT sames_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

