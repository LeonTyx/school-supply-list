create table if not exists event_date
(
    start_time date,
    end_time date,
    event_id int
        constraint event_date_event_id_fk
            references event
            on delete cascade,
    id serial not null
);

create unique index event_date_id_uindex
    on event_date (id);

alter table event_date
    add constraint event_date_pk
        primary key (id);

