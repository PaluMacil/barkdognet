-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table public.sys_user
(
    id                    integer generated always as identity
        constraint sys_user_pk
            primary key,
    email                 varchar(100)               not null
        constraint sys_user_pk_3
            unique,
    email_confirmed       boolean      default false not null,
    display_name          varchar(100)               not null
        constraint sys_user_pk_2
            unique,
    given_name            varchar(100) default NULL::character varying,
    phone_number          varchar(16),
    phone_number_verified boolean      default false,
    family_name           varchar(100) default NULL::character varying,
    locked                boolean      default false not null,
    password_hash         bytea,
    last_login_at         timestamp with time zone,
    created_at            TIMESTAMP WITH TIME ZONE
        DEFAULT now() NOT NULL
);
create table public.sys_role
(
    id           integer generated always as identity
        primary key,
    display_name varchar(50) not null
        constraint sys_role_pk
            unique,
    created_at   TIMESTAMP WITH TIME ZONE
        DEFAULT now() NOT NULL
);
create table public.m2m_user_role
(
    sys_user_id int not null,
    sys_role_id int not null,
    PRIMARY KEY (sys_user_id, sys_role_id),
    FOREIGN KEY (sys_user_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (sys_role_id) REFERENCES public.sys_role (id) ON DELETE CASCADE ON UPDATE CASCADE
);
create table public.sys_session
(
    id            integer generated always as identity
        constraint sys_session_pk
            primary key,
    sys_user_id   int  not null,
    session_token text not null
        constraint sys_session_session_token
            unique,
    created_at    TIMESTAMP WITH TIME ZONE
        DEFAULT now() NOT NULL,
    FOREIGN KEY (sys_user_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE
);
create table public.blog_category
(
    id            integer generated always as identity
        primary key,
    category_name varchar(100) not null
);
create table public.blog_post
(
    id          integer generated always as identity
        primary key,
    category_id int not null,
    title       varchar(200) not null,
    keywords    varchar(200) not null,
    body        text not null,
    author_id   int not null,
    created_at  TIMESTAMP WITH TIME ZONE
        DEFAULT now() NOT NULL,
    FOREIGN KEY (author_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (category_id) REFERENCES public.blog_category (id) ON DELETE CASCADE ON UPDATE CASCADE
);
create table public.blog_comment
(
    id           integer generated always as identity
        primary key,
    author_id    int not null,
    blog_post_id int not null,
    body         text not null,
    created_at   TIMESTAMP WITH TIME ZONE
        DEFAULT now() NOT NULL,
    FOREIGN KEY (author_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (blog_post_id) REFERENCES public.blog_post (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
