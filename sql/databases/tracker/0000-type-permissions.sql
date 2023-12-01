/*
 * 0000-type-permissions.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create a permissions type (enum)
 *       read - basic permission to view/read an entity
 *       create - a permission that includes 'read' but also allows entity creation
 *       update - a permission that includes 'create' and 'read' to modify an existing entity
 *       delete - a permission that includes 'delete' permissions as well as update, create and read.
 */
DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'permissions') THEN
            create type permissions as enum ('none','read', 'create','update','delete');
        END IF;
    END
$$;
