/*
 * 010-preventDelete.sql
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * A function which throws an exception to prevent deletes
 */

 create or replace function preventDelete() returns trigger as
 $$
 begin
     raise exception 'cannot delete records';
 end;
 $$ language plpgsql;
