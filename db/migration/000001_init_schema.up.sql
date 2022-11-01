CREATE TYPE "contact_types" AS ENUM (
  'instagram',
  'facebook',
  'whatsapp',
  'phone',
  'email'
);


CREATE TABLE
  "users" (
    "id" SERIAL PRIMARY KEY,
    "username" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL,
    "name" varchar NOT NULL,
    "phone" varchar UNIQUE NOT NULL,
    "email" varchar UNIQUE,
    "gender" int NOT NULL,
    "age" int NOT NULL,
    "avatar" text,
    "created_at" timestamp NOT NULL DEFAULT (now())
  );


CREATE TABLE
  "communities" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "address" varchar NOT NULL,
    "city" varchar NOT NULL,
    "coordinate" varchar,
    "contact_person" varchar NOT NULL,
    "logo" varchar,
    "created_at" timestamp NOT NULL DEFAULT (now())
  );


CREATE TABLE
  "communities_contacts" (
    "community_id" int NOT NULL,
    "contacts_type" contact_types NOT NULL,
    "value" text NOT NULL
  );


CREATE TABLE
  "communities_users" ("community_id" int NOT NULL, "user_id" int NOT NULL);


CREATE TABLE
  "masjids" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "address" varchar NOT NULL,
    "city" varchar NOT NULL,
    "coordinate" varchar NOT NULL,
    "phone" varchar,
    "logo" varchar,
    "created_at" timestamp NOT NULL DEFAULT (now())
  );


CREATE TABLE
  "ustadzs" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "avatar" varchar,
    "gender" int NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now())
  );


CREATE TABLE
  "ustadzs_contacts" (
    "ustadz_id" int NOT NULL,
    "contacts_type" contact_types NOT NULL,
    "value" text NOT NULL
  );


CREATE TABLE
  "events" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "venue" varchar NOT NULL,
    "community" int,
    "masjid" int,
    "date" timestamp NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now())
  );


CREATE TABLE
  "events_ustadzs" ("event_id" int NOT NULL, "ustadz_id" int NOT NULL);


CREATE TABLE
  "events_users" ("event_id" int NOT NULL, "user_id" int NOT NULL);


CREATE TABLE
  "user_favorited_events" ("user_id" int NOT NULL, "event_id" int NOT NULL);


CREATE TABLE
  "user_favorited_ustadzs" ("user_id" int NOT NULL, "ustadz_id" int NOT NULL);


CREATE TABLE
  "user_favorited_communities" ("user_id" int NOT NULL, "community_id" int NOT NULL);


ALTER TABLE
  "communities_contacts"
ADD
  FOREIGN KEY ("community_id") REFERENCES "communities" ("id");


ALTER TABLE
  "communities_users"
ADD
  FOREIGN KEY ("community_id") REFERENCES "communities" ("id");


ALTER TABLE
  "communities_users"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");


ALTER TABLE
  "ustadzs_contacts"
ADD
  FOREIGN KEY ("ustadz_id") REFERENCES "ustadzs" ("id");


ALTER TABLE
  "events"
ADD
  FOREIGN KEY ("community") REFERENCES "communities" ("id");


ALTER TABLE
  "events"
ADD
  FOREIGN KEY ("masjid") REFERENCES "masjids" ("id");


ALTER TABLE
  "events_ustadzs"
ADD
  FOREIGN KEY ("event_id") REFERENCES "events" ("id");


ALTER TABLE
  "events_ustadzs"
ADD
  FOREIGN KEY ("ustadz_id") REFERENCES "ustadzs" ("id");


ALTER TABLE
  "events_users"
ADD
  FOREIGN KEY ("event_id") REFERENCES "events" ("id");


ALTER TABLE
  "events_users"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");


ALTER TABLE
  "user_favorited_events"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");


ALTER TABLE
  "user_favorited_events"
ADD
  FOREIGN KEY ("event_id") REFERENCES "events" ("id");


ALTER TABLE
  "user_favorited_ustadzs"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");


ALTER TABLE
  "user_favorited_ustadzs"
ADD
  FOREIGN KEY ("ustadz_id") REFERENCES "ustadzs" ("id");


ALTER TABLE
  "user_favorited_communities"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");


ALTER TABLE
  "user_favorited_communities"
ADD
  FOREIGN KEY ("community_id") REFERENCES "communities" ("id");