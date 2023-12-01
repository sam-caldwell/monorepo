/*
 * 0120-func-getUserById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function getUserById(userId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'firstName', firstName,
            'lastName', lastName,
            'avatarId', avatarId,
            'email', email,
            'phoneNumber', phoneNumber,
            'description', description
        ))  as workflow into result
    from users
    where id == userId;
    return result;

end ;
$$ language plpgsql;
