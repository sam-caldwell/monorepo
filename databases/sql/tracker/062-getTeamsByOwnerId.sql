/*
 * 062-getTeamsByOwnerId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getTeamsByOwnerId(teamOwnerId uuid,
                                             pageLimit integer default 1000,
                                             pageOffset integer default 0) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(data)
    into result
    from (select jsonb_build_object(
                         'id', id,
                         'name', name,
                         'iconId', iconId,
                         'ownerId', ownerId,
                         'owner', owner,
                         'team', team,
                         'everyone', everyone,
                         'description', description
                     ) as data
          from teams
          where teamOwnerId = ownerId
          limit pageLimit offset pageOffset) as subquery;
    return result;
end;
$$ language plpgsql;
