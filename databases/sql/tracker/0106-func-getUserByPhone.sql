/*
 * 0106-func-getUserByPhone.sql
 * (c) 2023 Sam Caldwell.  See License.txt
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
        )  as workflow into result
    from users
    where phoneNumber = phone
    limit 1;
    return result;

end ;
$$ language plpgsql;
