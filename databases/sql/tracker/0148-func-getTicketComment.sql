/*
 * 0148-func-getTicketComment.sql
 * (c) 2023 Sam Caldwell.  See License.txt
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
        )) as workflow into result
    from comment
    where id == commentId
    limit 1;
    return result;
end;
$$ language plpgsql;
