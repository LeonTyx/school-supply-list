create table if not exists role
(
    role_id   serial                                not null
        constraint role_pk
            primary key,
    role_name varchar                               not null,
    role_desc varchar default ''::character varying not null
);

create unique index if not exists role_role_id_uindex
    on role (role_id);

create unique index if not exists role_role_name_uindex
    on role (role_name);

create table if not exists resource
(
    resource_id   serial                                not null
        constraint resource_pk
            primary key,
    resource_name varchar                               not null,
    resource_desc varchar default ''::character varying not null
);

create unique index if not exists resource_resource_id_uindex
    on resource (resource_id);

create unique index if not exists resource_resource_name_uindex
    on resource (resource_name);

create table if not exists role_resource_bridge
(
    rrb_id      serial                not null
        constraint resource_role_bridge_pk
            primary key,
    can_add     boolean default false not null,
    can_view    boolean default false not null,
    can_edit    boolean default false not null,
    can_delete  boolean default false not null,
    resource_id integer               not null
        constraint resource_role_bridge_resource_resource_id_fk
            references resource,
    role_id     integer               not null
        constraint resource_role_bridge_role_role_id_fk
            references role
            on delete cascade,
    constraint role_resource_bridge_pk
        unique (resource_id, role_id)
);

create unique index if not exists resource_role_bridge_rrb_id_uindex
    on role_resource_bridge (rrb_id);

create table if not exists user_role_bridge
(
    user_uuid uuid    not null
        constraint user_role_bridge_account_user_id_fk
            references account
            on delete cascade,
    role_id    integer not null
        constraint user_role_bridge_role_role_id_fk
            references role
            on delete cascade,
    constraint user_role_bridge_pk
        primary key (user_uuid, role_id)
);

