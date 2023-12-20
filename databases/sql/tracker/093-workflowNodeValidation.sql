/*
 * 093-workflowNodeValidation.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Trigger on insert/update to ensure prevStepId and nextStepId are in workflowSteps
 * unless they are null.  null prev/next are allowed.
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function workflowNodeLinkValidation()
    returns trigger as
$$
begin
    if (new.name = 'start') then
        if (new.prevstepid is not null) then
            raise exception 'constraint violation: start nodes must have null prevStepId';
        end if;
        if (new.nextstepid is null) then
            raise exception 'constraint violation: start nodes cannot have null prevStepId';
        end if;
    end if;
    if (new.name = 'terminate') then
        if (new.prevstepid is null) then
            raise exception 'constraint violation: terminate nodes cannot have null nextStepId';
        end if;
        if (new.nextstepid is not null) then
            raise exception 'constraint violation: terminate nodes must have null nextStepId';
        end if;
    end if;
    if (new.nextstepId is null) and (new.name != 'terminate') then
        raise exception 'constraint violation: only terminate nodes may have null nextStepId';
    end if;
    if (select count(id) from workflowSteps where new.prevStepId in (select id from workflowSteps)) = 0 then
        raise exception 'constraint violation: prevStepId must exist in workflowSteps.id';
    end if;
    if (select count(id) from workflowSteps where new.nextstepId in (select id from workflowSteps)) = 0 then
        raise exception 'constraint violation: nextstepId must exist in workflowSteps.id';
    end if;
    return new;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * check_foreign_key_insert
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create trigger checkForeignKeyInsert
    before insert
    on workflowSteps
    for each row
execute function workflowNodeLinkValidation();
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * check_foreign_key_update
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create trigger checkForeignKeyUpdate
    before update
    on workflowSteps
    for each row
execute function workflowNodeLinkValidation();
