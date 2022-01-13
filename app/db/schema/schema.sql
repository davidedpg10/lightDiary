CREATE TABLE "entries" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar,
  "message" varchar,
  "created_at" timestamptz default current_timestamp
);
