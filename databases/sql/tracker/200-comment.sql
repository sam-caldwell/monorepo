/*
 * 200-comment.sql
 * (c) 2023 Sam Caldwell. See License.txt.
 *
 * A collection of comments for any ticket.
 */
create table if not exists comment
(
    id       uuid primary key not null,
    ticketId uuid not null,
    authorId uuid not null,
    author   permissions not null default 'delete',
    team     permissions not null default 'read',
    everyone permissions not null default 'read',
    -- --
    created  timestamp        not null default now(),
    comment  text not null,
    foreign key (ticketId) references tickets (id),
    foreign key (authorId) references users (id),
    foreign key (id) references entity (id) on delete restrict
);

create index if not exists ndxCommentAuthor on comment (author);
create index if not exists ndxCommentTeam on comment (team);
create index if not exists ndxCommentEveryone on comment (everyone);

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
    commentId := (select createEntity('comment'::entityType));
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
