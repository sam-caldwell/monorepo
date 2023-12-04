/*
 * 0115-func-getProjectsById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function getProjectsById(projectId uuid) returns jsonb as
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
    where id == projectId;
    return result;

end ;
$$ language plpgsql;
