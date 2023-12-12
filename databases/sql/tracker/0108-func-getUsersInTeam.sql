/*
 * 0107-func-getUsersInTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Return all userIds in the given team
 */
create or replace function getUsersInTeam(teamId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'userId', userId
        )) as workflow
    into result
    from teamMembership
    where teamId == teamId;
    return result;
end;
$$ language plpgsql;
