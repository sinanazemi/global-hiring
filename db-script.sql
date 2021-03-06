/*
delete from accountcertificate;
delete from accountcourse;
delete from accounteducation;
delete from accounthonor;
delete from accountlanguage;
delete from accountproject;
delete from accountskill;
delete from accounttest;
delete from accountvolunteering;
delete from accountwork;
delete from account;
*/


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

CREATE TABLE MainService
(
 ID serial,
 Name character varying(100) not null,
 Question character varying(300),
 CONSTRAINT MainService_PK PRIMARY KEY (ID),
 CONSTRAINT MainService_Name_UK UNIQUE (Name)
)
WITH (
 OIDS=FALSE
);
ALTER TABLE MainService
 OWNER TO globeAdmin;


 INSERT INTO MainService(Name, Question) VALUES ('Customer Service', 'What type of Customer Service do you do?');
 INSERT INTO MainService(Name, Question) VALUES ('Admin Support', 'What type of Admin Support do you do?');
 INSERT INTO MainService(Name, Question) VALUES ('Sales & Marketing', 'What type of Sales & Marketing do you do?');
 INSERT INTO MainService(Name, Question) VALUES ('Accounting & Consulting', 'What type of Accounting & Consulting do you do?');
 INSERT INTO MainService(Name, Question) VALUES ('Legal', 'What type of Legal do you do?');
 INSERT INTO MainService(Name, Question) VALUES ('Translation', 'What type of Translation do you do?');
 INSERT INTO MainService(Name, Question) VALUES ('Writing', 'What type of Writing do you do?');
 INSERT INTO MainService(Name, Question) VALUES ('Design & Creative', 'What type of Design & Creative do you do?');
 INSERT INTO MainService(Name, Question) VALUES ('Engineering & Architecture', 'What type of Engineering & Architecture do you do?');
 INSERT INTO MainService(Name, Question) VALUES ('Data Science & Analytics', 'What type of Data Science & Analytics do you do?');
 INSERT INTO MainService(Name, Question) VALUES ('IT & Networking', 'What type of IT & Networking do you do?');
 INSERT INTO MainService(Name, Question) VALUES ('Web, Mobile & Software Dev', 'What type of Web, Mobile & Software Dev do you do?');


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

-- Table: public.city

-- DROP TABLE public.city;

CREATE TABLE city
(
  id serial,
  name character varying(100) NOT NULL,
  CONSTRAINT city_pk PRIMARY KEY (id),
  CONSTRAINT city_Name_UK UNIQUE (Name)
 )
WITH (
  OIDS=FALSE
);
ALTER TABLE city
  OWNER TO globeadmin;


