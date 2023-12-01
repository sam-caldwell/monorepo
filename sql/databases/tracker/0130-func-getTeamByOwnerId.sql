/*
 * 0130-func-getTeamByName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getTeamByOwnerId(teamOwnerId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'name', name,
            'iconId', iconId,
            'ownerId', ownerId,
            'perm_owner', owner,
            'perm_team', team,
            'perm_everyone', everyone,
            'description', description
        ))
    into result
    from teams
    where teamOwnerId == ownerId;
    return result;
end;
$$ language plpgsql;
