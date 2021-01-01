alter table supply_list
    add list_id serial not null;

create unique index supply_list_list_id_uindex
    on supply_list (list_id);

alter table supply_list
    add constraint supply_list_pk
        primary key (list_id);

