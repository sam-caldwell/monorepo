/*
 * 0300-0030-func-getWorkflowsByOwnerId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function getWorkflowsByOwnerId(workflowOwnerId uuid,
                                                pageLimit integer,
                                                pageOffset integer) returns jsonb as
$$
declare
    discard bool;
    result jsonb;
begin
    discard:=(select boundsCheck(pageLimit,1,1000));

    discard:=(select boundsCheck(pageOffset,0,1000));

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
        )) as workflow
    into result
    from workflow
    where ownerId = workflowOwnerId
    limit pageLimit offset pageOffset;
    return result;
end ;
$$ language plpgsql;
