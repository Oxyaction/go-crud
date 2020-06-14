-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS "user" (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR (255) DEFAULT NULL,
  email VARCHAR (255) NOT NULL ,
  password VARCHAR (255) NOT NULL,
  createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  CONSTRAINT user_email UNIQUE(email)
);

---- create above / drop below ----

drop table "user";
