/*
 * 0200-0020-func-deleteTicketAttachment.sql
 * (c) 2023 Sam Caldwell.  See License.txt
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
