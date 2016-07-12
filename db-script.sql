-- password: globePassword

-- Role: globeAdmin

-- DROP ROLE "globeAdmin";

CREATE ROLE "globeAdmin" LOGIN
  ENCRYPTED PASSWORD 'md523d0ca90a1d2f81ceba1ff4fdae13167'
  NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;


-- Database: "GlobalHiring"

-- DROP DATABASE "GlobalHiring";

CREATE DATABASE "GlobalHiring"
  WITH OWNER = "globAdmin"
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       LC_COLLATE = 'English_United States.1252'
       LC_CTYPE = 'English_United States.1252'
       CONNECTION LIMIT = -1;

-- Table: public."MainService"

-- DROP TABLE public."MainService";

CREATE TABLE "GlobalHiring".public."MainService"
(
 "ID" serial,
 "Name" character varying(100),
 CONSTRAINT "MainService_PK" PRIMARY KEY ("ID"),
 CONSTRAINT "MainService_Name_UK" UNIQUE ("Name")
)
WITH (
 OIDS=FALSE
);
ALTER TABLE public."MainService"
 OWNER TO "globAdmin";


 INSERT INTO public."MainService"("Name") VALUES ('Customer Service');
 INSERT INTO public."MainService"("Name") VALUES ('Admin Support');
 INSERT INTO public."MainService"("Name") VALUES ('Sales & Marketing');
 INSERT INTO public."MainService"("Name") VALUES ('Accounting & Consulting');
 INSERT INTO public."MainService"("Name") VALUES ('Legal');
 INSERT INTO public."MainService"("Name") VALUES ('Translation');
 INSERT INTO public."MainService"("Name") VALUES ('Writing');
 INSERT INTO public."MainService"("Name") VALUES ('Design & Creative');
 INSERT INTO public."MainService"("Name") VALUES ('Engineering & Architecture');
 INSERT INTO public."MainService"("Name") VALUES ('Data Science & Analytics');
 INSERT INTO public."MainService"("Name") VALUES ('IT & Networking');
 INSERT INTO public."MainService"("Name") VALUES ('Web, Mobile & Software Dev');


 -- Table: public."Skill"

-- DROP TABLE public."Skill";

CREATE TABLE public."Skill"
(
  "ID" integer NOT NULL DEFAULT nextval('"Skill_ID_seq"'::regclass),
  "Name" character varying(100) NOT NULL,
  "MainServiceID" integer NOT NULL,
  CONSTRAINT "Skill_PK" PRIMARY KEY ("ID"),
  CONSTRAINT "Skill_MainService_FK" FOREIGN KEY ("MainServiceID")
      REFERENCES public."MainService" ("ID") MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public."Skill"
  OWNER TO "globeAdmin";


INSERT INTO public."Skill"("Name", "MainServiceID") VALUES ('Customer Service', 1);
INSERT INTO public."Skill"("Name", "MainServiceID") VALUES ('Technical Support', 1);
INSERT INTO public."Skill"("Name", "MainServiceID") VALUES ('Other- Customer Service', 1);
