alter table if exists supply_list
    drop constraint supply_list_pk;

drop index if exists supply_list_list_id_uindex;

alter table if exists supply_list
    drop list_id;