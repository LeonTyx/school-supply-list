alter table item_list_bridge
    drop constraint item_list_bridge_supply_item_id_fk;

alter table item_list_bridge
    add constraint item_list_bridge_supply_item_id_fk
        foreign key (item_id) references supply_item;

