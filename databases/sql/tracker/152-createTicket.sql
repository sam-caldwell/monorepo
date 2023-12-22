/*
 * 152-createTicket.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function createTicket(project uuid, author uuid, assignee uuid, type uuid,
                                        pOwner permissions, pTeam permissions, pEveryone permissions,
                                        ticketSummary varchar(2048), ticketDescription text) returns uuid as
$$
declare
    ticketId       uuid := (select createEntity('team'::entityType));

    /*
     * for a given new ticket, we expect to find the starting step in its type-specific workflow
     * and then to use that value when initializing our ticket state.
     */
    startingStepId uuid := (select id
                            from workflowsteps
                            where workflowid in (select workflowid
                                                 from tickettypes
                                                 where id = type
                                                 limit 1));
begin
    /*
     * ToDo: verify that the summary follows expectations...sanitization?  Maybe enforce convention?
     *       Question: Do we want to do that here or in app land/UI?
     */
    insert into tickets (id, projectId, ticketTypeId,
                         authorId, assignedId,
                         owner, team, everyone,
                         summary, description, workflowStepId)

    values (ticketId, project, type,
            author, assignee,
            pOwner, pTeam, pEveryone,
            ticketSummary, ticketDescription, startingStepId);
    return ticketId;
end;
$$ language plpgsql;
