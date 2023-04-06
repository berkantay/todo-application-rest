CREATE TABLE "todos"(
    "id" bigserial PRIMARY KEY,
    "description" varchar NOT NULL,
    "deadline" varchar NOT NULL,
    "priority" varchar NOT NULL
)