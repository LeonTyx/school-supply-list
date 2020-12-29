alter table "supply_list"
    add school_id int not null default 0;

alter table "supply_list"
    add constraint "supply_list_school_school_id_fk"
        foreign key (school_id) references school;

