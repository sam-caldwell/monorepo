/*
 * 096-getWorkflowsByOwnerId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getWorkflowsByOwnerId(workflowOwnerId uuid,
                                                 pageLimit integer,
                                                 pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    execute boundsCheck(pageLimit, 1, 1000);
    execute boundsCheck(pageOffset, 0, 1000);
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'name', name,
            'iconId', iconId,
            'ownerId', ownerId,
            'teamId', teamId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'description', description
        )) as workflows
    into result
    from workflows
    where ownerId = workflowOwnerId
    limit pageLimit offset pageOffset;
    return result;
end ;
$$ language plpgsql;


