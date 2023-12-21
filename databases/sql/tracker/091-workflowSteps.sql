/*
 * 091-workflowSteps.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create table if not exists workflowSteps
(
    id          uuid primary key not null,
    workflowId  uuid             not null,
    name        varchar(64)      not null,
    -- a uuid representing the icon for the workflow.
    prevStepId  uuid             null, -- see checkForeignKeyInsert/Update (must be in workflowSteps.id)
    nextStepId  uuid             null, -- see checkForeignKeyInsert/Update (must be in workflowSteps.id)
    -- --
    created     timestamp        not null default now(),
    action      uuid             null, --if action is mapped, the action uuid will be stored here.
    description text,
    foreign key (workflowId) references workflows (id) on delete cascade,
    --we add an fk in 140+ after workflowactions exists...
    /*
     * see 093-workflowNodeValidation.sql
     *
     * We have two triggers (below) which enforce referential integrity
     * to ensure prevStepId and nextStepId are in the set of workflowSteps.id
     * if non-null uuid, but taking no action if the prevStepId or nextStepId
     * are null.
     */
    foreign key (id) references entity (id) on delete restrict,
    constraint validateWorkflowStepName check (validName(name))
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxWorkflowStepsWorkflowIdName on workflowSteps (workflowId, name);
create index if not exists ndxWorkflowStepsCreated on workflowSteps (created);
