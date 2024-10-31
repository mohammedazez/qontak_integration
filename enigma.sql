create table if not exists enigma.__ClientUpdate
(
    id         int auto_increment
    primary key,
    OutboundId int null,
    FromClient int null,
    ToClient   int null
);

create table if not exists enigma.api_logs
(
    id              bigint auto_increment
    primary key,
    url             varchar(255)                           not null,
    method          varchar(255)                           not null,
    request_header  text                                   not null,
    request_body    text                                   not null,
    response_status text                                   not null,
    response_header text                                   not null,
    response_body   text                                   not null,
    created_at      timestamp    default CURRENT_TIMESTAMP null,
    created_by      varchar(100) default (_utf8mb4'Admin') not null
    );

create index api_logs_created_at_IDX
    on enigma.api_logs (created_at);

create table if not exists enigma.client
(
    id          bigint auto_increment
    primary key,
    client_code varchar(255)                           null,
    client_name varchar(255)                           not null,
    status      int          default (1)               null,
    created_at  timestamp    default CURRENT_TIMESTAMP null,
    created_by  varchar(100) default (_utf8mb4'Admin') not null
    );

create table if not exists enigma.client_copy
(
    id          bigint auto_increment
    primary key,
    client_code varchar(255)                                 null,
    client_name varchar(255)                                 not null,
    status      int          default 1                       null,
    created_at  timestamp    default CURRENT_TIMESTAMP       null,
    created_by  varchar(100) default '_utf8mb4\\''Admin\\''' not null
    );

create table if not exists enigma.client_vendors
(
    id           bigint auto_increment
    primary key,
    channel_id   varchar(255) default (_utf8mb4'')      null,
    channel      varchar(100)                           not null,
    client_id    bigint                                 not null,
    vendor_id    bigint                                 not null,
    vendor_alias varchar(255)                           not null,
    status       int          default (1)               null,
    created_at   timestamp    default CURRENT_TIMESTAMP null,
    created_by   varchar(100) default (_utf8mb4'Admin') not null,
    constraint UQ_ALL
    unique (channel, client_id, vendor_id, vendor_alias)
    );

create table if not exists enigma.conversation_logs
(
    id                  bigint auto_increment
    primary key,
    client_id           int          default (0)               null,
    contact_id          varchar(255)                           null,
    message_id          varchar(255)                           null,
    conversation_id     varchar(255)                           null,
    conversation_status varchar(255)                           null,
    channel_id          varchar(255)                           null,
    platform            varchar(255)                           null,
    channel             varchar(255)                           null,
    direction           varchar(255)                           null,
    message_status      varchar(255)                           null,
    `from`              varchar(255)                           null,
    `to`                varchar(255)                           null,
    msisdn              varchar(255)                           null,
    contact_status      varchar(255)                           null,
    type                varchar(255) default (_utf8mb4'')      null,
    request_body        text                                   not null,
    message_json        text                                   not null,
    contact_json        text                                   not null,
    conversation_json   text                                   not null,
    status              int          default (0)               null,
    created_at          timestamp    default CURRENT_TIMESTAMP null
    );

create table if not exists enigma.flow_log
(
    id            bigint auto_increment
    primary key,
    client_id     int                                 null,
    outbound_json text                                null,
    status        int       default 1                 null,
    created_at    timestamp default CURRENT_TIMESTAMP null
)
    avg_row_length = 2482;

create table if not exists enigma.inbound
(
    id             int auto_increment
    primary key,
    contactid      varchar(200) null,
    contact        text         null,
    conversation   text         null,
    message        text         null,
    conversationid varchar(200) null,
    msisdn         varchar(30)  null,
    messageid      varchar(200) null,
    channelId      varchar(200) null,
    jsonplain      text         null,
    createddate    datetime     null
    );

create table if not exists enigma.inbound_raw_data
(
    id         bigint unsigned auto_increment
    primary key,
    message_id varchar(100)                       null,
    raw_data   text                               null,
    created_at datetime default CURRENT_TIMESTAMP null
    )
    collate = utf8mb4_general_ci;

