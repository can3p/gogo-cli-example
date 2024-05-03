
-- +migrate Up
CREATE TABLE system_settings (
  id uuid not null primary key,
  registration_open boolean not null
);

INSERT INTO system_settings (id, registration_open) values ('6443fc37-cad5-452f-b6c2-b13460f019e1', true);

-- +migrate Down
DROP TABLE system_settings;
