/*
 * 010-mimeType.sql
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * This is the enumerated set of supported image mimeTypes.
 */
DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'mimeType') THEN
            CREATE TYPE mimeType AS ENUM ('image/gif', 'image/jpeg', 'image/png', 'image/webp');
        END IF;
    END
$$;
