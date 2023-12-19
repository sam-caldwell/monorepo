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
    prevStepId  uuid             null,
    nextStepId  uuid             null,
    -- --
    created     timestamp        not null default now(),
    description text,
    foreign key (workflowId) references workflows (id),
    /*
     * see 093-verifyWorkFlowPrevNextStepsValid.sql
     *
     * We have two triggers (below) which enforce referential integrity
     * to ensure prevStepId and nextStepId are in the set of workflowSteps.id
     * if non-null uuid, but taking no action if the prevStepId or nextStepId
     * are null.
     */
    foreign key (id) references entity (id) on delete restrict
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxWorkflowStepsWorkflowIdName on workflowSteps (workflowId, name);
create index if not exists ndxWorkflowStepsCreated on workflowSteps (created);
