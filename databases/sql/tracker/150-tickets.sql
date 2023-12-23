/*
 * 150-tickets.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function isTicketTypeidSupported(thisProject uuid, thisType uuid) returns boolean as
$$
begin
    return (select count(*)
            from projecttickettypes
            where projectid = thisProject
              and thisType in (select ticketTypeId from projecttickettypes)) > 0;
end;
$$ language plpgsql;

create table if not exists tickets
(
    id             uuid primary key not null,
    /*
     * Project id allows tickets to be grouped into projects
     */
    projectId      uuid             not null,
    /*
     * A ticket must have a classifier (ticket type) which will determine which workflow is applied.
     */
    ticketTypeId   uuid             not null,
    /*
     * For a given ticket type, there is a workflow.  But the ticket must also know its current
     * step int that workflow which is tracked separately.
     */
    workflowStepId uuid             not null,
    /*
     * We must know who created the ticket and to whom the ticket is currently assigned.
     */
    authorId       uuid             not null,
    assignedId     uuid             not null,
    /*
     * permissions are granted to the owner, team and everyone --
     */
    owner          permissions      not null default 'delete',
    team           permissions      not null default 'read',
    everyone       permissions      not null default 'read',
    /*
     * descriptive text
     */
    summary        varchar(2048)    not null,
    description    text,
    /*
     * constraints
     */
    foreign key (id) references entity (id) on delete restrict,
    foreign key (projectId) references projects (id),
    foreign key (authorId) references users (id),
    foreign key (assignedId) references users (id),
    constraint ensureTicketTypeSupported check (isTicketTypeidSupported(projectId, ticketTypeId)),
    foreign key (workflowStepId) references workflowSteps (id)
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create index if not exists ndxTicketSummary on tickets (summary);
create index if not exists ndxTicketOwner on tickets (owner);
create index if not exists ndxTicketTeam on tickets (team);
create index if not exists ndxTicketEveryone on tickets (everyone);
