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
    description text,
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
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateUserDescription() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateUserDescription(userId uuid, newDescription text) returns integer as
$$
declare
    count integer;
begin
    update users set description=newDescription where id = userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateUserEmail() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateUserEmail(userId uuid, newEmail varchar(256)) returns integer as
$$
declare
    count integer;
begin
    update users set email=newEmail where id = userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateUserName() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateUserName(userId uuid, newFirstName varchar(64), newLastName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update users set firstName=newFirstName, lastName=newLastName where id = userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateUserPhone() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateUserPhone(userId uuid, newPhoneNumber varchar(20)) returns integer as
$$
declare
    count integer;
begin
    update users set phoneNumber=newPhoneNumber where id = userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateUserAvatar() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateUserAvatar(userId uuid, newAvatarId uuid) returns integer as
$$
declare
    count integer;
begin
    update users set avatarId=newAvatarId where id = userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;

