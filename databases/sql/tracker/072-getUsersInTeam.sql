/*
 * 072-getUsersInTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getUsersInTeam(thisTeamId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'userId', userId
        )) as workflow
    into result
    from teamMemberships
    where teamId = thisTeamId;
    return result;
end;
$$ language plpgsql;
