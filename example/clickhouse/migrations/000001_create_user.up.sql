CREATE TABLE IF NOT EXISTS `user`
(
    `id`         Int64,
    `username`   String,
    `password`   String,
    `created_at` DateTime,
    `updated_at` DateTime,
    `period`     UInt32
) ENGINE = MergeTree()
      PARTITION BY period
      ORDER BY (created_at, period);
