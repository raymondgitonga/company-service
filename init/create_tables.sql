create table person
(
    personId serial PRIMARY KEY,
    email   varchar(50) unique not null,
    role varchar(50) unique not null
);

create table company
(
    companyId serial PRIMARY KEY,
    name       varchar(50) not null,
    code       int not null,
    country    varchar(50) not null,
    website    varchar(50) not null,
    phone      varchar(15) not null
);

insert into person (email, role)
values ('raymond@test.com', 'admin'),
       ('gitonga@test.com', 'user');

insert into company (name, code, country, website, phone)
values ('Safaricom', 122, 'KE', 'www.saf.com', '7220000000')
