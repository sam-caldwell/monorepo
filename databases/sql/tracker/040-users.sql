/*
 * 040-users.sql
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
    foreign key (avatarId) references avatars (id) on delete restrict,
    foreign key (id) references entity (id) on delete restrict
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
 * createUser() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createUser(firstName varchar(64), lastName varchar(64), avatarId uuid, email varchar(256),
                                      phoneNumber varchar(20), description text) returns uuid as
$$
declare
    userId uuid;
begin
    userId := (select createEntity('avatar'::entityType));
    insert into users (id, firstName, lastName, avatarId, email, phoneNumber, description)
    values (userId, firstName, lastName, avatarId, email, phoneNumber, description);
    return userId;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteUsersById() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteUsersById(userId uuid) returns integer as
$$
declare
    count integer;
begin
    --Note: a user should not be deletable if it is associated with any
    --      project, ticket, workflow, team (as owner)
    delete from users where id = userId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getUserById() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getUserById(userId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
                   'id', id,
                   'firstName', firstName,
                   'lastName', lastName,
                   'avatarId', avatarId,
                   'email', email,
                   'phoneNumber', phoneNumber,
                   'description', description
               ) as user
    into result
    from users
    where id = userId
    limit 1;
    return result;

end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getUserByEmail() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getUserByEmail(emailAddress varchar(256)) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
                   'id', id,
                   'firstName', firstName,
                   'lastName', lastName,
                   'avatarId', avatarId,
                   'email', email,
                   'phoneNumber', phoneNumber,
                   'description', description
               ) as user
    into result
    from users
    where email = emailAddress
    limit 1;
    return result;

end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getUserByPhone() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getUserByPhone(phone varchar(20)) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
                   'id', id,
                   'firstName', firstName,
                   'lastName', lastName,
                   'avatarId', avatarId,
                   'email', email,
                   'phoneNumber', phoneNumber,
                   'description', description
               ) as user
    into result
    from users
    where phoneNumber = phone
    limit 1;
    return result;

end ;
$$ language plpgsql;
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
