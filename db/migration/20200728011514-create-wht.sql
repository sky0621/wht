
-- +migrate Up
CREATE TABLE wht (
  id bigserial NOT NULL,
  record_date timestamp NOT NULL UNIQUE,
  title varchar(256),
  created_by varchar(256),
  created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
  updated_by varchar(256),
  updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE content_text (
  id bigserial NOT NULL,
  wht_id bigint REFERENCES wht(id) NOT NULL,
  name varchar(256),
  text text NOT NULL,
  created_by varchar(256),
  created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
  updated_by varchar(256),
  updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE content_image (
  id bigserial NOT NULL,
  wht_id bigint REFERENCES wht(id) NOT NULL,
  name varchar(256),
  path varchar(1024) NOT NULL,
  created_by varchar(256),
  created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
  updated_by varchar(256),
  updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE content_image;
DROP TABLE content_text;
DROP TABLE wht;
