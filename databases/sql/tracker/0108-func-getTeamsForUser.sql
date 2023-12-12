/*
 * 0108-func-getTeamsForUser.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Given a userId, return all matching teams
 */
create or replace function getTeamsForUser(id uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'teamId', teamId
        )) as team
    into result
    from teamMembership
    where userId = id;
    return result;
end;
$$ language plpgsql;
