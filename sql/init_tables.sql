create table if not exists events (
    id serial primary key,
    UUID text not null,
    title text,
    datetime text,
    duration text,
    description text,
    userid text,
    notify text
);