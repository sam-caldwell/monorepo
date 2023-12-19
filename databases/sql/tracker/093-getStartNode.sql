/*
 * 093-getStartNode.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * When we create a workflow, there must be a start and terminal node
 * representing the first and last steps in the workflow.
 *
 * - Get the start node (default) for a given workflow.
 */
create or replace function getStartNode(thisWorkflow uuid) returns uuid as
$$
declare
    entityId uuid := (select id
                      from workflowSteps
                      where workflowId = thisWorkflow
                        and name = 'terminate');
begin
    if entityId = '00000000-0000-0000-0000-000000000000' then
        raise exception 'internal error. missing start node for this workflow';
    end if;
    return entityId;
end ;
$$ language plpgsql;
