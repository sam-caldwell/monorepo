/*
 * 090-workflows.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * The workflow record is the top-level object representing a workflow for a given ticket.
 * A record in workflowsteps table links to this representing a list of steps in the workflow,
 * starting with a 'start' node and ending with a 'terminate' node.
 */
create table if not exists workflows
(
    id          uuid primary key not null, -- workflowsId uniquely identifies the workflows --
    name        varchar(64)      not null, -- each workflows has a unique name --
    -- --
    iconId      uuid             not null, -- a uuid representing the icon for the workflows. --
    ownerId     uuid             not null, -- a workflows will have an owner and team which have permissions --
    teamId      uuid             not null, -- permissions are granted to the owner, team and everyone --
    -- --
    owner       permissions      not null default 'delete',
    team        permissions      not null default 'read',
    everyone    permissions      not null default 'read',
    -- --
    created     timestamp        not null default now(),
    -- --
    description text,-- --
    constraint validateWorkflowName check (validName(name)),
    foreign key (ownerId) references users (id),
    foreign key (teamId) references teams (id),
    foreign key (iconId) references icons (id),
    foreign key (id) references entity (id) on delete restrict
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxWorkflowsName on workflows (name);
create index if not exists ndxWorkflowsCreated on workflows (created);
create index if not exists ndxWorkflowsOwner on workflows (owner);
create index if not exists ndxWorkflowsTeam on workflows (team);
create index if not exists ndxWorkflowsEveryone on workflows (everyone);


