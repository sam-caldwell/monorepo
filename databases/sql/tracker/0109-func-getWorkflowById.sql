/*
 * 0120-func-getWorkflowById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function getWorkflowById(workflowId uuid) returns jsonb as
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
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'description', description
        )) as workflow
    into result
    from workflow
    where id == workflowId;
    return result;
end;
$$ language plpgsql;