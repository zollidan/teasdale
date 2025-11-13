CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "keycloak_id" varchar(255) UNIQUE NOT NULL,
  "avatar_url" text,
  "bio" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "artists" (
  "id" uuid PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "bio" text,
  "image_url" text,
  "country" varchar(100),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "genres" (
  "id" uuid PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL,
  "description" text,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "albums" (
  "id" uuid PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "artist_id" uuid NOT NULL,
  "cover_url" text,
  "release_date" date,
  "description" text,
  "album_type" varchar(50) DEFAULT 'album',
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "tracks" (
  "id" uuid PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "album_id" uuid NOT NULL,
  "duration" integer,
  "track_number" integer,
  "file_url" text NOT NULL,
  "plays_count" integer DEFAULT 0,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "track_genres" (
  "track_id" uuid NOT NULL,
  "genre_id" uuid NOT NULL
);

CREATE TABLE "likes" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "track_id" uuid NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "track_id" uuid NOT NULL,
  "content" text NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "reviews" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "album_id" uuid NOT NULL,
  "rating" integer NOT NULL,
  "title" varchar(255),
  "content" text NOT NULL,
  "is_editor" boolean DEFAULT false,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "news" (
  "id" uuid PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "content" text NOT NULL,
  "author_id" uuid NOT NULL,
  "image_url" text,
  "published_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "users" ("keycloak_id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "artists" ("name");

CREATE INDEX ON "genres" ("name");

CREATE INDEX ON "albums" ("artist_id");

CREATE INDEX ON "albums" ("title");

CREATE INDEX ON "albums" ("release_date");

CREATE INDEX ON "albums" ("album_type");

CREATE INDEX ON "tracks" ("album_id");

CREATE INDEX ON "tracks" ("title");

CREATE INDEX ON "track_genres" ("track_id", "genre_id");

CREATE UNIQUE INDEX ON "likes" ("user_id", "track_id");

CREATE INDEX ON "likes" ("user_id");

CREATE INDEX ON "likes" ("track_id");

CREATE INDEX ON "comments" ("user_id");

CREATE INDEX ON "comments" ("track_id");

CREATE INDEX ON "comments" ("created_at");

CREATE UNIQUE INDEX ON "reviews" ("user_id", "album_id");

CREATE INDEX ON "reviews" ("album_id");

CREATE INDEX ON "reviews" ("user_id");

CREATE INDEX ON "reviews" ("is_editor");

CREATE INDEX ON "news" ("author_id");

CREATE INDEX ON "news" ("published_at");

COMMENT ON COLUMN "albums"."album_type" IS 'album, single, ep';

COMMENT ON COLUMN "tracks"."duration" IS 'продолжительность в секундах';

COMMENT ON COLUMN "reviews"."rating" IS 'от 1 до 5';

COMMENT ON COLUMN "reviews"."is_editor" IS 'редакторский обзор';

COMMENT ON COLUMN "news"."author_id" IS 'editor user_id';

ALTER TABLE "albums" ADD FOREIGN KEY ("artist_id") REFERENCES "artists" ("id") ON DELETE CASCADE;

ALTER TABLE "tracks" ADD FOREIGN KEY ("album_id") REFERENCES "albums" ("id") ON DELETE CASCADE;

ALTER TABLE "track_genres" ADD FOREIGN KEY ("track_id") REFERENCES "tracks" ("id") ON DELETE CASCADE;

ALTER TABLE "track_genres" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("id") ON DELETE CASCADE;

ALTER TABLE "likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "likes" ADD FOREIGN KEY ("track_id") REFERENCES "tracks" ("id") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("track_id") REFERENCES "tracks" ("id") ON DELETE CASCADE;

ALTER TABLE "reviews" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "reviews" ADD FOREIGN KEY ("album_id") REFERENCES "albums" ("id") ON DELETE CASCADE;

ALTER TABLE "news" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id") ON DELETE SET NULL;
