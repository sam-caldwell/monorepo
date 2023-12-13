/*
 * 0250-0200-func-getUsersInTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Return all userIds in the given team
 */
create or replace function getUsersInTeam(id uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'userId', userId
        )) as workflow
    into result
    from teamMembership
    where teamId = id;
    return result;
end;
$$ language plpgsql;
