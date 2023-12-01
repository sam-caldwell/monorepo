/*
 * 0004-workflowSteps.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create user profiles for the tracker.
 */
create table if not exists workflowSteps
(
    id          uuid primary key default gen_random_uuid(),
    workflowId  uuid        not null,
    name        varchar(64) not null,
    -- a uuid representing the icon for the workflow.
    iconId      uuid        not null,
    prevStepId  uuid        not null,
    nextStepId  uuid        not null,
    description text,
    foreign key (workflowId) references workflow (id),
    foreign key (nextStepId) references workflowSteps (id),
    foreign key (nextStepId) references workflowSteps (id),
    foreign key (iconId) references icons (id)
);
