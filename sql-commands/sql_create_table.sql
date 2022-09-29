\! tput setaf 1; "___________________________create table users ____________________________";
\! tput setaf 2;

create table users (
  id      varchar primary key,
  email   integer unique
);

\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table posts ____________________________";
\! tput setaf 2;

create table posts (
  id                varchar primary key,
  id_user           varchar,
  id_category       varchar,
  full_description  varchar,
  short_description varchar,
  creation_date     timestamp,
  address           varchar,
  label             varchar
);

\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table category ____________________________";
\! tput setaf 2;

create table category (
  id    varchar primary key,
  name  varchar unique
);

\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table tags ____________________________";
\! tput setaf 2;

create table tags (
  id    varchar primary key,
  name  varchar unique
);

\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table photos ____________________________";
\! tput setaf 2;

create table photos (
  id      varchar primary key,
  id_post varchar,
  name    varchar unique
);

\! tput setaf 3; "____________________________________________________________________________________";
\! tput setaf 1; "___________________________create table post_tags ____________________________";
\! tput setaf 2;

create table post_tags (
  id_tag   varchar,
  id_post  varchar
);
