/*
 * 150-tickets.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create table if not exists tickets
(
    id             uuid primary key not null,
    projectId      uuid             not null,
    authorUserId   uuid             not null,
    assignedUserId uuid             not null,
    ticketTypeId   uuid             not null,
    -- permissions are granted to the owner, team and everyone --
    owner          permissions      not null default 'delete',
    team           permissions      not null default 'read',
    everyone       permissions      not null default 'read',
    -- descriptive text --
    summary        varchar(2048)    not null,
    workflowStepId uuid             not null,
    -- --
    created        timestamp        not null default now(),
    description    text,
    foreign key (projectId) references projects (id),
    foreign key (authorUserId) references users (id),
    foreign key (assignedUserId) references users (id),
    foreign key (ticketTypeId) references ticketTypes (id),
    foreign key (workflowStepId) references workflowSteps (id),
    foreign key (id) references entity (id) on delete restrict
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create index if not exists ndxTicketSummary on tickets (summary);
create index if not exists ndxTicketCreated on tickets (created);
create index if not exists ndxTicketOwner on tickets (owner);
create index if not exists ndxTicketTeam on tickets (team);
create index if not exists ndxTicketEveryone on tickets (everyone);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createTicketType()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createTicket(tName varchar(64), tProject uuid, tAuthor uuid, tAssignee uuid,
                                        ticketType uuid, pOwner permissions, pTeam permissions, pEveryone permissions,
                                        tSummary varchar(2048), tWorkflowStepId uuid, tDescription text) returns uuid as
$$
declare
    ticketId uuid;
begin
    ticketId := (select createEntity('team'::entityType));
    insert into tickets (id, name, projectId, authorId, assigneeId, ticketTypeId, owner, team, everyone, summary,
                         workflowStepId, description)
    values (ticketId, tName, tProject, tAuthor, tAssignee, ticketType,pOwner, pTeam, pEveryone, tSummary, tWorkflowStepId,
            tDescription);
    return ticketId;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteTicket()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteTicket(ticketId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from tickets where id = ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketByAssignee()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketByAssignee(thisUserId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'projectId', projectId,
            'authorUserId', authorUserId,
            'assignedUserId', thisUserId,
            'ticketTypeId', ticketTypeId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'summary', summary,
            'workflowStepId', workflowStepId,
            'description', description
        )) as workflow
    into result
    from tickets
    where assignedUserId = assignedUserId;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketByAuthor()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketByAuthor(authorId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'projectId', projectId,
            'authorUserId', authorUserId,
            'assignedUserId', assignedUserId,
            'ticketTypeId', ticketTypeId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'subject', summary,
            'workflowStepId', workflowStepId,
            'description', description
        )) as workflow
    into result
    from tickets
    where authorUserId = authorId;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketByPermEveryone()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketByPermEveryone(thisPermission permissions, pageLimit integer,
                                                   pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'projectId', projectId,
            'authorUserId', authorUserId,
            'assignedUserId', assignedUserId,
            'ticketTypeId', ticketTypeId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'subject', summary,
            'workflowStepId', workflowStepId,
            'description', description
        )) as workflow
    into result
    from tickets
    where everyone = thisPermission
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketByPermOwner()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketByPermOwner(thisPermission permissions, pageLimit integer,
                                                pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'projectId', projectId,
            'authorUserId', authorUserId,
            'assignedUserId', assignedUserId,
            'ticketTypeId', ticketTypeId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'subject', subject,
            'workflowStepId', workflowStepId,
            'description', description
        )) as workflow
    into result
    from ticket
    where owner == thisPermission
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketByPermTeam()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketByPermTeam(thisPermission permissions, pageLimit integer,
                                               pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'projectId', projectId,
            'authorUserId', authorUserId,
            'assignedUserId', assignedUserId,
            'ticketTypeId', ticketTypeId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'subject', subject,
            'workflowStepId', workflowStepId,
            'description', description
        )) as workflow
    into result
    from ticket
    where team == thisPermission
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketByProject()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketByProject(projectId uuid, pageLimit integer,
                                              pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'projectId', projectId,
            'authorUserId', authorUserId,
            'assignedUserId', assignedUserId,
            'ticketTypeId', ticketTypeId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'subject', subject,
            'workflowStepId', workflowStepId,
            'description', description
        )) as workflow
    into result
    from ticket
    where projectId == projectId
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketByTicketType()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketByTicketType(ticketTypeId uuid, pageLimit integer,
                                                 pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'projectId', projectId,
            'authorUserId', authorUserId,
            'assignedUserId', assignedUserId,
            'ticketTypeId', ticketTypeId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'subject', subject,
            'workflowStepId', workflowStepId,
            'description', description
        )) as workflow
    into result
    from ticket
    where ticketTypeId == ticketTypeId
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketByWorkflowStepId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketByWorkflowStepId(stepId uuid, pageLimit integer,
                                                     pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'projectId', projectId,
            'authorUserId', authorUserId,
            'assignedUserId', assignedUserId,
            'ticketTypeId', ticketTypeId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'subject', subject,
            'workflowStepId', workflowStepId,
            'description', description
        )) as workflow
    into result
    from ticket
    where workflowStepId == stepId
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;
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
    update ticket set assignedUserId=assignee where id = ticketId;
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
    update ticket set authorUserId=authorUserId where id = ticketId;
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
    update ticket set projectId=projectId where id = ticketId;
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
    update ticket set owner=owner, team=team, everyone=everyone where id = ticketId;
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
    update ticket set ticketTypeId=ticketTypeId where id = ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createTicketAttachment()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createTicketAttachment(ticketId uuid, authorId uuid,
                                                  permAuthor permissions, permTeam permissions,
                                                  permEveryone permissions) returns uuid as
$$
declare
    commentId uuid;
begin
    commentId := gen_random_uuid();
    insert into attachment (id, ticketId, authorId, permAuthor, permTeam, permEveryone, everyone)
    values (commentId, ticketId, authorId, permAuthor, permTeam, permEveryone, url);
    return commentId;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteTicketAttachment()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteTicketAttachment(commentId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from attachment where id = commentId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketAttachmentByTicket()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketAttachmentByTicket(ticketId uuid,
                                                       pageLimit integer,
                                                       pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'ticketId', ticketId,
            'authorId', authorId,
            'author', author,
            'team', team,
            'everyone', everyone
        )) as workflow
    into result
    from attachment
    where ticketId == ticketId
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createTicketComment()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createTicketComment(ticketId uuid, authorId uuid,
                                               permAuthor permissions, permTeam permissions, permEveryone permissions,
                                               comment text) returns uuid as
$$
declare
    commentId uuid;
begin
    commentId := gen_random_uuid();
    insert into comment (id, ticketId, authorId, permAuthor, permTeam, permEveryone, everyone, comment)
    values (commentId, ticketId, authorId, permAuthor, permTeam, permEveryone, comment);
    return commentId;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteTicketComment()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteTicketComment(commentId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from comment where id = commentId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketComment()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketComment(commentId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'ticketId', ticketId,
            'authorId', authorId,
            'author', author,
            'team', team,
            'everyone', everyone,
            'comment', comment
        )) as workflow
    into result
    from comment
    where id == commentId
    limit 1;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketCommentByTicket()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketCommentByTicket(ticketId uuid,
                                                    pageLimit integer,
                                                    pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'ticketId', ticketId,
            'authorId', authorId,
            'author', author,
            'team', team,
            'everyone', everyone,
            'comment', comment
        )) as workflow
    into result
    from comment
    where ticketId == ticketId
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketComment()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketComment(commentId uuid,
                                               ticketId uuid,
                                               authorId uuid,
                                               author permissions,
                                               team permissions,
                                               everyone permissions,
                                               comment text) returns integer as
$$
declare
    count integer;
begin
    update comment
    set ticketId=ticketId,
        authorId=authorId,
        author=author,
        team=team,
        everyone=everyone,
        comment=comment
    where id = commentId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
