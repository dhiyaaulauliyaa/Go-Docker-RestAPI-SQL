CREATE TYPE "contact_types" AS ENUM ('instagram', 'facebook', 'whatsapp');


CREATE TABLE
  "users" (
    "id" SERIAL PRIMARY KEY,
    "username" varchar NOT NULL,
    "name" varchar NOT NULL,
    "phone" int NOT NULL,
    "gender" int NOT NULL,
    "age" int NOT NULL,
    "avatar" text,
    "created_at" timestamp NOT NULL DEFAULT (now())
  );


CREATE TABLE
  "masjids" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "address" varchar NOT NULL,
    "city" varchar NOT NULL,
    "coordinate" varchar NOT NULL,
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
  "events" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "venue" varchar NOT NULL,
    "masjid" int,
    "date" timestamp NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now())
  );


CREATE TABLE
  "events_ustadzs" ("event_id" int NOT NULL, "ustadz_id" int NOT NULL);


CREATE TABLE
  "user_events" ("user_id" int NOT NULL, "event_id" int NOT NULL);


CREATE TABLE
  "user_favorited_events" ("user_id" int NOT NULL, "event_id" int NOT NULL);


CREATE TABLE
  "user_favorited_ustadzs" ("user_id" int NOT NULL, "ustadz_id" int NOT NULL);


CREATE TABLE
  "user_favorited_masjids" ("user_id" int NOT NULL, "masjid_id" int NOT NULL);


CREATE TABLE
  "contacts" (
    "ustadz_id" int NOT NULL,
    "contacts_type" contact_types NOT NULL,
    "value" text NOT NULL
  );


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
  "user_events"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");


ALTER TABLE
  "user_events"
ADD
  FOREIGN KEY ("event_id") REFERENCES "events" ("id");


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
  "user_favorited_masjids"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");


ALTER TABLE
  "user_favorited_masjids"
ADD
  FOREIGN KEY ("masjid_id") REFERENCES "masjids" ("id");


ALTER TABLE
  "contacts"
ADD
  FOREIGN KEY ("ustadz_id") REFERENCES "ustadzs" ("id");