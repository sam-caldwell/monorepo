/*
 * 096-updateWorkflowsProperties.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsDescription()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsDescription(workflowId uuid, descriptionText text) returns integer as
$$
declare
    count integer;
begin
    update workflows set description=descriptionText where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsIconId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsIconId(workflowId uuid, newId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflows set iconId=newId where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsName()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsName(workflowId uuid, workflowName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update workflows set name=workflowName where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsOwnerId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsOwnerId(workflowId uuid, newId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflows set ownerId=newId where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsPermEveryone()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsPermEveryone(workflowId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflows set everyone=newPermission where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsPermOwner()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsPermOwner(workflowId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflows set owner=newPermission where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsPermTeam()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsPermTeam(workflowId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflows set team=newPermission where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsTeamId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsTeamId(workflowId uuid, newId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflows set teamId=newId where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
