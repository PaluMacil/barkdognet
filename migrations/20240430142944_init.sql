-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE public.sys_tenant
(
    id            INTEGER GENERATED ALWAYS AS IDENTITY
        CONSTRAINT sys_tenant_pk
            PRIMARY KEY,
    display_name  VARCHAR(255)                           NOT NULL,
    active        BOOLEAN                  DEFAULT true  NOT NULL,
    api_subdomain VARCHAR(255)                           NOT NULL,
    ui_domain     VARCHAR(255),
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

COMMENT ON COLUMN public.sys_tenant.api_subdomain IS 'The domain where the API is hosted.';
COMMENT ON COLUMN public.sys_tenant.ui_domain IS 'The domain where the UI is hosted. Can be null if API and UI are hosted together.';
COMMENT ON COLUMN public.sys_tenant.active IS 'Flag to indicate whether the tenant is active or inactive.';

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
    bio                               varchar(500) default ''    NOT NULL,
    social_links                      json         DEFAULT '[]'  NOT NULL,
    locked                            boolean      default false not null,
    password_hash                     bytea,
    last_login_at                     timestamp with time zone,
    created_at                        TIMESTAMP WITH TIME ZONE
                                                   DEFAULT now() NOT NULL
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
    tenant_id    INTEGER                    NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE
                              DEFAULT now() NOT NULL,
    FOREIGN KEY (tenant_id) REFERENCES public.sys_tenant (id)
);
CREATE INDEX idx_sys_role_tenant_id ON public.sys_role (tenant_id);

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
    ip_address    varchar(60)  not null,
    user_agent    varchar(200) not null,
    tenant_id     INTEGER      NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE
        DEFAULT now()          NOT NULL,
    FOREIGN KEY (sys_user_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (tenant_id) REFERENCES public.sys_tenant (id)
);
CREATE INDEX idx_sys_session_sys_user_id ON public.sys_session (sys_user_id);
CREATE INDEX idx_sys_session_tenant_id ON public.sys_session (tenant_id);

create table public.blog_category
(
    id            integer generated always as identity
        primary key,
    category_name varchar(100) not null,
    tenant_id     INTEGER      NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE
        DEFAULT now()          NOT NULL,
    FOREIGN KEY (tenant_id) REFERENCES public.sys_tenant (id)
);
CREATE INDEX idx_blog_category_category_name ON public.blog_category (category_name);
CREATE INDEX idx_blog_category_tenant_id ON public.blog_category (tenant_id);

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
    tenant_id   INTEGER      NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE
        DEFAULT now()        NOT NULL,
    FOREIGN KEY (author_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (category_id) REFERENCES public.blog_category (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (tenant_id) REFERENCES public.sys_tenant (id)
);
CREATE INDEX idx_blog_post_slug ON public.blog_post (slug);
CREATE INDEX idx_blog_post_status ON public.blog_post (status);
CREATE INDEX idx_blog_post_author_id ON public.blog_post (author_id);
CREATE INDEX idx_blog_post_category_id ON public.blog_post (category_id);
CREATE INDEX idx_blog_post_tenant_id ON public.blog_post (tenant_id);

CREATE TYPE comment_status AS ENUM ('approved', 'pending', 'spam');

create table public.blog_comment
(
    id           integer generated always as identity
        primary key,
    author_id    int                              not null,
    blog_post_id int                              not null,
    body         text                             not null,
    status       comment_status default 'pending' not null,
    tenant_id    INTEGER                          NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE
                                DEFAULT now()     NOT NULL,
    FOREIGN KEY (author_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (blog_post_id) REFERENCES public.blog_post (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (tenant_id) REFERENCES public.sys_tenant (id)
);
CREATE INDEX idx_blog_comment_status ON public.blog_comment (status);
CREATE INDEX idx_blog_comment_author_id ON public.blog_comment (author_id);
CREATE INDEX idx_blog_comment_blog_post_id ON public.blog_comment (blog_post_id);
CREATE INDEX idx_blog_comment_tenant_id ON public.blog_comment (tenant_id);

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
    display_name varchar(50) not null,
    created_at   TIMESTAMP WITH TIME ZONE
        DEFAULT now()        NOT NULL,
    tenant_id    INTEGER     NOT NULL,
    FOREIGN KEY (tenant_id) REFERENCES public.sys_tenant (id)
);
CREATE INDEX idx_blog_tag_tenant_id ON public.blog_tag (tenant_id);

create table public.m2m_blog_post_tag
(
    blog_post_id int not null,
    tag_id       int not null,
    PRIMARY KEY (blog_post_id, tag_id),
    FOREIGN KEY (blog_post_id) REFERENCES public.blog_post (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES public.blog_tag (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE public.oidc_provider
(
    id              INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    display_name    VARCHAR(255)                           NOT NULL,
    issuer_url      VARCHAR(255)                           NOT NULL,
    discovery_url   VARCHAR(255)                           NOT NULL,
    scopes          TEXT[]                                 NOT NULL DEFAULT ARRAY ['oidc', 'profile', 'email'],
    client_id       VARCHAR(255)                           NOT NULL,
    client_secret   VARCHAR(255)                           NOT NULL,
    redirect_url    VARCHAR(255)                           NOT NULL,
    access_type     VARCHAR(50)                            NOT NULL,
    azure_tenant_id VARCHAR(100),
    active          BOOLEAN                  DEFAULT true  NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    tenant_id       INTEGER                                NOT NULL,
    FOREIGN KEY (tenant_id) REFERENCES public.sys_tenant (id)
);
CREATE INDEX idx_oidc_provider_tenant_id ON public.oidc_provider (tenant_id);

CREATE TABLE public.user_oidc
(
    sys_user_id      INTEGER                                NOT NULL,
    oidc_provider_id INTEGER                                NOT NULL,
    sub              VARCHAR(255)                           NOT NULL,
    created_at       TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    PRIMARY KEY (sys_user_id, oidc_provider_id),
    FOREIGN KEY (sys_user_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (oidc_provider_id) REFERENCES public.oidc_provider (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX idx_user_oidc_sys_user_id ON public.user_oidc (sys_user_id);
CREATE INDEX idx_user_oidc_oidc_provider_id ON public.user_oidc (oidc_provider_id);
CREATE INDEX idx_user_oidc_sub ON public.user_oidc (sub);

CREATE TABLE public.invitation
(
    id              INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    inviter_id      INTEGER      NOT NULL,
    invitee_email   VARCHAR(100) NOT NULL,
    invitation_code UUID         NOT NULL    DEFAULT gen_random_uuid(),
    pending         BOOLEAN      NOT NULL    DEFAULT true,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    expires_at      TIMESTAMP WITH TIME ZONE DEFAULT (now() + interval '30 days') NOT NULL,
    accepted_at     TIMESTAMP WITH TIME ZONE,
    tenant_id       INTEGER      NOT NULL,
    FOREIGN KEY (inviter_id) REFERENCES public.sys_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (tenant_id) REFERENCES public.sys_tenant (id),
    UNIQUE (invitation_code)
);
CREATE INDEX idx_invitations_invitee_email ON public.invitation (invitee_email);
CREATE INDEX idx_invitations_pending ON public.invitation (pending);
CREATE INDEX idx_invitations_created_at ON public.invitation (created_at);
CREATE INDEX idx_invitations_tenant_id ON public.invitation (tenant_id);


-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
