--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = dibs, pg_catalog;

ALTER TABLE ONLY dibs."Shares" DROP CONSTRAINT "UserID";
ALTER TABLE ONLY dibs."Bookings" DROP CONSTRAINT "UserID";
ALTER TABLE ONLY dibs."Shares" DROP CONSTRAINT "ResourceID";
ALTER TABLE ONLY dibs."Bookings" DROP CONSTRAINT "ResourceID";
ALTER TABLE ONLY dibs."Resources" DROP CONSTRAINT "OwnerID";
ALTER TABLE ONLY dibs."Users" DROP CONSTRAINT "Users_pkey";
ALTER TABLE ONLY dibs."Shares" DROP CONSTRAINT "Shares_pkey";
ALTER TABLE ONLY dibs."Resources" DROP CONSTRAINT "Resources_pkey";
ALTER TABLE ONLY dibs."Bookings" DROP CONSTRAINT "Bookings_pkey";
ALTER TABLE dibs."Users" ALTER COLUMN "ID" DROP DEFAULT;
ALTER TABLE dibs."Resources" ALTER COLUMN "ID" DROP DEFAULT;
ALTER TABLE dibs."Bookings" ALTER COLUMN "ID" DROP DEFAULT;
DROP SEQUENCE dibs."Users_ID_seq";
DROP TABLE dibs."Users";
DROP TABLE dibs."Shares";
DROP SEQUENCE dibs."Resources_ID_seq";
DROP TABLE dibs."Resources";
DROP SEQUENCE dibs."Bookings_ID_seq";
DROP TABLE dibs."Bookings";
DROP SCHEMA dibs;
--
-- Name: dibs; Type: SCHEMA; Schema: -; Owner: dibsagent
--

CREATE SCHEMA dibs;


ALTER SCHEMA dibs OWNER TO dibsagent;

SET search_path = dibs, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: Bookings; Type: TABLE; Schema: dibs; Owner: dibsagent; Tablespace: 
--

CREATE TABLE "Bookings" (
    "ID" integer NOT NULL,
    "User" integer NOT NULL,
    "ResourceID" integer NOT NULL,
    "Start" timestamp without time zone NOT NULL,
    "End" timestamp without time zone NOT NULL,
    "Remarks" text
);


ALTER TABLE dibs."Bookings" OWNER TO dibsagent;

--
-- Name: Bookings_ID_seq; Type: SEQUENCE; Schema: dibs; Owner: dibsagent
--

CREATE SEQUENCE "Bookings_ID_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE dibs."Bookings_ID_seq" OWNER TO dibsagent;

--
-- Name: Bookings_ID_seq; Type: SEQUENCE OWNED BY; Schema: dibs; Owner: dibsagent
--

ALTER SEQUENCE "Bookings_ID_seq" OWNED BY "Bookings"."ID";


--
-- Name: Resources; Type: TABLE; Schema: dibs; Owner: dibsagent; Tablespace: 
--

CREATE TABLE "Resources" (
    "ID" integer NOT NULL,
    "Owner" integer NOT NULL,
    "Name" character varying(128) NOT NULL,
    "Description" text
);


ALTER TABLE dibs."Resources" OWNER TO dibsagent;

--
-- Name: Resources_ID_seq; Type: SEQUENCE; Schema: dibs; Owner: dibsagent
--

CREATE SEQUENCE "Resources_ID_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE dibs."Resources_ID_seq" OWNER TO dibsagent;

--
-- Name: Resources_ID_seq; Type: SEQUENCE OWNED BY; Schema: dibs; Owner: dibsagent
--

ALTER SEQUENCE "Resources_ID_seq" OWNED BY "Resources"."ID";


--
-- Name: Shares; Type: TABLE; Schema: dibs; Owner: dibsagent; Tablespace: 
--

CREATE TABLE "Shares" (
    "ResourceID" integer NOT NULL,
    "UserID" integer NOT NULL
);


ALTER TABLE dibs."Shares" OWNER TO dibsagent;

--
-- Name: Users; Type: TABLE; Schema: dibs; Owner: dibsagent; Tablespace: 
--

CREATE TABLE "Users" (
    "ID" integer NOT NULL,
    "Username" character varying(32) NOT NULL,
    "Alias" character varying(128) NOT NULL,
    "Password" character varying(128) NOT NULL,
    "Email" character varying(45)
);


ALTER TABLE dibs."Users" OWNER TO dibsagent;

--
-- Name: Users_ID_seq; Type: SEQUENCE; Schema: dibs; Owner: dibsagent
--

CREATE SEQUENCE "Users_ID_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE dibs."Users_ID_seq" OWNER TO dibsagent;

--
-- Name: Users_ID_seq; Type: SEQUENCE OWNED BY; Schema: dibs; Owner: dibsagent
--

