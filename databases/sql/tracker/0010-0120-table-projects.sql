/*
 * 0010-0400-table-projects.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 */
create table if not exists projects
(
    id          uuid primary key not null,
    -- each workflow has a unique name --
    name        varchar(64) not null,
    -- a uuid representing the icon for the workflow. --
    iconId        uuid not null,
    -- a project will have an owner and team which have permissions --
    ownerId     uuid        not null,
    teamId      uuid        not null,
    owner       permissions not null default 'delete',
    team        permissions not null default 'none',
    everyone    permissions not null default 'none',
    defaultTicketType uuid not null,
    -- --
    created  timestamp        not null default now(),
    -- descriptive text --
    description text,
    foreign key (ownerId) references users (id),
    foreign key (teamId) references teams (id),
    foreign key (iconId) references icons(id),
    foreign key (defaultTicketType) references ticketTypes(id),
    foreign key (id) references entity (id) on delete restrict
);

create unique index if not exists ndxProjectsName on projects (name);
create index if not exists ndxProjectsCreated on projects (created);
create index if not exists ndxProjectsDefaultTicketType on projects (defaultTicketType);
