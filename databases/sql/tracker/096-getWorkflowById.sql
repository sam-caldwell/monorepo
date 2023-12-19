/*
 * 096-getWorkflowById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getWorkflowById(thisWorkflowId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
                   'id', id,
                   'name', name,
                   'iconId', iconId,
                   'ownerId', ownerId,
                   'teamId', teamId,
                   'owner', owner,
                   'team', team,
                   'everyone', everyone,
                   'description', description
               ) as workflows
    into result
    from workflows
    where id = thisWorkflowId
    limit 1;
    return result;
end;
$$ language plpgsql;
