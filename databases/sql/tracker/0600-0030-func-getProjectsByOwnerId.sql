/*
 * 0020-0030-func-getProjectsByOwnerId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function getProjectsByOwnerId(OwnerId uuid) returns jsonb as
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
    where ownerId == OwnerId;
    return result;

end ;
$$ language plpgsql;
