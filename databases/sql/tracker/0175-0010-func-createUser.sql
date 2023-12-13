/*
 * 0175-0010-func-createUser.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function createUser(firstName varchar(64), lastName varchar(64), avatarId uuid, email varchar(256),
                                      phoneNumber varchar(20), description text) returns uuid as
$$
declare
    userId uuid;
begin
    userId := gen_random_uuid();
    insert into users (id, firstName, lastName, avatarId, email, phoneNumber, description)
    values (userId, firstName, lastName, avatarId, email, phoneNumber, description);
    return userId;
end;
$$ language plpgsql;