create index inbound_raw_data_created_at_IDX
    on enigma.inbound_raw_data (created_at desc);

create index inbound_raw_data_message_id_IDX
    on enigma.inbound_raw_data (message_id);

create table if not exists enigma.inbound_wa
(
    id             bigint auto_increment
    primary key,
    client_id      int                                    not null,
    type           varchar(100)                           not null,
    timestamp_unix int                                    not null,
    timestamp_dt   datetime                               not null,
    text_json      text                                   null,
    location_json  text                                   null,
    image_json     text                                   null,
    document_json  text                                   null,
    status         int          default (0)               null,
    messages       text                                   not null,
    created_at     timestamp    default CURRENT_TIMESTAMP null,
    created_by     varchar(100) default (_utf8mb3'Admin') not null
    );

create table if not exists enigma.minion_status_log
(
    id                    int auto_increment
    primary key,
    outbound_id           int default 0 null,
    channel               varchar(50)   null,
    trx_id                varchar(100)  null,
    sesame_id             varchar(100)  null,
    conversation_id       varchar(100)  null,
    message_id            varchar(100)  null,
    channel_id            varchar(100)  null,
    message_content       text          null,
    direction             varchar(50)   null,
    message_status        varchar(50)   null,
    status                varchar(50)   null,
    `from`                varchar(50)   null,
    `to`                  varchar(50)   null,
    platform              varchar(100)  null,
    sql_server_id         int           null,
    created_at            timestamp     null,
    created_at_sql_server timestamp     null,
    transaction_at        timestamp     null,
    last_updated_at       timestamp     null,
    is_process            int default 0 null
    )
    avg_row_length = 16384;

create table if not exists enigma.ms_status
(
    id      int auto_increment
    primary key,
    channel varchar(50) null,
    status  varchar(50) null,
    hirarki int         null
    )
    avg_row_length = 1489;

create table if not exists enigma.outbound
(
    id                bigint auto_increment
    primary key,
    client_id         int                                    not null,
    channel           varchar(255)                           not null,
    channel_id        varchar(255)                           null,
    `from`            varchar(255)                           not null,
    `to`              varchar(255)                           null,
    cc                varchar(255)                           null,
    bcc               varchar(255)                           null,
    wa_template_id    varchar(255)                           null,
    sesame_id         varchar(255)                           not null,
    conversation_id   varchar(255)                           null,
    timestamp_unix    int                                    not null,
    timestamp_dt      datetime                               not null,
    message_id        varchar(255)                           null,
    sesame_message_id varchar(255)                           null,
    title             varchar(255)                           null,
    type              varchar(255)                           null,
    content           text                                   null,
    type_json         text                                   null,
    request_json      text                                   null,
    response_json     text                                   null,
    status            int          default (0)               null,
    created_at        timestamp    default CURRENT_TIMESTAMP null,
    created_by        varchar(100) default (_utf8mb3'Admin') not null,
    message_plain     text                                   null,
    trx_id            varchar(100)                           null,
    x                 varchar(255)                           null
    );

create index outbound_title
    on enigma.outbound (title);

create index to_id
    on enigma.outbound (`to`);

create index wa_template_id
    on enigma.outbound (wa_template_id);

create table if not exists enigma.outbound_archieve
(
    id                bigint                                   not null
    primary key,
    client_id         int                                      not null,
    channel           varchar(255)                             not null,
    channel_id        varchar(255)                             null,
    `from`            varchar(255)                             not null,
    `to`              varchar(255)                             null,
    cc                varchar(255)                             null,
    bcc               varchar(255)                             null,
    wa_template_id    varchar(255)                             null,
    sesame_id         varchar(255)                             not null,
    conversation_id   varchar(255)                             null,
    timestamp_unix    int                                      not null,
    timestamp_dt      datetime                                 not null,
    message_id        varchar(255)                             null,
    sesame_message_id varchar(255)                             null,
    title             varchar(255)                             null,
    content           text                                     null,
    type              varchar(255)                             null,
    type_json         text                                     null,
    request_json      text                                     null,
    response_json     text                                     null,
    status            int          default 0                   null,
    created_at        timestamp    default CURRENT_TIMESTAMP   null,
    created_by        varchar(100) default '_utf8mb3''Admin''' not null,
    message_plain     text                                     null,
    trx_id            varchar(100)                             null
    )
    avg_row_length = 40788;

