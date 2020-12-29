alter table "supply_list"
    drop constraint "supply_list_school_school_id_fk";

alter table "supply_list"
    drop column school_id;

