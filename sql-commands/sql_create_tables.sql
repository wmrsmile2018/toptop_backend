\! tput setaf 1; "___________________________create table users ____________________________";
\! tput setaf 2;

create table users (
  id      varchar primary key,
  phone   integer unique
);

\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table places ____________________________";
\! tput setaf 2;

create table places (
  id                varchar primary key,
  id_user           varchar,
  id_category       varchar,
  full_description  varchar,
  short_description varchar,
  creation_date     timestamp,
  address           varchar,
  label             varchar,
  inn               varchar,
  priority          varchar
);

\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table category ____________________________";
\! tput setaf 2;

create table category (
  id    varchar primary key,
  name  varchar unique
);

\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table user_tags ____________________________";
\! tput setaf 2;

create table user_tags (
  id    varchar primary key,
  name  varchar unique
);

\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table place_user_tags ____________________________";
\! tput setaf 2;

create table place_user_tags (
  id_user_tags  varchar ,
  id_place       varchar
);

\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table photos ____________________________";
\! tput setaf 2;

create table photos (
  id      varchar primary key,
  id_place varchar,
  name    varchar unique
);

\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table filters ____________________________";
\! tput setaf 2;

create table filters (
  id    varchar primary key,
  name  varchar
);


\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table filters_tags ____________________________";
\! tput setaf 2;

create table place_filters (
  id_place  varchar,
  id_filter   varchar
);

