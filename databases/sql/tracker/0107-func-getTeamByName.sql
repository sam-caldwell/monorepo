/*
 * 0107-func-getTeamByName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getTeamByName(teamName varchar(64)) returns jsonb as
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
    where name = teamName;
    return result;
end;
$$ language plpgsql;
