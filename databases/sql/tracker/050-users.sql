/*
 * 050-users.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create user profiles for the tracker.
 */
create table if not exists users
(
    -- the userId used to uniquely identify the users --
    id          uuid primary key not null,
    -- the user's first and last names --
    firstName   varchar(64)      not null,
    lastName    varchar(64)      not null,
    avatarId    uuid             not null,
    created     timestamp        not null default now(),
    -- the user's identifying email address (globally unique) --
    email       varchar(256)     not null,
    -- an optional phone number --
    phoneNumber varchar(20)      null,
    -- descriptive text --
    description text             null,
    foreign key (avatarId) references avatars (id) on delete cascade,
    foreign key (id) references entity (id) on delete restrict,
    constraint validateFirstName check (validFirstLastName(firstName)),
    constraint validateLastName check (validFirstLastName(lastName)),
    constraint validateEmail check (validEmailAddress(email)),
    constraint validatePhone check (validPhoneNumber(phoneNumber))
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create index if not exists ndxUsersCreated on users (created);
create unique index if not exists ndxUsersEmail on users (email);
create index if not exists ndxUsersFirstName on users (firstName);
create index if not exists ndxUsersLastName on users (lastName);
create unique index if not exists ndxUsersPhoneNumber on users (phoneNumber);
