/*
 * 096-getWorkflowsByTeamId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getWorkflowsByTeamId(workflowTeamId uuid,
                                                pageLimit integer,
                                                pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    execute boundsCheck(pageLimit, 1, 1000);
    execute boundsCheck(pageOffset, 0, 1000);
    select jsonb_agg(data)
    into result
    from (select jsonb_build_object(
                         'id', id,
                         'name', name,
                         'iconId', iconId,
                         'ownerId', ownerId,
                         'teamId', teamId,
                         'owner', owner,
                         'team', team,
                         'everyone', everyone,
                         'description', description
                     ) as data
          from workflows
          where teamId = workflowTeamId
          limit pageLimit offset pageOffset) as subquery;
    return coalesce(result, '[]'::jsonb);
end ;
$$ language plpgsql;
