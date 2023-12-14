/*
 * 0010-0500-table-workflowSteps.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create user profiles for the tracker.
 */
create table if not exists workflowSteps
(
    id          uuid primary key not null,
    workflowId  uuid        not null,
    name        varchar(64) not null,
    -- a uuid representing the icon for the workflow.
    iconId      uuid        not null,
    prevStepId  uuid        not null,
    nextStepId  uuid        not null,
    -- --
    created  timestamp        not null default now(),
    description text,
    foreign key (workflowId) references workflow (id),
    foreign key (prevStepId) references workflowSteps (id),
    foreign key (nextStepId) references workflowSteps (id),
    foreign key (iconId) references icons (id),
    foreign key (id) references entity (id) on delete restrict
);

create unique index if not exists ndxWorkflowStepsName on workflowSteps (name);
create index if not exists ndxPWorkflowStepsIconId on workflowSteps (iconId);
create index if not exists ndxPWorkflowStepsPrevStepId on workflowSteps (prevStepId);
create index if not exists ndxPWorkflowStepsNextStepId on workflowSteps (nextStepId);
create index if not exists ndxWorkflowCreated on workflowSteps (created);
