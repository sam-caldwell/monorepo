/*
 * 0002-workflow.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This is the top-level table for representing a workflow.
 * This workflow is then composed of records in workflowSteps.
 */
create table if not exists workflow
(
    -- workflowId uniquely identifies the workflow --
    id          uuid primary key     default gen_random_uuid(),
    -- each workflow has a unique name --
    name        varchar(64) not null unique,
    -- a uuid representing the icon for the workflow.
    iconId      uuid        not null,
    -- a workflow will have an owner and team which have permissions --
    ownerId     uuid        not null,
    teamId      uuid        not null,
    -- permissions are granted to the owner, team and everyone --
    owner       permissions not null default 'delete',
    team        permissions not null default 'read',
    everyone    permissions not null default 'read',
    -- descriptive text --
    description text,
    foreign key (ownerId) references users (id),
    foreign key (teamId) references teams (id),
    foreign key (iconId) references icons (id)
);
