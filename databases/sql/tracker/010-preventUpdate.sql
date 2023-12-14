/*
 * 010-preventUpdate.sql
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * A function which throws an exception to prevent updates
 */

create or replace function preventUpdate() returns trigger as
$$
begin
    raise exception 'cannot update records';
end;
$$ language plpgsql;
