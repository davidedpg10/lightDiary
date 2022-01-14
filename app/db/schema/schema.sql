CREATE TABLE "entries" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar NOT NULL,
  "message" varchar NOT NULL,
  "mood" varchar[],
  "created_at" timestamptz default current_timestamp
);
