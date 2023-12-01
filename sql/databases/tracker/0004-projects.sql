create table if not exists projects
(
    id          uuid primary key default gen_random_uuid(),
    -- each workflow has a unique name --
    name        varchar(64) not null unique,
    -- a uuid representing the icon for the workflow. --
    iconId        uuid not null,
    -- a project will have an owner and team which have permissions --
    ownerId     uuid        not null,
    teamId      uuid        not null,
    owner       permissions not null default 'delete',
    team        permissions not null default 'none',
    everyone    permissions not null default 'none',
    defaultTicketType uuid not null,
    -- a project is mapped to a workflow --
    workflowId  uuid        not null,
    -- descriptive text --
    description text,
    foreign key (ownerId) references users (id),
    foreign key (teamId) references teams (id),
    foreign key (workflowId) references workflow (id),
    foreign key (iconId) references icons(id),
    foreign key (defaultTicketType) references ticketTypes(id)
);
