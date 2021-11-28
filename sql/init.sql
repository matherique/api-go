create table if not exists users (
  id serial not null primary key,
  name varchar(100) not null,
  email varchar(50) not null, 
  createdAt timestamp not null default CURRENT_TIMESTAMP,
  updatedAT timestamp not null default CURRENT_TIMESTAMP
)
