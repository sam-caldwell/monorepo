/*
 * 052-getUsersByEmail.sql
 * (c) 2023 Sam Caldwell. See License.txt
 */
create or replace function getUsersByEmail(emailAddress varchar(256)) returns jsonb as
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
