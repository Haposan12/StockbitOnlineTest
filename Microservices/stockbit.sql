DROP TABLE IF EXISTS `search_log`;

CREATE TABLE `search_log` (
                              `id` character varying NOT NULL,
                              `data` jsonb NOT NULL ,
                              `created_at` timestamp without time zone ,
                              `updated_at` timestamp without time zone ,
                              `deleted_at` timestamp without time zone ,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
