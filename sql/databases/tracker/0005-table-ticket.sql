/*
 * 0005-table-ticket.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create table if not exists ticket
(
    id             uuid primary key       default gen_random_uuid(),
    projectId      uuid          not null,
    authorUserId   uuid          not null,
    assignedUserId uuid          not null,
    ticketTypeId   uuid          not null,
    -- permissions are granted to the owner, team and everyone --
    owner          permissions   not null default 'delete',
    team           permissions   not null default 'read',
    everyone       permissions   not null default 'read',
    -- descriptive text --
    subject        varchar(2048) not null,
    workflowStepId uuid          not null,
    description    text,
    foreign key (projectId) references projects (id),
    foreign key (authorUserId) references users (id),
    foreign key (assignedUserId) references users (id),
    foreign key (ticketTypeId) references ticketTypes (id),
    foreign key (workflowStepId) references workflowSteps(id)
);
