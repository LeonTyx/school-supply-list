alter table item_list_bridge
    drop constraint item_list_bridge_supply_list_list_id_fk;

alter table item_list_bridge
    add constraint item_list_bridge_supply_list_list_id_fk
        foreign key (list_id) references supply_list
            on delete cascade;

