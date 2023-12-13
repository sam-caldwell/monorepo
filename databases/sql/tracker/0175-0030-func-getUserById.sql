/*
 * 0004-0030-func-getUserById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
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
        )  as user into result
    from users
    where id = userId
    limit 1;
    return result;

end ;
$$ language plpgsql;
