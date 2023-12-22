/*
 * 122-getProjectTicketTypes.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function getProjectTicketTypes(thisPid uuid,
                                                 maxRecords integer default 1000,
                                                 startAt integer default 0)
    returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(data) into result
    from (
             select jsonb_build_object(
                            'id', tt.id,
                            'name', tt.name,
                            'iconId', tt.iconid::uuid,
                            'workflowId', tt.workflowid::uuid,
                            'created', tt.created,
                            'description', tt.description
                        ) as data
             from projectTicketTypes ptt
                      join tickettypes tt
                           on ptt.ticketTypeId = tt.id
             where ptt.projectId = thisPid
             order by tt.name asc
             limit maxRecords offset startAt
         ) as subquery;

    return coalesce(result, '[]'::jsonb);
end;
$$ language plpgsql;
