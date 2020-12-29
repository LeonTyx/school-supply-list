create table "supply-list"
(
    grade int not null,
    list_name varchar default 'Placeholder List' not null
);

create unique index "supply-list_list_name_uindex"
    on "supply-list" (list_name);