create table if not exists enigma.outbound_bz
(
    id                bigint auto_increment
    primary key,
    client_id         int                                    not null,
    channel           varchar(255)                           not null,
    channel_id        varchar(255)                           null,
    `from`            varchar(255)                           not null,
    `to`              varchar(255)                           null,
    cc                varchar(255)                           null,
    bcc               varchar(255)                           null,
    wa_template_id    varchar(255)                           null,
    sesame_id         varchar(255)                           not null,
    conversation_id   varchar(255)                           null,
    timestamp_unix    int                                    not null,
    timestamp_dt      datetime                               not null,
    message_id        varchar(255)                           null,
    sesame_message_id varchar(255)                           null,
    title             varchar(255)                           null,
    content           text                                   null,
    type              varchar(255)                           null,
    type_json         text                                   null,
    request_json      text                                   null,
    response_json     text                                   null,
    status            int          default 0                 null,
    created_at        timestamp    default CURRENT_TIMESTAMP null,
    created_by        varchar(100) default 'Admin'           not null,
    message_plain     text                                   null,
    trx_id            varchar(100)                           null
    )
    row_format = DYNAMIC;

create index created_at
    on enigma.outbound_bz (created_at);

create index outbound_id
    on enigma.outbound_bz (id);

create table if not exists enigma.retry_logs
(
    id              bigint auto_increment
    primary key,
    url             varchar(255)                        not null,
    method          varchar(255)                        not null,
    request_header  text                                not null,
    request_body    text                                not null,
    response_status text                                not null,
    response_header text                                not null,
    response_body   text                                not null,
    created_at      timestamp default CURRENT_TIMESTAMP null,
    api_logs_id     bigint                              not null
    );

create table if not exists enigma.roles
(
    id         bigint auto_increment
    primary key,
    name       varchar(100)                           not null,
    created_at timestamp    default CURRENT_TIMESTAMP null,
    created_by varchar(100) default (_utf8mb4'Admin') not null
    );

create table if not exists enigma.sessions
(
    id           bigint auto_increment
    primary key,
    client_id    bigint                                 null,
    vendor_id    int                                    not null,
    vendor_alias varchar(255)                           not null,
    token        varchar(255)                           not null,
    status       int          default (1)               null,
    expired_at   timestamp    default CURRENT_TIMESTAMP null,
    created_at   timestamp    default CURRENT_TIMESTAMP null,
    created_by   varchar(100) default (_utf8mb3'Admin') not null
    );

create table if not exists enigma.status_conversation
(
    id               bigint auto_increment
    primary key,
    conversation_id  varchar(255)                             null,
    `to`             varchar(50)                              null,
    `from`           varchar(50)                              null,
    channel_id       varchar(255)                             null,
    template_name    varchar(255)                             null,
    status           varchar(255)                             null,
    created_datetime timestamp    default CURRENT_TIMESTAMP   not null,
    created_at       timestamp    default CURRENT_TIMESTAMP   null,
    created_by       varchar(100) default '_utf8mb4''Admin''' not null
    );

