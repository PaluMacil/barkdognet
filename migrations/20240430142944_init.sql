-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table public.sys_user
(
    id                                integer generated always as identity
        constraint sys_user_pk
            primary key,
    email                             varchar(100)               not null
        constraint sys_user_pk_3
            unique,
    email_verified                    boolean      default false not null,
    email_verification_hash           bytea,
    email_verification_code_issued_at timestamp with time zone,
    password_reset_hash               bytea,
    password_reset_code_issued_at     timestamp with time zone,
    display_name                      varchar(100)               not null
        constraint sys_user_pk_2
            unique,
    given_name                        varchar(100) default ''    NOT NULL,
    family_name                       varchar(100) default ''    NOT NULL,
    phone_number                      varchar(16),
    phone_number_verified             boolean      default false not null,
    location                          varchar(200) default ''    NOT NULL,
    locked                            boolean      default false not null,
    password_hash                     bytea,
    last_login_at                     timestamp with time zone,
    created_at                        TIMESTAMP WITH TIME ZONE
                                                   DEFAULT now() NOT NULL,
    bio                               varchar(500) default ''    NOT NULL,
    social_links                      json                       NOT NULL DEFAULT '[]'
);
CREATE INDEX idx_sys_user_email_verified ON public.sys_user (email_verified);
CREATE INDEX idx_sys_user_created_at ON public.sys_user (created_at);

create table public.sys_role
(
    id           integer generated always as identity
        primary key,
    display_name varchar(50)                not null
        constraint sys_role_pk
            unique,
    description  varchar(200) default ''    not null,
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
CREATE INDEX idx_m2m_user_role_sys_user_id ON public.m2m_user_role (sys_user_id);
CREATE INDEX idx_m2m_user_role_sys_role_id ON public.m2m_user_role (sys_role_id);

create table public.sys_session
(
    id            integer generated always as identity
        constraint sys_session_pk
            primary key,
    sys_user_id   int          not null,
    session_token text         not null
        constraint sys_session_session_token
            unique,
    created_at    TIMESTAMP WITH TIME ZONE
        DEFAULT now()          NOT NULL,
    ip_address    varchar(60)  not null,
    user_agent    varchar(200) not null,
    FOREIGN KEY (sys_user_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_sys_session_sys_user_id ON public.sys_session (sys_user_id);

create table public.blog_category
(
    id            integer generated always as identity
        primary key,
    category_name varchar(100) not null
);
CREATE INDEX idx_blog_category_category_name ON public.blog_category (category_name);

CREATE TYPE post_status AS ENUM ('draft', 'published', 'hidden', 'archived');

create table public.blog_post
(
    id          integer generated always as identity
        primary key,
    category_id int          not null,
    title       varchar(200) not null,
    slug        varchar(200) not null,
    summary     text,
    status      post_status  not null,
    keywords    varchar(200) not null,
    body        text         not null,
    author_id   int          not null,
    created_at  TIMESTAMP WITH TIME ZONE
        DEFAULT now()        NOT NULL,
    FOREIGN KEY (author_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (category_id) REFERENCES public.blog_category (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_blog_post_slug ON public.blog_post (slug);
CREATE INDEX idx_blog_post_status ON public.blog_post (status);
CREATE INDEX idx_blog_post_author_id ON public.blog_post (author_id);
CREATE INDEX idx_blog_post_category_id ON public.blog_post (category_id);

CREATE TYPE comment_status AS ENUM ('approved', 'pending', 'spam');

create table public.blog_comment
(
    id           integer generated always as identity
        primary key,
    author_id    int                              not null,
    blog_post_id int                              not null,
    body         text                             not null,
    status       comment_status default 'pending' not null,
    created_at   TIMESTAMP WITH TIME ZONE
                                DEFAULT now()     NOT NULL,
    FOREIGN KEY (author_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (blog_post_id) REFERENCES public.blog_post (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_blog_comment_status ON public.blog_comment (status);
CREATE INDEX idx_blog_comment_author_id ON public.blog_comment (author_id);
CREATE INDEX idx_blog_comment_blog_post_id ON public.blog_comment (blog_post_id);

create table public.blog_post_like
(
    user_id    int                                    not null,
    post_id    int                                    not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    PRIMARY KEY (user_id, post_id),
    FOREIGN KEY (user_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (post_id) REFERENCES public.blog_post (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_blog_post_like_post_id ON public.blog_post_like (post_id);

create table public.blog_comment_like
(
    user_id    int                                    not null,
    comment_id int                                    not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    PRIMARY KEY (user_id, comment_id),
    FOREIGN KEY (user_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (comment_id) REFERENCES public.blog_comment (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_blog_comment_like_comment_id ON public.blog_comment_like (comment_id);

create table public.blog_tag
(
    id           integer generated always as identity primary key,
    display_name varchar(50) not null
);

create table public.m2m_blog_post_tag
(
    blog_post_id int not null,
    tag_id       int not null,
    PRIMARY KEY (blog_post_id, tag_id),
    FOREIGN KEY (blog_post_id) REFERENCES public.blog_post (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES public.blog_tag (id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
