/*
 * 052-createUser.sql
 * (c) 2023 Sam Caldwell. See License.txt
 */
create or replace function createUser(f varchar(64), l varchar(64), a uuid, e varchar(256),
                                      p varchar(20), d text) returns uuid as
$$
declare
    entityId uuid;
begin
    entityId := (select createEntity('user'::entityType));
    insert into users (id, firstName, lastName, avatarId, email, phoneNumber, description)
    values (entityId, f, l, a, e, p, d);
    return entityId;
end;
$$ language plpgsql;
