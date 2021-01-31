DELETE from supply_item;

alter table supply_item
    add category varchar;

alter table supply_item
    add list_id int not null;

alter table supply_item
    add constraint supply_item_supply_list_list_id_fk
        foreign key (list_id) references supply_list
            on delete cascade;