create table if not exists enigma.status_delivered
(
    id                  bigint auto_increment
    primary key,
    client_id           int          default (0)               null,
    sesame_message_id   varchar(255) default ''                null,
    sesame_id           varchar(255) default ''                null,
    contact_id          varchar(255) default ''                null,
    message_id          varchar(255) default ''                null,
    conversation_id     varchar(255) default ''                null,
    conversation_status varchar(255) default ''                null,
    channel_id          varchar(255) default ''                null,
    platform            varchar(255) default ''                null,
    channel             varchar(255) default ''                null,
    direction           varchar(255) default ''                null,
    message_status      varchar(255) default ''                null,
    `from`              varchar(255) default ''                null,
    `to`                varchar(255) default ''                null,
    msisdn              varchar(255) default ''                null,
    contact_status      varchar(255) default ''                null,
    type                varchar(255) default (_utf8mb4'')      null,
    request_body        text                                   not null,
    response_body       text                                   null,
    api_log_id          int          default (0)               null,
    message_json        text                                   not null,
    contact_json        text                                   not null,
    conversation_json   text                                   not null,
    status              int          default (2)               null,
    created_at          timestamp    default CURRENT_TIMESTAMP null,
    last_updated_at     timestamp    default CURRENT_TIMESTAMP null
    );

create table if not exists enigma.status_failed
(
    id                  bigint auto_increment
    primary key,
    client_id           int          default (0)               null,
    sesame_message_id   varchar(255) default ''                null,
    sesame_id           varchar(255) default ''                null,
    contact_id          varchar(255) default ''                null,
    message_id          varchar(255) default ''                null,
    conversation_id     varchar(255) default ''                null,
    conversation_status varchar(255) default ''                null,
    channel_id          varchar(255) default ''                null,
    platform            varchar(255) default ''                null,
    channel             varchar(255) default ''                null,
    direction           varchar(255) default ''                null,
    message_status      varchar(255) default ''                null,
    `from`              varchar(255) default ''                null,
    `to`                varchar(255) default ''                null,
    msisdn              varchar(255) default ''                null,
    contact_status      varchar(255) default ''                null,
    type                varchar(255) default (_utf8mb4'')      null,
    request_body        text                                   not null,
    response_body       text                                   null,
    api_log_id          int          default (0)               null,
    message_json        text                                   not null,
    contact_json        text                                   not null,
    conversation_json   text                                   not null,
    status              int          default (99)              null,
    created_at          timestamp    default CURRENT_TIMESTAMP null
    );

create table if not exists enigma.status_log
(
    id               int auto_increment
    primary key,
    flow_id          int          default 0                 null,
    client_id        int                                    null,
    channel          varchar(50)                            null,
    destination      varchar(50)                            null,
    message_id       varchar(255) default ''                null,
    message_template varchar(255)                           null,
    max_orders       int                                    null,
    status           varchar(255) default ''                null,
    status_json      text                                   null,
    transaction_at   timestamp    default CURRENT_TIMESTAMP null,
    last_updated_at  timestamp    default CURRENT_TIMESTAMP null,
    is_process       int          default 0                 null,
    constraint unique_index_message_id_status
    unique (message_id, status),
    constraint unq_status_log_message_id_status
    unique (message_id, status)
    )
    avg_row_length = 16384;

create table if not exists enigma.status_pending
(
    id                  bigint auto_increment
    primary key,
    client_id           int          default (0)               null,
    sesame_message_id   varchar(255) default ''                null,
    sesame_id           varchar(255) default ''                null,
    contact_id          varchar(255) default ''                null,
    message_id          varchar(255) default ''                null,
    conversation_id     varchar(255) default ''                null,
    conversation_status varchar(255) default ''                null,
    channel_id          varchar(255) default ''                null,
    platform            varchar(255) default ''                null,
    channel             varchar(255) default ''                null,
    direction           varchar(255) default ''                null,
    message_status      varchar(255) default ''                null,
    `from`              varchar(255) default ''                null,
    `to`                varchar(255) default ''                null,
    msisdn              varchar(255) default ''                null,
    contact_status      varchar(255) default ''                null,
    type                varchar(255) default (_utf8mb4'')      null,
    request_body        text                                   not null,
    response_body       text                                   null,
    api_log_id          int          default (0)               null,
    message_json        text                                   not null,
    contact_json        text                                   not null,
    conversation_json   text                                   not null,
    status              int          default (0)               null,
    created_at          timestamp    default CURRENT_TIMESTAMP null,
    last_updated_at     timestamp    default CURRENT_TIMESTAMP null
    );

