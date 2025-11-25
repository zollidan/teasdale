import { integer, pgTable, varchar, real } from "drizzle-orm/pg-core";

export const albumTable = pgTable("albums", {
  id: integer().primaryKey().generatedAlwaysAsIdentity(),
  title: varchar({ length: 255 }).notNull(),
  artistName: varchar({ length: 255 }).notNull(),
  coverUrl: varchar({ length: 512 }),
  rating: real(),
  itunesArtistId: integer(),
  artworkPreviewUrl: varchar({ length: 512 }),
  artworkFullUrl: varchar({ length: 512 }),
  explicitness: varchar({ length: 255 }),
  trackCount: integer(),
  copyright: varchar({ length: 512 }),
  country: varchar({ length: 255 }),
  releasedDate: varchar({ length: 255 }),
  genre: varchar({ length: 255 }),
});
