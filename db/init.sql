
create table if not exists public.article
(
    "Article_id"           integer default nextval('"article_Id_seq"'::regclass) not null
        constraint article_pk
            primary key
        constraint "article_Id_key"
            unique,
    "Content"              varchar,
);

alter table public.article
    owner to table_admin;

create unique index if not exists article_id_uindex
    on public.article ("Article_id");