create table if not exists enigma.status_read
(
    id                  bigint auto_increment
    primary key,
    client_id           int          default (0)               null,
    sesame_message_id   varchar(255) default ''                null,
    sesame_id           varchar(255) default ''                null,
    contact_id          varchar(255) default ''                null,
    message_id          varchar(255) default ''                null,
    conversation_id     varchar(255) default ''                null,
    conversation_status varchar(255) default ''                null,
    channel_id          varchar(255) default ''                null,
    platform            varchar(255) default ''                null,
    channel             varchar(255) default ''                null,
    direction           varchar(255) default ''                null,
    message_status      varchar(255) default ''                null,
    `from`              varchar(255) default ''                null,
    `to`                varchar(255) default ''                null,
    msisdn              varchar(255) default ''                null,
    contact_status      varchar(255) default ''                null,
    type                varchar(255) default (_utf8mb4'')      null,
    request_body        text                                   not null,
    response_body       text                                   null,
    api_log_id          int          default (0)               null,
    message_json        text                                   not null,
    contact_json        text                                   not null,
    conversation_json   text                                   not null,
    status              int          default (3)               null,
    created_at          timestamp    default CURRENT_TIMESTAMP null
    );

create table if not exists enigma.status_rejected
(
    id                  bigint auto_increment
    primary key,
    client_id           int          default (0)               null,
    sesame_message_id   varchar(255) default ''                null,
    sesame_id           varchar(255) default ''                null,
    contact_id          varchar(255) default ''                null,
    message_id          varchar(255) default ''                null,
    conversation_id     varchar(255) default ''                null,
    conversation_status varchar(255) default ''                null,
    channel_id          varchar(255) default ''                null,
    platform            varchar(255) default ''                null,
    channel             varchar(255) default ''                null,
    direction           varchar(255) default ''                null,
    message_status      varchar(255) default ''                null,
    `from`              varchar(255) default ''                null,
    `to`                varchar(255) default ''                null,
    msisdn              varchar(255) default ''                null,
    contact_status      varchar(255) default ''                null,
    type                varchar(255) default (_utf8mb4'')      null,
    request_body        text                                   not null,
    response_body       text                                   null,
    api_log_id          int          default (0)               null,
    message_json        text                                   not null,
    contact_json        text                                   not null,
    conversation_json   text                                   not null,
    status              int          default (98)              null,
    created_at          timestamp    default CURRENT_TIMESTAMP null
    );

create table if not exists enigma.status_sent
(
    id                  bigint auto_increment
    primary key,
    client_id           int          default (0)               null,
    sesame_message_id   varchar(255) default ''                null,
    sesame_id           varchar(255) default ''                null,
    contact_id          varchar(255) default ''                null,
    message_id          varchar(255) default ''                null,
    conversation_id     varchar(255) default ''                null,
    conversation_status varchar(255) default ''                null,
    channel_id          varchar(255) default ''                null,
    platform            varchar(255) default ''                null,
    channel             varchar(255) default ''                null,
    direction           varchar(255) default ''                null,
    message_status      varchar(255) default ''                null,
    `from`              varchar(255) default ''                null,
    `to`                varchar(255) default ''                null,
    msisdn              varchar(255) default ''                null,
    contact_status      varchar(255) default ''                null,
    type                varchar(255) default (_utf8mb4'')      null,
    request_body        text                                   not null,
    response_body       text                                   null,
    api_log_id          int          default (0)               null,
    message_json        text                                   not null,
    contact_json        text                                   not null,
    conversation_json   text                                   not null,
    status              int          default (1)               null,
    created_at          timestamp    default CURRENT_TIMESTAMP null,
    last_updated_at     timestamp    default CURRENT_TIMESTAMP null
    );

