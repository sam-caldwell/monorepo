/*
 * 010-boundsCheckUpper.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function boundsCheckUpper(n integer, upperBound integer) returns boolean as
$$
begin
    if (n > upperBound) then
        raise exception 'upper bounds check error(boundary: %)(value:%)',upperBound,n;
    end if;
    return true;
end ;
$$ language plpgsql;
