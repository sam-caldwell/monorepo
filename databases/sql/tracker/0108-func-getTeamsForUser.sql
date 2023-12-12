/*
 * 0107-func-getTeamsForUser.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Given a userId, return all matching teams
 */
create or replace function getTeamsForUser(userId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
            'teamId', teamId
        ) as team
    into result
    from teamMembership
    where userId = userId
    limit 1;
    return result;
end;
$$ language plpgsql;