create table if not exists enigma.temporary_table_data_email_blast
(
    id                  varchar(255) null,
    email_verified_at   varchar(255) null,
    created_at          varchar(255) null,
    status              varchar(255) null,
    source_sub_detail   varchar(255) null,
    source_detail       varchar(255) null,
    flag_sub_detail     varchar(255) null,
    Category            varchar(255) null,
    `source detail new` varchar(255) null,
    name                varchar(255) null,
    email               varchar(255) null,
    phone               varchar(255) null,
    reg_source          varchar(255) null,
    login_at            varchar(255) null,
    source              varchar(255) null,
    status_login        varchar(255) null,
    status_edm          varchar(255) null,
    edm_status_date     varchar(255) null
    );

create table if not exists enigma.user_roles
(
    id         bigint auto_increment
    primary key,
    user_id    bigint                                 not null,
    role_id    bigint                                 not null,
    status     int          default (1)               null,
    created_at timestamp    default CURRENT_TIMESTAMP null,
    created_by varchar(100) default (_utf8mb4'Admin') not null,
    constraint UQ_ALL
    unique (user_id, role_id)
    );

create table if not exists enigma.users
(
    id         bigint auto_increment
    primary key,
    username   varchar(255)                           not null,
    password   varchar(255)                           not null,
    client_id  varchar(255)                           not null,
    hash_key   varchar(255) default ''                null,
    status     int          default (1)               null,
    created_at timestamp    default CURRENT_TIMESTAMP null,
    created_by varchar(100) default (_utf8mb3'Admin') not null,
    constraint username
    unique (username)
    );

create table if not exists enigma.vendor
(
    id            bigint auto_increment
    primary key,
    vendor_alias  varchar(255)                           not null,
    url           varchar(255)                           not null,
    uri_login     varchar(255)                           not null,
    username      varchar(255)                           not null,
    password      varchar(255)                           not null,
    method        varchar(100)                           null,
    auth_type     varchar(255)                           null,
    header_prefix varchar(255)                           null,
    access_token  varchar(255)                           null,
    consumer_key  varchar(255)                           null,
    status        int          default (1)               null,
    created_at    timestamp    default CURRENT_TIMESTAMP null,
    created_by    varchar(100) default (_utf8mb4'Admin') not null
    );

create table if not exists enigma.vendor_services
(
    id            bigint auto_increment
    primary key,
    vendor_id     int                                    not null,
    vendor_alias  varchar(255)                           not null,
    service_name  varchar(255)                           not null,
    uri           varchar(255)                           not null,
    method        varchar(100)                           not null,
    header_prefix varchar(255)                           null,
    status        int          default (1)               null,
    created_at    timestamp    default CURRENT_TIMESTAMP null,
    created_by    varchar(100) default (_utf8mb4'Admin') not null,
    from_sender   varchar(100) default ''                null,
    reply_to      varchar(100) default ''                null,
    constraint vendor_id
    unique (vendor_id, vendor_alias, service_name, from_sender)
    );

create table if not exists enigma.webhooks
(
    id                 bigint auto_increment
    primary key,
    client_id          bigint                                 not null,
    client_code        varchar(255)                           not null,
    trx_type           varchar(100) default (_utf8mb4'')      not null,
    url                text                                   not null,
    method             varchar(50)                            not null,
    header_prefix      varchar(255) default ''                null,
    token              text                                   not null,
    events             text                                   not null,
    expected_http_code int                                    not null,
    retry              int                                    not null,
    timeout            int                                    not null,
    status             int          default 1                 null,
    created_at         timestamp    default CURRENT_TIMESTAMP null,
    created_by         varchar(100) default 'Admin'           not null
    )
    avg_row_length = 8192;

create table if not exists enigma.z_cek_customer_fair2024
(
    email        varchar(255) null,
    id           int          null,
    name         varchar(255) null,
    luna_cust_id int          null
    );

