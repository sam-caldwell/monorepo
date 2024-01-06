/*
 * 010-isValidEmail.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
CREATE OR REPLACE FUNCTION is_valid_email(inpEmail VARCHAR) RETURNS BOOLEAN AS $$
DECLARE
    email_pattern VARCHAR := '^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$';
BEGIN
    -- Check if inpEmail matches the email pattern
    RETURN inpEmail ~ email_pattern;
END;
$$ LANGUAGE plpgsql;
