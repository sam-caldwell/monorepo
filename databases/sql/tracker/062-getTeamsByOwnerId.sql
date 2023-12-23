/*
 * 062-getTeamsByOwnerId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getTeamsByOwnerId(teamOwnerId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
                   'id', id,
                   'name', name,
                   'iconId', iconId,
                   'ownerId', ownerId,
                   'owner', owner,
                   'team', team,
                   'everyone', everyone,
                   'description', description
               ) as team
    into result
    from teams
    where teamOwnerId = ownerId
    limit 1;
    return result;
end;
$$ language plpgsql;
