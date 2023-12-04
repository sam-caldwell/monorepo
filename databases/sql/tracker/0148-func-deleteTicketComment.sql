/*
 * 0148-func-deleteTicketComment.sql
 * (c) 2023 Sam Caldwell.  See License.txt
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
