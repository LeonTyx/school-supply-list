create table school
(
    school_name varchar default 'Little Placeholder Elementary' not null,
    school_id serial not null
);

create unique index school_school_id_uindex
    on school (school_id);

alter table school
    add constraint school_pk
        primary key (school_id);