INSERT INTO city(Name) VALUES ('Abadan');
INSERT INTO city(Name) VALUES ('Abadeh');
INSERT INTO city(Name) VALUES ('Abyek');
INSERT INTO city(Name) VALUES ('Abhar');
INSERT INTO city(Name) VALUES ('Abyaneh');
INSERT INTO city(Name) VALUES ('Ahar');
INSERT INTO city(Name) VALUES ('Ahvaz');
INSERT INTO city(Name) VALUES ('Alavicheh');
INSERT INTO city(Name) VALUES ('Aliabad');
INSERT INTO city(Name) VALUES ('Aligoodarz');
INSERT INTO city(Name) VALUES ('Alvand');
INSERT INTO city(Name) VALUES ('Amlash');
INSERT INTO city(Name) VALUES ('Amol');
INSERT INTO city(Name) VALUES ('Andimeshk');
INSERT INTO city(Name) VALUES ('Andisheh');
INSERT INTO city(Name) VALUES ('Arak');
INSERT INTO city(Name) VALUES ('Ardabil');
INSERT INTO city(Name) VALUES ('Ardakan');
INSERT INTO city(Name) VALUES ('Asalem');
INSERT INTO city(Name) VALUES ('Asalouyeh');
INSERT INTO city(Name) VALUES ('Ashkezar');
INSERT INTO city(Name) VALUES ('Ashlagh');
INSERT INTO city(Name) VALUES ('Ashtiyan');
INSERT INTO city(Name) VALUES ('Astaneh Arak');
INSERT INTO city(Name) VALUES ('Astaneh-e Ashrafiyyeh');
INSERT INTO city(Name) VALUES ('Astara');
INSERT INTO city(Name) VALUES ('Babol');
INSERT INTO city(Name) VALUES ('Babolsar');
INSERT INTO city(Name) VALUES ('Baharestan');
INSERT INTO city(Name) VALUES ('Balov');
INSERT INTO city(Name) VALUES ('Bardaskan');
INSERT INTO city(Name) VALUES ('Bam');
INSERT INTO city(Name) VALUES ('Bampur');
INSERT INTO city(Name) VALUES ('Bandar Abbas');
INSERT INTO city(Name) VALUES ('Bandar Anzali');
INSERT INTO city(Name) VALUES ('Bandar Charak');
INSERT INTO city(Name) VALUES ('Bandar Imam');
INSERT INTO city(Name) VALUES ('Bandar Lengeh');
INSERT INTO city(Name) VALUES ('Bandar Torkman');
INSERT INTO city(Name) VALUES ('Baneh');
INSERT INTO city(Name) VALUES ('Bastak');
INSERT INTO city(Name) VALUES ('Behbahan');
INSERT INTO city(Name) VALUES ('Behshahr');
INSERT INTO city(Name) VALUES ('Bijar');
INSERT INTO city(Name) VALUES ('Birjand');
INSERT INTO city(Name) VALUES ('Bistam');
INSERT INTO city(Name) VALUES ('Bojnourd');
INSERT INTO city(Name) VALUES ('Bonab');
INSERT INTO city(Name) VALUES ('Borazjan');
INSERT INTO city(Name) VALUES ('Borujerd');
INSERT INTO city(Name) VALUES ('Bukan');
INSERT INTO city(Name) VALUES ('Bushehr');
INSERT INTO city(Name) VALUES ('Chabahar');
INSERT INTO city(Name) VALUES ('Damavand');
INSERT INTO city(Name) VALUES ('Damghan');
INSERT INTO city(Name) VALUES ('Darab');
INSERT INTO city(Name) VALUES ('Dargaz');
INSERT INTO city(Name) VALUES ('Daryan');
INSERT INTO city(Name) VALUES ('Darreh Shahr');
INSERT INTO city(Name) VALUES ('Deylam');
INSERT INTO city(Name) VALUES ('Deyr');
INSERT INTO city(Name) VALUES ('Dezful');
INSERT INTO city(Name) VALUES ('Dezghan');
INSERT INTO city(Name) VALUES ('Dibaj');
INSERT INTO city(Name) VALUES ('Doroud');
INSERT INTO city(Name) VALUES ('Eghlid');
INSERT INTO city(Name) VALUES ('Esfahan');
INSERT INTO city(Name) VALUES ('Esfarayen');
INSERT INTO city(Name) VALUES ('Eslamabad');
INSERT INTO city(Name) VALUES ('Eslamabad-e Gharb');
INSERT INTO city(Name) VALUES ('Eslamshahr');
INSERT INTO city(Name) VALUES ('Evaz');
INSERT INTO city(Name) VALUES ('Farahan');
INSERT INTO city(Name) VALUES ('Fasa');
INSERT INTO city(Name) VALUES ('Ferdows');
INSERT INTO city(Name) VALUES ('Feshak');
INSERT INTO city(Name) VALUES ('Feshk');
INSERT INTO city(Name) VALUES ('Firouzabad');
INSERT INTO city(Name) VALUES ('Fouman');
INSERT INTO city(Name) VALUES ('Fasham, Tehran');
INSERT INTO city(Name) VALUES ('Gachsaran');
INSERT INTO city(Name) VALUES ('Garmeh-Jajarm');
INSERT INTO city(Name) VALUES ('Gavrik');
INSERT INTO city(Name) VALUES ('Ghale Ganj');
INSERT INTO city(Name) VALUES ('Gerash');
INSERT INTO city(Name) VALUES ('Genaveh');
INSERT INTO city(Name) VALUES ('Ghaemshahr');
INSERT INTO city(Name) VALUES ('Golbahar');
INSERT INTO city(Name) VALUES ('Golpayegan');
INSERT INTO city(Name) VALUES ('Gonabad');
INSERT INTO city(Name) VALUES ('Gonbad-e Kavous');
INSERT INTO city(Name) VALUES ('Gorgan');
INSERT INTO city(Name) VALUES ('Hamadan');
INSERT INTO city(Name) VALUES ('Hashtgerd');
INSERT INTO city(Name) VALUES ('Hashtpar');
INSERT INTO city(Name) VALUES ('Hashtrud');
INSERT INTO city(Name) VALUES ('Heris');
INSERT INTO city(Name) VALUES ('Hidaj');
INSERT INTO city(Name) VALUES ('Haji Abad');
INSERT INTO city(Name) VALUES ('Ij');
INSERT INTO city(Name) VALUES ('Ilam');
INSERT INTO city(Name) VALUES ('Iranshahr');
INSERT INTO city(Name) VALUES ('Islamshahr');
INSERT INTO city(Name) VALUES ('Izadkhast');
INSERT INTO city(Name) VALUES ('Izeh');
INSERT INTO city(Name) VALUES ('Jajarm');
INSERT INTO city(Name) VALUES ('Jask');
INSERT INTO city(Name) VALUES ('Jahrom');
INSERT INTO city(Name) VALUES ('Jaleq');
INSERT INTO city(Name) VALUES ('Javanrud');
INSERT INTO city(Name) VALUES ('Jiroft');
INSERT INTO city(Name) VALUES ('Jolfa');
INSERT INTO city(Name) VALUES ('Kahnuj');
INSERT INTO city(Name) VALUES ('Kamyaran');
INSERT INTO city(Name) VALUES ('Kangan');
INSERT INTO city(Name) VALUES ('Kangavar');
INSERT INTO city(Name) VALUES ('Karaj');
INSERT INTO city(Name) VALUES ('Kashan');
INSERT INTO city(Name) VALUES ('Kashmar');
INSERT INTO city(Name) VALUES ('Kazeroun');
INSERT INTO city(Name) VALUES ('Kerman');
INSERT INTO city(Name) VALUES ('Kermanshah');
INSERT INTO city(Name) VALUES ('Khalkhal');
INSERT INTO city(Name) VALUES ('Khalkhāl');
INSERT INTO city(Name) VALUES ('Khomein');
INSERT INTO city(Name) VALUES ('Khomeynishahr');
INSERT INTO city(Name) VALUES ('Khonj');
INSERT INTO city(Name) VALUES ('Khormuj');
INSERT INTO city(Name) VALUES ('Khorramabad');
INSERT INTO city(Name) VALUES ('Khorramshahr');
INSERT INTO city(Name) VALUES ('Khorashad');
INSERT INTO city(Name) VALUES ('Koumleh');
INSERT INTO city(Name) VALUES ('Khvoy');
INSERT INTO city(Name) VALUES ('Kilan');
INSERT INTO city(Name) VALUES ('Kish');
INSERT INTO city(Name) VALUES ('Koker');
INSERT INTO city(Name) VALUES ('Kosar');
INSERT INTO city(Name) VALUES ('Kordkuy');
INSERT INTO city(Name) VALUES ('Kong');
INSERT INTO city(Name) VALUES ('Kuhdasht');
INSERT INTO city(Name) VALUES ('Laft');
INSERT INTO city(Name) VALUES ('Lahijan');
INSERT INTO city(Name) VALUES ('Langaroud');
INSERT INTO city(Name) VALUES ('Lar');
INSERT INTO city(Name) VALUES ('Latian');
INSERT INTO city(Name) VALUES ('Lavasan');
INSERT INTO city(Name) VALUES ('Lamerd');
INSERT INTO city(Name) VALUES ('Mahabad');
INSERT INTO city(Name) VALUES ('Mahan');
INSERT INTO city(Name) VALUES ('Mahshahr');
INSERT INTO city(Name) VALUES ('Majlesi');
INSERT INTO city(Name) VALUES ('Maku');
INSERT INTO city(Name) VALUES ('Malard');
INSERT INTO city(Name) VALUES ('Malayer');
INSERT INTO city(Name) VALUES ('Manjil');
INSERT INTO city(Name) VALUES ('Manoojan');
INSERT INTO city(Name) VALUES ('Maragheh');
INSERT INTO city(Name) VALUES ('Marand');
INSERT INTO city(Name) VALUES ('Marivan');
INSERT INTO city(Name) VALUES ('Marvdasht');
INSERT INTO city(Name) VALUES ('Masal');
INSERT INTO city(Name) VALUES ('Mashhad');
INSERT INTO city(Name) VALUES ('Masjed Soleyman');
INSERT INTO city(Name) VALUES ('Mehran');
INSERT INTO city(Name) VALUES ('Meshkinshahr');
INSERT INTO city(Name) VALUES ('Meyaneh');
INSERT INTO city(Name) VALUES ('Meybod');
INSERT INTO city(Name) VALUES ('Miandoab');
INSERT INTO city(Name) VALUES ('Mianeh');
INSERT INTO city(Name) VALUES ('Mianeh-ye Bardangan');
INSERT INTO city(Name) VALUES ('Mianej');
INSERT INTO city(Name) VALUES ('Minab');
INSERT INTO city(Name) VALUES ('Minoodasht');
INSERT INTO city(Name) VALUES ('Mohajeran');
INSERT INTO city(Name) VALUES ('Naghadeh');
INSERT INTO city(Name) VALUES ('Nobandeyaan');
INSERT INTO city(Name) VALUES ('Nahavand');
INSERT INTO city(Name) VALUES ('Nain');
INSERT INTO city(Name) VALUES ('Najafabad');
INSERT INTO city(Name) VALUES ('Namin');
INSERT INTO city(Name) VALUES ('Natanz');
INSERT INTO city(Name) VALUES ('Nazarabad');
INSERT INTO city(Name) VALUES ('Nishapur');
INSERT INTO city(Name) VALUES ('Nīr');
INSERT INTO city(Name) VALUES ('Nowshahr');
INSERT INTO city(Name) VALUES ('Nurabad');
INSERT INTO city(Name) VALUES ('Omidiyeh');
INSERT INTO city(Name) VALUES ('Oshnaviyeh');
INSERT INTO city(Name) VALUES ('Oskou');
INSERT INTO city(Name) VALUES ('Ormand');
INSERT INTO city(Name) VALUES ('Orumiyeh');
INSERT INTO city(Name) VALUES ('Pakdasht');
INSERT INTO city(Name) VALUES ('Parand');
INSERT INTO city(Name) VALUES ('Pardis');
INSERT INTO city(Name) VALUES ('Parsabad');
INSERT INTO city(Name) VALUES ('Paveh');
INSERT INTO city(Name) VALUES ('Piranshahr');
INSERT INTO city(Name) VALUES ('Pishva');
INSERT INTO city(Name) VALUES ('Poldasht');
INSERT INTO city(Name) VALUES ('Poulad-shahr');
INSERT INTO city(Name) VALUES ('Qaemshahr');
INSERT INTO city(Name) VALUES ('Qaen');
INSERT INTO city(Name) VALUES ('Qamsar');
INSERT INTO city(Name) VALUES ('Qasr-e Shirin');
INSERT INTO city(Name) VALUES ('Qazvin');
INSERT INTO city(Name) VALUES ('Qods');
INSERT INTO city(Name) VALUES ('Qom');
INSERT INTO city(Name) VALUES ('Qorveh');
INSERT INTO city(Name) VALUES ('Quchan');
INSERT INTO city(Name) VALUES ('Rafsanjan');
INSERT INTO city(Name) VALUES ('Ramin');
INSERT INTO city(Name) VALUES ('Ramsar');
INSERT INTO city(Name) VALUES ('Ramshar');
INSERT INTO city(Name) VALUES ('Rasht');
INSERT INTO city(Name) VALUES ('Ray');
INSERT INTO city(Name) VALUES ('Razmian');
INSERT INTO city(Name) VALUES ('Rezvanshahr');
INSERT INTO city(Name) VALUES ('Roudbar');
INSERT INTO city(Name) VALUES ('Roodbar-e-Jonoub');
INSERT INTO city(Name) VALUES ('Roudsar');
INSERT INTO city(Name) VALUES ('Runiz');
INSERT INTO city(Name) VALUES ('Sabzevar');
INSERT INTO city(Name) VALUES ('Sadra');
INSERT INTO city(Name) VALUES ('Sahand');
INSERT INTO city(Name) VALUES ('Salmas');
INSERT INTO city(Name) VALUES ('Sanandaj');
INSERT INTO city(Name) VALUES ('Saqqez');
INSERT INTO city(Name) VALUES ('Sarab');
INSERT INTO city(Name) VALUES ('Sarableh');
INSERT INTO city(Name) VALUES ('Sarakhs');
INSERT INTO city(Name) VALUES ('Saravan');
INSERT INTO city(Name) VALUES ('Sardasht');
INSERT INTO city(Name) VALUES ('Sari');
INSERT INTO city(Name) VALUES ('Sarvestan');
INSERT INTO city(Name) VALUES ('Saveh');
INSERT INTO city(Name) VALUES ('Senejan');
INSERT INTO city(Name) VALUES ('Semnan');
INSERT INTO city(Name) VALUES ('Shabestar');
INSERT INTO city(Name) VALUES ('Shaft');
INSERT INTO city(Name) VALUES ('Shahinshahr');
INSERT INTO city(Name) VALUES ('Shahr-e Kord');
INSERT INTO city(Name) VALUES ('Shahrezā');
INSERT INTO city(Name) VALUES ('Shahriar');
INSERT INTO city(Name) VALUES ('Shahroud');
INSERT INTO city(Name) VALUES ('Shahsavar');
INSERT INTO city(Name) VALUES ('Shiraz');
INSERT INTO city(Name) VALUES ('Shirvan');
INSERT INTO city(Name) VALUES ('Shushtar');
INSERT INTO city(Name) VALUES ('Siahkal');
INSERT INTO city(Name) VALUES ('Sirjan');
INSERT INTO city(Name) VALUES ('Sourmagh');
INSERT INTO city(Name) VALUES ('Sowme-e-Sara');
INSERT INTO city(Name) VALUES ('Sarpole Zahab');
INSERT INTO city(Name) VALUES ('Tabas');
INSERT INTO city(Name) VALUES ('Tabriz');
INSERT INTO city(Name) VALUES ('Tafresh');
INSERT INTO city(Name) VALUES ('Taft');
INSERT INTO city(Name) VALUES ('Takab');
INSERT INTO city(Name) VALUES ('Tehran');
INSERT INTO city(Name) VALUES ('Torqabeh');
INSERT INTO city(Name) VALUES ('Torbat-e Heydarieh');
INSERT INTO city(Name) VALUES ('Torbat-e Jam');
INSERT INTO city(Name) VALUES ('Touyserkan');
INSERT INTO city(Name) VALUES ('Tous');
INSERT INTO city(Name) VALUES ('Tonekabon');
INSERT INTO city(Name) VALUES ('Varamin');
INSERT INTO city(Name) VALUES ('Yasouj');
INSERT INTO city(Name) VALUES ('Yazd');
INSERT INTO city(Name) VALUES ('Zabol');
INSERT INTO city(Name) VALUES ('Zahedan');
INSERT INTO city(Name) VALUES ('Zanjan');
INSERT INTO city(Name) VALUES ('Zarand');
INSERT INTO city(Name) VALUES ('Zarrinshahr');


