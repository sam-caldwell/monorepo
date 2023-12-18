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
            CREATE TYPE entityType AS ENUM ('avatar', 'icon', 'user', 'team','teamAssociation',
                'workflow','workflow_step','workflow_action','ticket_type','project','ticket','projectTypes',
                'ticketTypeAssociation','attachment','comment','property','other');
        END IF;
    END
$$;
