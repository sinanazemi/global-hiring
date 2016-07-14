-- password: globePassword

-- Role: globeAdmin

-- DROP ROLE "globeAdmin";

CREATE ROLE globeAdmin LOGIN
  ENCRYPTED PASSWORD 'md5e60cf860535525c970ebac4e32e50184'
  NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;


-- Database: "GlobalHiring"

-- DROP DATABASE "GlobalHiring";

CREATE DATABASE GlobalHiring
  WITH OWNER = globeAdmin
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       LC_COLLATE = 'English_United States.1252'
       LC_CTYPE = 'English_United States.1252'
       CONNECTION LIMIT = -1;

-- Table: public."MainService"

-- DROP TABLE public."MainService";

CREATE TABLE GlobalHiring.MainService
(
 ID serial,
 Name character varying(100),
 CONSTRAINT MainService_PK PRIMARY KEY (ID),
 CONSTRAINT MainService_Name_UK UNIQUE (Name)
)
WITH (
 OIDS=FALSE
);
ALTER TABLE MainService
 OWNER TO globeAdmin;


 INSERT INTO MainService(Name) VALUES ('Customer Service');
 INSERT INTO MainService(Name) VALUES ('Admin Support');
 INSERT INTO MainService(Name) VALUES ('Sales & Marketing');
 INSERT INTO MainService(Name) VALUES ('Accounting & Consulting');
 INSERT INTO MainService(Name) VALUES ('Legal');
 INSERT INTO MainService(Name) VALUES ('Translation');
 INSERT INTO MainService(Name) VALUES ('Writing');
 INSERT INTO MainService(Name) VALUES ('Design & Creative');
 INSERT INTO MainService(Name) VALUES ('Engineering & Architecture');
 INSERT INTO MainService(Name) VALUES ('Data Science & Analytics');
 INSERT INTO MainService(Name) VALUES ('IT & Networking');
 INSERT INTO MainService(Name) VALUES ('Web, Mobile & Software Dev');


 -- Table: public."Skill"

-- DROP TABLE public."Skill";

CREATE TABLE Skill
(
  ID serial,
  Name character varying(100) NOT NULL,
  MainServiceID integer NOT NULL,
  CONSTRAINT Skill_PK PRIMARY KEY (ID),
  CONSTRAINT Skill_MainService_FK FOREIGN KEY (MainServiceID)
      REFERENCES MainService (ID) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
ALTER TABLE Skill
  OWNER TO globeAdmin;


INSERT INTO Skill(Name, MainServiceID) VALUES ('Customer Service', 1);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Technical Support', 1);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Other- Customer Service', 1);

INSERT INTO Skill(Name, MainServiceID) VALUES ('Data Entry' , 2);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Personal / Virtual Assistant' , 2);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Project Management' , 2);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Transcription' , 2);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Web Research' , 2);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Other - Admin Support' , 2);

INSERT INTO Skill(Name, MainServiceID) VALUES ('Display Advertising' , 3);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Email & Marketing Automation' , 3);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Lead Generation' , 3);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Market & Customer Research' , 3);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Marketing Strategy' , 3);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Public Relations' , 3);
INSERT INTO Skill(Name, MainServiceID) VALUES ('SEM - Search Engine Marketing' , 3);
INSERT INTO Skill(Name, MainServiceID) VALUES ('SEO - Search Engine Optimization' , 3);
INSERT INTO Skill(Name, MainServiceID) VALUES ('SMM - Social Media Marketing' , 3);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Telemarketing & Telesales' , 3);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Other - Sales & Marketing' , 3);

INSERT INTO Skill(Name, MainServiceID) VALUES ('Accounting' , 4);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Financial Planning' , 4);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Human Resources' , 4);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Management Consulting' , 4);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Other - Accounting & Consulting' , 4);

INSERT INTO Skill(Name, MainServiceID) VALUES ('Contract Law' , 5);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Corporate Law' , 5);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Criminal Law' , 5);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Family Law' , 5);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Intellectual Property Law' , 5);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Paralegal Services' , 5);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Other - Legal' , 5);

INSERT INTO Skill(Name, MainServiceID) VALUES ('General Translation' , 6);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Legal Translation' , 6);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Medical Translation' , 6);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Technical Translation' , 6);

INSERT INTO Skill(Name, MainServiceID) VALUES ('Academic Writing & Research' , 7);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Article & Blog Writing' , 7);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Copywriting' , 7);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Creative Writing' , 7);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Editing & Proofreading' , 7);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Grant Writing' , 7);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Resumes & Cover Letters' , 7);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Technical Writing' , 7);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Web Content' , 7);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Other - Writing' , 7);

INSERT INTO Skill(Name, MainServiceID) VALUES ('Animation' , 8);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Audio Production' , 8);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Graphic Design' , 8);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Illustraion' , 8);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Logo Design & Branding' , 8);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Photography' , 8);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Presentaions' , 8);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Video Production' , 8);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Voice Talent' , 8);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Other - Design & Creative' , 8);

INSERT INTO Skill(Name, MainServiceID) VALUES ('3d Modeling & CAD' , 9);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Architecture' , 9);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Chemical Engineering' , 9);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Civil & Structural Engineering' , 9);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Contract Manufacturing' , 9);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Electrical Engineering' , 9);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Interior Design' , 9);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Mechanical Engineering' , 9);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Product Design' , 9);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Other - Engineering & Architecture' , 9);

INSERT INTO Skill(Name, MainServiceID) VALUES ('A/B Testing' , 10);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Data Visualization' , 10);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Data Extraction / ETL' , 10);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Data Mining & Management' , 10);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Machine Learning' , 10);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Quantative Analysis' , 10);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Other - Data Science & Analytics' , 10);

INSERT INTO Skill(Name, MainServiceID) VALUES ('Database Administration' , 11);
INSERT INTO Skill(Name, MainServiceID) VALUES ('ERP / CRM Software' , 11);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Information Security' , 11);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Network & System Administration' , 11);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Other - IT & Networking' , 11);

INSERT INTO Skill(Name, MainServiceID) VALUES ('Desktop Software Development' , 12);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Ecommerce Development' , 12);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Game Development' , 12);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Mobile Development' , 12);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Product Management' , 12);
INSERT INTO Skill(Name, MainServiceID) VALUES ('QA & Testing' , 12);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Scripts & Utilities' , 12);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Web Development' , 12);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Web & Mobile Design' , 12);
INSERT INTO Skill(Name, MainServiceID) VALUES ('Other - Software Development' , 12);
