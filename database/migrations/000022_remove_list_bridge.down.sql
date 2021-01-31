create table if not exists item_list_bridge
(
    ilb_id   serial  not null
        constraint item_list_bridge_pk
            primary key,
    list_id  integer not null
        constraint item_list_bridge_supply_list_list_id_fk
            references supply_list
            on delete cascade,
    item_id  integer not null
        constraint item_list_bridge_supply_item_id_fk
            references supply_item
            on delete restrict,
    category varchar,
    constraint item_list_bridge_pk_2
        unique (list_id, item_id)
);

create unique index if not exists item_list_bridge_ilb_id_uindex
    on item_list_bridge (ilb_id);

