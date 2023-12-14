/*
 * 010-isValidUrl.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
/*
 * Supported URL Patterns:
 *
 *     http://simpleHostName
 *     https://google.com
 *     http://google.com
 *     http://google.com/
 *     http://www.simpleHostName
 *     http://www.google.com
 *     http://www.google.com/
 *     http://www.happyUrl.tld/good.image.name.png
 *     http://www.happyUrl.tld/good.image.name.png
 *     http://www.happyUrl.tld/good.image.name.png/?query=param
 *     http://www.happyUrl.tld/good.image.name.png/?query1=param1&query1=param1
 *     https://www.happyUrl.tld/good.image.name.png
 *     https://www.happyUrl.tld/good.image.name.png/?query=param
 *     https://www.happyUrl.tld/good.image.name.png/?query1=param1&query1=param1
 *     http://happyUrl.tld/good.image.name.png
 *     http://happyUrl.tld/good.image.name.png
 *     http://happyUrl.tld/good.image.name.png/?query=param
 *     http://happyUrl.tld/good.image.name.png/?query1=param1&query1=param1
 *     https://happyUrl.tld/good.image.name.png
 *     https://happyUrl.tld/good.image.name.png/?query=param
 *     https://happyUrl.tld/good.image.name.png/?query1=param1&query1=param1
 */

create or replace function isValidUrl(url text)
    returns boolean
as
$$
declare
    regex  text := '^(http|https):\/\/(?:www\.|)?([-a-zA-Z0-9@:%._\+~#=]+\.[a-z]+\b)*(\/[\/\d\w\.-]*)*(?:[\?])*(.+)*$';
begin
    return (select regexp_match(url, regex,'i') is not null);
end;
$$ language plpgsql;
