
-- +migrate Up
CREATE TABLE wht (
  id bigserial NOT NULL,
  record_date timestamp NOT NULL UNIQUE,
  path varchar(1024) NOT NULL,
  created_by varchar(256),
  created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
  updated_by varchar(256),
  updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE wht;
