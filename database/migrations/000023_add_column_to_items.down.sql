DELETE from supply_item;

alter table supply_item drop column category;

alter table supply_item drop constraint supply_item_supply_list_list_id_fk;

alter table supply_item drop column list_id;

alter table supply_item
    add constraint supply_item_supply_list_list_id_fk
        foreign key (list_id) references supply_list
            on delete cascade;