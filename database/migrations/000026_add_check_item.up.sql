create table if not exists checked_items
(
    item_id int not null
        constraint checked_items_supply_item_id_fk
            references supply_item
            on delete cascade,
    user_uuid uuid not null
        constraint checked_items_account_user_id_fk
            references account (user_id)
            on delete cascade,
    constraint checked_items_pk
        primary key (item_id, user_uuid)
);

