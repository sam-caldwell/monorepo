/*
 * 0115-func-getProjectsByTeamId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function getProjectsByTeamId(TeamId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'name', name,
            'iconId', iconId,
            'ownerId', ownerId,
            'teamId', teamId,
            'permissionOwner', owner,
            'permissionTeam', team,
            'permissionEveryone', everyone,
            'defaultTicketType', defaultTicketType,
            'description', description
        ))  as data into result
    from projects
    where teamId == TeamId;
    return result;

end ;
$$ language plpgsql;
