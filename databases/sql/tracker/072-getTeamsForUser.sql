/*
 * 072-getTeamsForUser.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getTeamsForUser(thisUserId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'teamId', teamId
        )) as team
    into result
    from teamMemberships
    where userId = thisUserId;
    return result;
end;
$$ language plpgsql;
