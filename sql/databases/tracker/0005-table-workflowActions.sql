/*
 * 0005-table-workflowActions.sql
 * (c) 2023 Sam Caldwell.  See License.txt.
 */
create table if not exists workflowActions
(
    id             uuid primary key default gen_random_uuid(),
    workflowStepId uuid          not null,
    name           varchar(64)   not null unique,
    topic          varchar(2048) not null, -- the topic to be published to the message processor
    message        varchar(2048) not null, -- the message to send (formatting string)
    description    text,
    foreign key (workflowStepId) references workflowSteps (id)
);
