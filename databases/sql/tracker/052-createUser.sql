/*
 * 052-createUser.sql
 * (c) 2023 Sam Caldwell. See License.txt
 */
create or replace function createUser(newFirstName varchar(64), newLastName varchar(64), newAvatarId uuid,
                                      newEmailAddress varchar(256), newPhoneNumber varchar(20),
                                      newDescription text default ''::text) returns uuid as
$$
declare
    entityId uuid;
begin

    entityId := (select createEntity('user'::entityType));

    if not validName(newFirstName) then
        raise exception 'invalid firstname';
    end if;

    if not validName(newLastName) then
        raise exception 'invalid lastname';
    end if;

    if not validemailaddress(newEmailAddress) then
        raise exception 'invalid email address';
    end if;

    if not validphonenumber(newPhoneNumber) then
        raise exception 'invalid phone number';
    end if;

    insert into users (id, firstName, lastName, avatarId, email, phoneNumber, description)
    values (entityId, newFirstName, newLastName, newAvatarId, newEmailAddress, newPhoneNumber, newDescription);

    return entityId;
end;
$$ language plpgsql;
