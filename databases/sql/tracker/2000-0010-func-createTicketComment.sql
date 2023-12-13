/*
 * 2000-0010-func-createTicketComment.sql
 * (c) 2023 Sam Caldwell.  See License.txt
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
