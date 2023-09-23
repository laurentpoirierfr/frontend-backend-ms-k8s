CREATE TABLE "menu_items" (
  "id" integer PRIMARY KEY,
  "parent_id" integer,
  "title" varchar,
  "url" varchar,
  "target" varchar,
  "profiles" integer
);

CREATE TABLE "profiles" (
  "id" integer PRIMARY KEY,
  "key" integer,
  "title" varchar,
  "description" varchar
);

ALTER TABLE "menu_items" ADD FOREIGN KEY ("parent_id") REFERENCES "menu_items" ("id");
