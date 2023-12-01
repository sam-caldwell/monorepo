/*
 * 0148-func-updateTicketComment.sql
 * (c) 2023 Sam Caldwell.  See License.txt
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
