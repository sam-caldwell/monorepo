/*
 * 0145-func-createTicket.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function createTicket(name varchar(64), projectId uuid, authorId uuid, assigneeId uuid,
                                        ticketTypeId uuid, owner permissions, team permissions, everyone permissions,
                                        summary varchar(2048), workflowStepId uuid, description text) returns uuid as
$$
declare
    ticketId uuid;
begin
    ticketId := gen_random_uuid();
    insert into users (id, name, projectId, authorId, assigneeId, ticketTypeId, owner, team, everyone, summary,
                       workflowStepId, description)
    values (ticketId, name, projectId, authorId, assigneeId, owner, team, everyone, summary, workflowStepId,
            description);
    return ticketId;
end;
$$ language plpgsql;
