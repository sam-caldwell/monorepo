/*
 * 052-getUsersByPhone.sql
 * (c) 2023 Sam Caldwell. See License.txt
 */
create or replace function getUsersByPhone(thisPhone varchar(20),
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
                         'phoneNumber', thisPhone,
                         'description', description
                     ) as data
          from users
          where phoneNumber = thisPhone
          limit pageLimit offset pageOffset) as subquery;
    return coalesce(result, '[]'::jsonb);
end;
$$ language plpgsql;
