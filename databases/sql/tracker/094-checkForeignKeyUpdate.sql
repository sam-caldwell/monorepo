/*
 * 094-checkForeignKeyUpdate.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Trigger on insert/update to ensure prevStepId and nextStepId are in workflowSteps
 * unless they are null.  null prev/next are allowed.
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace trigger checkForeignKeyUpdate
    before update
    on workflowSteps
    for each row
execute function workflowNodeValidation();
