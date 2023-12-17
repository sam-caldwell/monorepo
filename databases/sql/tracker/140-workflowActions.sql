/*
 * 140-workflowActions.sql
 * (c) 2023 Sam Caldwell.  See License.txt.
 */
create table if not exists workflowActions
(
    id             uuid primary key not null,
    StepId uuid          not null,
    name           varchar(64)   not null unique,
    topic          varchar(2048) not null, -- the topic to be published to the message processor
    message        varchar(2048) not null, -- the message to send (formatting string)
    -- --
    created  timestamp        not null default now(),
    description    text,
    foreign key (StepId) references workflowSteps (id),
    foreign key (id) references entity (id) on delete restrict
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxWorkflowActionsStepIdName on workflowActions (StepId, name);
create index if not exists ndxWorkflowActionsTopic on workflowActions (topic);
create index if not exists ndxWorkflowActionsCreated on workflowActions (created);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createTicketType()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
