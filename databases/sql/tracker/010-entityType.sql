/*
 * 010-entityType.sql
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * An entity is a universal object identifier (uuid) for the entire tracker database.
 * This is intended as the basis of our accountability system.  The entity system provides
 * a write-only object registry with timestamps and context.
 */
DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'entityType') THEN
            CREATE TYPE entityType AS ENUM ('attachment', 'avatar', 'comment', 'icon', 'project', 'projectTypes',
                'property', 'team', 'teamAssociation', 'ticket', 'ticketType', 'ticketTypeAssociation', 'user',
                'workflow', 'workflowAction', 'workflowStep');
        END IF;
    END
$$;
