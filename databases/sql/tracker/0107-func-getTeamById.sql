/*
 * 0107-func-getTeamById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getTeamById(teamId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
            'id', id,
            'name', name,
            'iconId', iconId,
            'ownerId', ownerId,
            'perm_owner', owner,
            'perm_team', team,
            'perm_everyone', everyone,
            'description', description
        ) as team
    into result
    from teams
    where id = teamId;
    return result;
end;
$$ language plpgsql;
