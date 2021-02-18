create table event
(
    id serial not null,
    title varchar not null,
    description varchar not null
);

create unique index event_id_uindex
    on event (id);

alter table event
    add constraint event_pk
        primary key (id);

