create table if not exists supply_item
(
    id          serial             not null,
    supply_name varchar default '' not null,
    supply_desc varchar
);

create unique index supply_item_id_uindex
    on supply_item (id);

alter table supply_item
    add constraint supply_item_pk
        primary key (id);