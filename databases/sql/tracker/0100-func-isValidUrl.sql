/*
 * 0100-func-isValidUrl.sql
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

CREATE OR REPLACE FUNCTION isValidUrl(url text)
    RETURNS bool
AS $$
BEGIN
    DECLARE
        regex text := '^(http|https)?:\/\/(?:www\.|)?([-a-zA-Z0-9@:%._\+~#=]+\.[a-z]+\b)*(\/[\/\d\w\.-]*)*(?:[\?])*(.+)*$';
    BEGIN
        IF (select count(*) as c from regexp_match(url, regex)) > 0 THEN
            RETURN TRUE;
        ELSE
            RETURN FALSE;
        END IF;
    END;
END;
$$ LANGUAGE plpgsql;
