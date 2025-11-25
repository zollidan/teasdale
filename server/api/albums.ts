import { eq } from "drizzle-orm";
import { db } from "~~/db";
import { albumTable } from "~~/db/schema";

export default defineEventHandler(async (event) => {
  const query = getQuery(event);

  if (query.title && query.artist) {
    const album = await db
      .select()
      .from(albumTable)
      .where(eq(albumTable.title, String(query.title)));

    return album;
  } else {
    const albums = await db.select().from(albumTable);

    if (!albums || albums.length === 0) {
      return {
        error: "No albums found",
      };
    }

    return {
      albums,
    };
  }
});
