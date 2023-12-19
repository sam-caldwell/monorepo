/*
 * 096-createWorkflow.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 *  - create a workflow object.
 *  - initialize the workflowSteps table with a start and terminate node.
 */
create or replace function createWorkflow(workflowName varchar(64), workflowIconId uuid, workflowOwnerId uuid,
                                          workflowTeamId uuid, workflowPermissionOwner permissions,
                                          workflowPermissionTeam permissions, workflowPermissionEveryone permissions,
                                          workflowPermissionDescription text) returns uuid as
$$
declare
    workflowId uuid := (select createEntity('workflow'::entityType));
begin
    insert into workflows (id, name, iconId, ownerId, teamId, owner, team, everyone, description)
    values (workflowId, workflowName, workflowIconId, workflowOwnerId, workflowTeamId,
            workflowPermissionOwner, workflowPermissionTeam, workflowPermissionEveryone,
            workflowPermissionDescription);

    /*
     *  For every new workflow, there must be a start and stop workflow step
     */
    if initializeWorkflowSteps(workflowId) != 1 then
        raise exception 'initialization failed';
    end if;
    return workflowId;
end;
$$ language plpgsql;