ALTER SEQUENCE "Users_ID_seq" OWNED BY "Users"."ID";


--
-- Name: ID; Type: DEFAULT; Schema: dibs; Owner: dibsagent
--

ALTER TABLE ONLY "Bookings" ALTER COLUMN "ID" SET DEFAULT nextval('"Bookings_ID_seq"'::regclass);


--
-- Name: ID; Type: DEFAULT; Schema: dibs; Owner: dibsagent
--

ALTER TABLE ONLY "Resources" ALTER COLUMN "ID" SET DEFAULT nextval('"Resources_ID_seq"'::regclass);


--
-- Name: ID; Type: DEFAULT; Schema: dibs; Owner: dibsagent
--

ALTER TABLE ONLY "Users" ALTER COLUMN "ID" SET DEFAULT nextval('"Users_ID_seq"'::regclass);


--
-- Data for Name: Bookings; Type: TABLE DATA; Schema: dibs; Owner: dibsagent
--



--
-- Name: Bookings_ID_seq; Type: SEQUENCE SET; Schema: dibs; Owner: dibsagent
--

SELECT pg_catalog.setval('"Bookings_ID_seq"', 1, false);


--
-- Data for Name: Resources; Type: TABLE DATA; Schema: dibs; Owner: dibsagent
--



--
-- Name: Resources_ID_seq; Type: SEQUENCE SET; Schema: dibs; Owner: dibsagent
--

SELECT pg_catalog.setval('"Resources_ID_seq"', 1, false);


--
-- Data for Name: Shares; Type: TABLE DATA; Schema: dibs; Owner: dibsagent
--



--
-- Data for Name: Users; Type: TABLE DATA; Schema: dibs; Owner: dibsagent
--

INSERT INTO "Users" VALUES (1, 'aburian', 'Andrew Burian', '$2a$10$llL.U6kVd6ZJIbiWtnRqrOFLC612LTu5cXGWyu51N2u4sqaxdgUji', 'andrew@burian.ca');


--
-- Name: Users_ID_seq; Type: SEQUENCE SET; Schema: dibs; Owner: dibsagent
--

SELECT pg_catalog.setval('"Users_ID_seq"', 1, true);


--
-- Name: Bookings_pkey; Type: CONSTRAINT; Schema: dibs; Owner: dibsagent; Tablespace: 
--

ALTER TABLE ONLY "Bookings"
    ADD CONSTRAINT "Bookings_pkey" PRIMARY KEY ("ID");


--
-- Name: Resources_pkey; Type: CONSTRAINT; Schema: dibs; Owner: dibsagent; Tablespace: 
--

ALTER TABLE ONLY "Resources"
    ADD CONSTRAINT "Resources_pkey" PRIMARY KEY ("ID");


--
-- Name: Shares_pkey; Type: CONSTRAINT; Schema: dibs; Owner: dibsagent; Tablespace: 
--

ALTER TABLE ONLY "Shares"
    ADD CONSTRAINT "Shares_pkey" PRIMARY KEY ("ResourceID", "UserID");


--
-- Name: Users_pkey; Type: CONSTRAINT; Schema: dibs; Owner: dibsagent; Tablespace: 
--

ALTER TABLE ONLY "Users"
    ADD CONSTRAINT "Users_pkey" PRIMARY KEY ("ID");


--
-- Name: OwnerID; Type: FK CONSTRAINT; Schema: dibs; Owner: dibsagent
--

ALTER TABLE ONLY "Resources"
    ADD CONSTRAINT "OwnerID" FOREIGN KEY ("Owner") REFERENCES "Users"("ID");


--
-- Name: ResourceID; Type: FK CONSTRAINT; Schema: dibs; Owner: dibsagent
--

ALTER TABLE ONLY "Bookings"
    ADD CONSTRAINT "ResourceID" FOREIGN KEY ("ResourceID") REFERENCES "Resources"("ID");


--
-- Name: ResourceID; Type: FK CONSTRAINT; Schema: dibs; Owner: dibsagent
--

ALTER TABLE ONLY "Shares"
    ADD CONSTRAINT "ResourceID" FOREIGN KEY ("ResourceID") REFERENCES "Resources"("ID");


--
-- Name: UserID; Type: FK CONSTRAINT; Schema: dibs; Owner: dibsagent
--

ALTER TABLE ONLY "Bookings"
    ADD CONSTRAINT "UserID" FOREIGN KEY ("User") REFERENCES "Users"("ID");


--
-- Name: UserID; Type: FK CONSTRAINT; Schema: dibs; Owner: dibsagent
--

ALTER TABLE ONLY "Shares"
    ADD CONSTRAINT "UserID" FOREIGN KEY ("UserID") REFERENCES "Users"("ID");


--
-- PostgreSQL database dump complete
--