-- Table: Account

-- DROP TABLE Account;

CREATE TABLE Account
(
  id serial,
  name character varying(100) NOT NULL,
  email character varying(100) NOT NULL,
  cityID integer NOT NULL,
  phone character varying(30) NOT NULL,
  password character varying(100) NOT NULL,
  Description character varying(2000),
  JobTitle character varying(2000),
  isStudent BOOLEAN DEFAULT FALSE NOT NULL,
  image bytea NOT NULL,
  CONSTRAINT account_pk PRIMARY KEY (id),
  CONSTRAINT account_city_fk FOREIGN KEY (cityID)
      REFERENCES city (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
ALTER TABLE account
  OWNER TO globeadmin;


-- Table: AccountLanguage

-- DROP TABLE AccountLanguage;

CREATE TABLE AccountLanguage
(
 ID serial,
 Name character varying(100) not null,
 Profeciency character (1) not null,
 accountID integer NOT NULL,
 CONSTRAINT AccountLanguage_PK PRIMARY KEY (ID),
 CONSTRAINT AccountLanguage_Account_FK FOREIGN KEY (accountID)
     REFERENCES account (id) MATCH SIMPLE
     ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
 OIDS=FALSE
);
ALTER TABLE AccountLanguage
 OWNER TO globeAdmin;

-- Table: Degree

-- DROP TABLE Degree;

CREATE TABLE Degree
(
  ID serial,
  Name character varying(100) not null,
  CONSTRAINT Degree_PK PRIMARY KEY (ID),
  CONSTRAINT Degree_Name_UK UNIQUE (Name)
)
WITH (
OIDS=FALSE
);
ALTER TABLE Degree
OWNER TO globeAdmin;

insert into degree (Name) values ('other');
insert into degree (Name) values ('High School');
insert into degree (Name) values ('Associate`s Degree');
insert into degree (Name) values ('Bachelor`s Degree');
insert into degree (Name) values ('Master`s Degree');
insert into degree (Name) values ('Master of Business Administration (M.B.A.)');
insert into degree (Name) values ('Juris Doctor (J.D.)');
insert into degree (Name) values ('Doctor of Medicine (M.D.)');
insert into degree (Name) values ('Doctor of Philosophy (Ph.D.)');
insert into degree (Name) values ('Engineer`s Degree');


-- Table: AccountEducation

-- DROP TABLE AccountEducation;

create TABLE AccountEducation
(
 ID serial,
 School character varying(100) not null,
 FromDate integer NOT NULL,
 ToDate integer NOT NULL,
 DegreeID integer NOT NULL,
 Field character varying(2000) not null,
 Grade NUMERIC(5, 2),
 accountID integer NOT NULL,
 CONSTRAINT AccountEducation_PK PRIMARY KEY (ID),
 CONSTRAINT AccountEducation_Degree_FK FOREIGN KEY (DegreeID)
     REFERENCES Degree (id) MATCH SIMPLE
     ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT AccountEducation_Account_FK FOREIGN KEY (accountID)
     REFERENCES account (id) MATCH SIMPLE
     ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
 OIDS=FALSE
);
ALTER TABLE AccountEducation
 OWNER TO globeAdmin;


-- Table: AccountSkill

-- DROP TABLE AccountSkill;

create TABLE AccountSkill
(
ID serial,
accountID integer NOT NULL,
skillID integer NOT NULL,
Profeciency character (1) not null,
CONSTRAINT AccountSkill_PK PRIMARY KEY (ID),
CONSTRAINT AccountSkill_Account_FK FOREIGN KEY (accountID)
    REFERENCES account (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION,
CONSTRAINT AccountSkill_Skill_FK FOREIGN KEY (skillID)
    REFERENCES skill (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
OIDS=FALSE
);
ALTER TABLE AccountSkill
OWNER TO globeAdmin;

-- Table: AccountCertificate

-- DROP TABLE AccountCertificate;

create TABLE AccountCertificate
(
ID serial,
accountID integer NOT NULL,
Name character varying(200) not null,
Authority character varying(200) not null,
License character varying(200) not null,
Url character varying(300) not null,
Description character varying(2000),
CONSTRAINT AccountCertificate_PK PRIMARY KEY (ID),
CONSTRAINT AccountCertificate_Account_FK FOREIGN KEY (accountID)
    REFERENCES account (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
OIDS=FALSE
);
ALTER TABLE AccountCertificate
OWNER TO globeAdmin;


-- Table: AccountWork

-- DROP TABLE AccountWork;

create TABLE AccountWork
(
ID serial,
Company character varying(200) not null,
Location character varying(600) not null,
Title character varying(200) not null,
Role character(1) not null,
FromMonth integer NOT NULL,
FromYear integer NOT NULL,
ToMonth integer,
ToYear integer,
currently boolean not null,
Description character varying(2000),
accountID integer NOT NULL,
CONSTRAINT AccountWork_PK PRIMARY KEY (ID),
CONSTRAINT AccountWork_Account_FK FOREIGN KEY (accountID)
    REFERENCES account (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
OIDS=FALSE
);
ALTER TABLE AccountWork
OWNER TO globeAdmin;

-- Table: VolunteeringCause

-- DROP TABLE VolunteeringCause;

CREATE TABLE VolunteeringCause
(
  ID serial,
  Name character varying(200) not null,
  CONSTRAINT VolunteeringCause_PK PRIMARY KEY (ID),
  CONSTRAINT VolunteeringCause_Name_UK UNIQUE (Name)
)
WITH (
OIDS=FALSE
);
ALTER TABLE VolunteeringCause
OWNER TO globeAdmin;

insert into VolunteeringCause (Name) values ('Animal Welfare');
insert into VolunteeringCause (Name) values ('Arts and Culture');
insert into VolunteeringCause (Name) values ('Children');
insert into VolunteeringCause (Name) values ('Civil Rights and Social Action');
insert into VolunteeringCause (Name) values ('Disaster and Humanitarian Releif');
insert into VolunteeringCause (Name) values ('Economic Empowerment');
insert into VolunteeringCause (Name) values ('Education');
insert into VolunteeringCause (Name) values ('Environment');
insert into VolunteeringCause (Name) values ('Health');
insert into VolunteeringCause (Name) values ('Human Rights');
insert into VolunteeringCause (Name) values ('Politics');
insert into VolunteeringCause (Name) values ('Poverty Alleviation');
insert into VolunteeringCause (Name) values ('Science and technology');
insert into VolunteeringCause (Name) values ('Social Services');

-- Table: AccountVolunteering

-- DROP TABLE AccountVolunteering;

create TABLE AccountVolunteering
(
ID serial,
Organization character varying(200) not null,
Role character varying(200) not null,
Cause integer not null,
FromMonth integer NOT NULL,
FromYear integer NOT NULL,
ToMonth integer,
ToYear integer,
Description character varying(2000),
accountID integer NOT NULL,
CONSTRAINT AccountVolunteering_PK PRIMARY KEY (ID),
CONSTRAINT AccountVolunteering_Account_FK FOREIGN KEY (accountID)
    REFERENCES account (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION,
CONSTRAINT AccountVolunteering_Cause_FK FOREIGN KEY (Cause)
    REFERENCES VolunteeringCause (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
OIDS=FALSE
);
ALTER TABLE AccountVolunteering
OWNER TO globeAdmin;

-- View: Occupation

-- DROP VIEW Occupation;

create or replace view Occupation as
select 0 as ID, -1 as accountID, 'Other' as Name
union
select (1 + 10*ID) as ID, accountID, school as name from accounteducation
union
select (2 + 10*ID) as ID, accountID, company as name from accountwork;
ALTER VIEW Occupation
OWNER TO globeAdmin;

-- Table: AccountCourse

-- DROP TABLE AccountCourse;

create TABLE AccountCourse
(
ID serial,
Name character varying(100) not null,
Number character varying(100) not null,
OccupationID integer not null,
Description character varying(2000),
accountID integer NOT NULL,
CONSTRAINT AccountCourse_PK PRIMARY KEY (ID),
CONSTRAINT AccountCourse_Account_FK FOREIGN KEY (accountID)
    REFERENCES account (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
OIDS=FALSE
);
ALTER TABLE AccountCourse
OWNER TO globeAdmin;

-- Table: AccountHonor

-- DROP TABLE AccountHonor;

create TABLE AccountHonor
(
ID serial,
Title character varying(100) not null,
OccupationID integer not null,
Month integer NOT NULL,
Year integer NOT NULL,
Description character varying(2000),
accountID integer NOT NULL,
CONSTRAINT AccountHonor_PK PRIMARY KEY (ID),
CONSTRAINT AccountHonor_Account_FK FOREIGN KEY (accountID)
    REFERENCES account (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
OIDS=FALSE
);
ALTER TABLE AccountHonor
OWNER TO globeAdmin;


-- Table: AccountTest

-- DROP TABLE AccountTest;

create TABLE AccountTest
(
ID serial,
Name character varying(100) not null,
OccupationID integer not null,
Month integer NOT NULL,
Year integer NOT NULL,
Score NUMERIC(5, 2) NOT Null,
Description character varying(2000),
accountID integer NOT NULL,
CONSTRAINT AccountTest_PK PRIMARY KEY (ID),
CONSTRAINT AccountTest_Account_FK FOREIGN KEY (accountID)
    REFERENCES account (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
OIDS=FALSE
);
ALTER TABLE AccountTest
OWNER TO globeAdmin;

-- Table: AccountProject

-- DROP TABLE AccountProject;

create TABLE AccountProject
(
ID serial,
Name character varying(100) not null,
OccupationID integer not null,
Month integer NOT NULL,
Year integer NOT NULL,
URL character varying(500),
Description character varying(2000),
accountID integer NOT NULL,
CONSTRAINT AccountProject_PK PRIMARY KEY (ID),
CONSTRAINT AccountProject_Account_FK FOREIGN KEY (accountID)
    REFERENCES account (id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
OIDS=FALSE
);
ALTER TABLE AccountProject
OWNER TO globeAdmin;

-- Table: image

-- DROP TABLE image;

CREATE TABLE image (

  id serial,
  category character varying(5) NOT NULL,
  type character varying(20) NOT NULL,
  parentID integer,
  image bytea NOT NULL,
  CONSTRAINT image_pk PRIMARY KEY (id)
);
ALTER TABLE image
OWNER TO globeAdmin;
