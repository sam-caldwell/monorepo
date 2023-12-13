/*
 * 0010-0000-table-ticketTypes.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Identify the ticket classifications.
 */
create table if not exists ticketTypes
(
    id          uuid primary key default gen_random_uuid(),
    -- each workflow has a unique name --
    name        varchar(64) not null unique,
    -- a uuid representing the icon for the workflow.
    iconId      uuid        not null,
    -- a project ticketType is mapped to a workflow --
    workflowId  uuid        not null,
    -- descriptive text --
    description text,
    foreign key (iconId) references icons (id),
    foreign key (workflowId) references workflow (id)
);
