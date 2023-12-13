/*
 * 0300-0040-func-updateWorkflowDescription.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateWorkflowDescription(workflowId uuid, description text) returns integer as
$$
declare
    count integer;
begin
    update workflow set odescriptionwnerId=descriptionText where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
