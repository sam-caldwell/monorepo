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
    -- Ticket identifier (id) is the publicly identifiable entityId used to uniquely identify a ticket.
    id             uuid primary key not null,
    -- Project id allows tickets to be grouped into projects
    projectId      uuid             not null,
    -- A ticket must have a classifier (ticket type) which will determine which workflow is applied.
    ticketTypeId   uuid             not null,
    -- For a given ticket type, there is a workflow.  But the ticket must also know its current step int that workflow
    -- which is tracked separately.
    workflowStepId uuid             not null,
    -- We must know who created the ticket and to whom the ticket is currently assigned.
    authorId   uuid             not null,
    assignedId uuid             not null,
    -- permissions are granted to the owner, team and everyone --
    owner          permissions      not null default 'delete',
    team           permissions      not null default 'read',
    everyone       permissions      not null default 'read',
    -- descriptive text --
    summary        varchar(2048)    not null,
    description    text,
    -- a projectId must be in table projects
    foreign key (projectId) references projects (id),
    -- an author or assigned user must be in users
    foreign key (authorId) references users (id),
    foreign key (assignedId) references users (id),
    -- a ticketTypeId must be in the set of ticket types SUPPORTED BY THE project
    constraint ensureTicketTypeSupported check (isTicketTypeidSupported(projectId, ticketTypeId)),
    -- workflow step must be in workflowsteps.
    foreign key (workflowStepId) references workflowSteps (id),
    -- ticket id must be an entity.
    foreign key (id) references entity (id) on delete restrict
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

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketAssignee()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketAssignee(ticketId uuid, assignee uuid) returns integer as
$$
declare
    count integer;
begin
    update tickets set assignedUserId=assignee where id = ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketAuthor()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketAuthor(ticketId uuid, authorUserId uuid) returns integer as
$$
declare
    count integer;
begin
    update tickets set authorUserId=authorUserId where id = ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketProject()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketProject(ticketId uuid, projectId uuid) returns integer as
$$
declare
    count integer;
begin
    update tickets set projectId=projectId where id = ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketPermissions()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketPermissions(ticketId uuid, owner permissions, team permissions,
                                                   everyone permissions) returns integer as
$$
declare
    count integer;
begin
    update tickets set owner=owner, team=team, everyone=everyone where id = ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketType()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketType(ticketId uuid, ticketTypeId uuid) returns integer as
$$
declare
    count integer;
begin
    update tickets set ticketTypeId=ticketTypeId where id = ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
