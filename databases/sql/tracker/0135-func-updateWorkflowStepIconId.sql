/*
 * 0135-func-updateWorkflowStepIconId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateWorkflowStepDescription(stepId uuid, iconId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflowSteps set iconId=iconId where id=stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
