--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1 (Debian 14.1-1.pgdg110+1)
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
-- Name: accounts; Type: TABLE; Schema: public; Owner: gouser
--

CREATE TABLE public.accounts (
    id integer NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    phone_number character varying(16) NOT NULL,
    gender bit(1) NOT NULL,
    birth_date date NOT NULL,
    is_verified boolean DEFAULT false,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    tag integer DEFAULT 1
);


ALTER TABLE public.accounts OWNER TO gouser;

--
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: gouser
--

CREATE SEQUENCE public.accounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts_id_seq OWNER TO gouser;

--
-- Name: accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gouser
--

ALTER SEQUENCE public.accounts_id_seq OWNED BY public.accounts.id;


--
-- Name: admins; Type: TABLE; Schema: public; Owner: gouser
--

CREATE TABLE public.admins (
    id integer NOT NULL,
    username text NOT NULL,
    password text NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.admins OWNER TO gouser;

--
-- Name: admins_id_seq; Type: SEQUENCE; Schema: public; Owner: gouser
--

CREATE SEQUENCE public.admins_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.admins_id_seq OWNER TO gouser;

--
-- Name: admins_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gouser
--

ALTER SEQUENCE public.admins_id_seq OWNED BY public.admins.id;


--
-- Name: khotbah; Type: TABLE; Schema: public; Owner: gouser
--

CREATE TABLE public.khotbah (
    id integer NOT NULL,
    thumbnail character varying(256),
    title character varying(256) NOT NULL,
    link character varying(256) NOT NULL,
    pendeta_name character varying(256) NOT NULL,
    ibadah_date timestamp without time zone NOT NULL,
    link_warta character varying(256) NOT NULL
);


ALTER TABLE public.khotbah OWNER TO gouser;

--
-- Name: khotbah_id_seq; Type: SEQUENCE; Schema: public; Owner: gouser
--

CREATE SEQUENCE public.khotbah_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.khotbah_id_seq OWNER TO gouser;

--
-- Name: khotbah_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gouser
--

ALTER SEQUENCE public.khotbah_id_seq OWNED BY public.khotbah.id;


--
-- Name: profiles; Type: TABLE; Schema: public; Owner: gouser
--

CREATE TABLE public.profiles (
    profile_id integer NOT NULL,
    blood_type character varying(2),
    job character varying(256),
    age integer,
    address text,
    account_id integer
);


ALTER TABLE public.profiles OWNER TO gouser;

--
-- Name: profiles_profile_id_seq; Type: SEQUENCE; Schema: public; Owner: gouser
--

CREATE SEQUENCE public.profiles_profile_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.profiles_profile_id_seq OWNER TO gouser;

--
-- Name: profiles_profile_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gouser
--

ALTER SEQUENCE public.profiles_profile_id_seq OWNED BY public.profiles.profile_id;


--
-- Name: accounts id; Type: DEFAULT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.accounts ALTER COLUMN id SET DEFAULT nextval('public.accounts_id_seq'::regclass);


--
-- Name: admins id; Type: DEFAULT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.admins ALTER COLUMN id SET DEFAULT nextval('public.admins_id_seq'::regclass);


--
-- Name: khotbah id; Type: DEFAULT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.khotbah ALTER COLUMN id SET DEFAULT nextval('public.khotbah_id_seq'::regclass);


--
-- Name: profiles profile_id; Type: DEFAULT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.profiles ALTER COLUMN profile_id SET DEFAULT nextval('public.profiles_profile_id_seq'::regclass);


--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: gouser
--

-- COPY public.accounts (id, first_name, last_name, email, password, phone_number, gender, birth_date, is_verified, created_at, updated_at, tag) FROM stdin;
-- 3	benaya123	rio	benaya.terong@gmail.com	$2a$10$ZwAHaWyHyx2JH20VOqUDFO2OemFY9OYrBAehbh5KW7oru8u8K7ocG	085814060717	0	2000-07-24	t	2022-01-23 15:45:30.699306	2022-01-24 13:16:27.296916	1
-- 8	benaya	rio	benaya.terong1@gmail.com	$2a$10$lgDyyxmAFlanOJyccndziOiIFXe719.B45kLg6VjtDFuZ4QzMpYMO	085814060717	0	2000-07-24	f	2022-01-25 13:40:27.552121	2022-01-25 13:40:27.552121	1
-- 1	benaya edit	rio	benaya.c@mail.ugm.ac.id	$2a$10$hsRZ1OxvsWQmaXTXvfAuYucbB9ToK8KsuaETbkcfM3rw.LzVugdhq	085814060717	0	2000-07-24	t	2022-01-23 15:35:28.079192	2022-01-27 14:52:01.661392	1
-- \.


--
-- Data for Name: admins; Type: TABLE DATA; Schema: public; Owner: gouser
--

COPY public.admins (id, username, password, created_at, updated_at) FROM stdin;
1	admin_test_dev	$2a$10$.57D7YXmKFCgp37LmQagy.XS0EQKeFhIl5IxiNJcRg4lmgjlX7ogq	2022-01-24 14:02:28.342295	2022-01-24 14:02:28.342295
\.


--
-- Data for Name: khotbah; Type: TABLE DATA; Schema: public; Owner: gouser
--

-- COPY public.khotbah (id, thumbnail, title, link, pendeta_name, ibadah_date, link_warta) FROM stdin;
-- 1	url thumbnail	judul khotbah	url video khotbah	pdt aldrich dheja saleki	2022-01-30 21:00:00	url warta
-- 2	url thumbnail2	judul khotbah2	url video khotbah2	pdt aldrich dheja saleki	2022-02-02 21:00:00	url warta2
-- 3	url thumbnail2	judul khotbah2	url video khotbah2	pdt aldrich dheja saleki	2022-02-02 21:00:00	url warta2
-- 4	url thumbnail3	judul khotbah kemarin	url video khotbah3	pdt aldrich dheja saleki	2022-01-26 09:00:00	url warta2
-- 5	url thumbnail3	judul khotbah kemarin	url video khotbah3	pdt aldrich dheja saleki	2022-01-25 09:00:00	url warta2
-- 6	url thumbnail3	judul khotbah kemarin baru banget	url video khotbah3	pdt aldrich dheja saleki	2022-01-26 10:00:00	url warta2
-- \.


--
-- Data for Name: profiles; Type: TABLE DATA; Schema: public; Owner: gouser
--

-- COPY public.profiles (profile_id, blood_type, job, age, address, account_id) FROM stdin;
-- 1	O	software engineer	21	jalan damai	1
-- \.


--
-- Name: accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gouser
--

SELECT pg_catalog.setval('public.accounts_id_seq', 8, true);


--
-- Name: admins_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gouser
--

SELECT pg_catalog.setval('public.admins_id_seq', 1, true);


--
-- Name: khotbah_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gouser
--

SELECT pg_catalog.setval('public.khotbah_id_seq', 6, true);


--
-- Name: profiles_profile_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gouser
--

SELECT pg_catalog.setval('public.profiles_profile_id_seq', 2, true);


--
-- Name: accounts accounts_email_key; Type: CONSTRAINT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_email_key UNIQUE (email);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: admins admins_pkey; Type: CONSTRAINT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_pkey PRIMARY KEY (id);


--
-- Name: admins admins_username_key; Type: CONSTRAINT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_username_key UNIQUE (username);


--
-- Name: khotbah khotbah_pkey; Type: CONSTRAINT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.khotbah
    ADD CONSTRAINT khotbah_pkey PRIMARY KEY (id);


--
-- Name: profiles profiles_account_id_key; Type: CONSTRAINT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_account_id_key UNIQUE (account_id);


--
-- Name: profiles profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_pkey PRIMARY KEY (profile_id);


--
-- Name: profiles profiles_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: gouser
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.accounts(id);


--
-- PostgreSQL database dump complete
--

