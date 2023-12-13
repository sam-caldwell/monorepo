/*
 * 0200-0010-func-createTicketAttachment.sql
 * (c) 2023 Sam Caldwell.  See License.txt
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
