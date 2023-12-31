/*
 * 120-projectsTicketTypes.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Associates a project with its supported ticket types.
 * For example, an agile project may have epics, issues, and bugs.
 * An incident project may have incident and postmortem tickets.
 */

create table if not exists projectTicketTypes
(
    id           uuid primary key not null,
    -- associate a project to one or more ticket types --
    projectId    uuid             not null,
    ticketTypeId uuid             not null,
    -- we need to know when this is created --
    created      timestamp        not null default now(),
    foreign key (projectId) references projects (id),
    foreign key (ticketTypeId) references ticketTypes (id)
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
-- no indexes needed
