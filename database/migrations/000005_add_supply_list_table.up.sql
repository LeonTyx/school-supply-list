create table "supply_list"
(
    grade int not null,
    list_name varchar default 'Placeholder List' not null
);

create unique index "supply_list_list_name_uindex"
    on "supply_list" (list_name);

