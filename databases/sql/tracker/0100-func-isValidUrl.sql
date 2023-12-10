/*
 * 0100-func-isValidUrl.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function isValidUrl(url text) returns bool as
$$
begin
    if (select count(*) from regexp_match(url, '^(http|https)?://.*\?.*=.*$')) = 0 then
        return false;
    end if;
    return true;
end;
$$
    language plpgsql;
