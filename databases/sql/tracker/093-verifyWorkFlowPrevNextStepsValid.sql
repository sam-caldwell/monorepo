/*
 * 093-verifyWorkFlowPrevNextStepsValid.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Trigger on insert/update to ensure prevStepId and nextStepId are in workflowSteps
 * unless they are null.  null prev/next are allowed.
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function verifyWorkFlowPrevNextStepsValid()
    returns trigger as
$$
declare
    prevNotInSet boolean:=(select count(id) from workflowSteps where new.prevStepId in (select id from workflowSteps)) = 0;
    nextNotInSet boolean:=(select count(id) from workflowSteps where new.nextStepId in (select id from workflowSteps)) = 0;
begin
    if (new.prevStepId is not null) and prevNotInSet then
        RAISE EXCEPTION 'constraint violation: prevStepId must exist in workflowSteps.id';
    end if;
    IF (new.nextStepId is not null) and nextNotInSet then
        RAISE EXCEPTION 'constraint violation: nextStepId must exist in workflowSteps.id';
    end if;
    return new;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * check_foreign_key_insert
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create trigger check_foreign_key_insert
    before insert
    on workflowSteps
    for each row
execute function verifyWorkFlowPrevNextStepsValid();
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * check_foreign_key_update
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create trigger check_foreign_key_update
    before update
    on workflowSteps
    for each row
execute function verifyWorkFlowPrevNextStepsValid();
