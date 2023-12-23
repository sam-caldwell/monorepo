/*
 * 052-getUsersByEmail.sql
 * (c) 2023 Sam Caldwell. See License.txt
 */
create or replace function getUsersByEmail(emailAddress varchar(256),
                                           pageLimit integer,
                                           pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(data)
    into result
    from (select jsonb_build_object(
                         'id', id,
                         'firstName', firstName,
                         'lastName', lastName,
                         'avatarId', avatarId,
                         'email', email,
                         'phoneNumber', phoneNumber,
                         'description', description
                     ) as data
          from users
          where email = emailAddress
          limit pageLimit offset pageOffset) as subquery;
    return result;
end ;
$$ language plpgsql;
