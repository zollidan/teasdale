import { db } from "~~/db";
import { albumTable } from "~~/db/schema";
import { and, eq } from "drizzle-orm";

interface ItunesAlbum {
  wrapperType: string;
  collectionType: string;
  artistId: number;
  collectionId: number;
  amgArtistId: number;
  artistName: string;
  collectionName: string;
  collectionCensoredName: string;
  artistViewUrl: string;
  collectionViewUrl: string;
  artworkUrl60: string;
  artworkUrl100: string;
  collectionPrice: number;
  collectionExplicitness: string;
  trackCount: number;
  copyright: string;
  country: string;
  currency: string;
  releaseDate: string;
  primaryGenreName: string;
}

export default defineEventHandler(async (event) => {
  const body = await readBody<ItunesAlbum>(event);

  if (!body.collectionName || !body.artistName) {
    setResponseStatus(event, 400);
    return {
      error:
        "Missing required fields: collectionName and artistName are required",
    };
  }

  const existingAlbum = await db
    .select()
    .from(albumTable)
    .where(
      and(
        eq(albumTable.title, body.collectionName),
        eq(albumTable.artistName, body.artistName)
      )
    )
    .limit(1);

  if (existingAlbum && existingAlbum.length > 0) {
    setResponseStatus(event, 200);
    return {
      album: existingAlbum[0],
      message: "Album already exists",
    };
  }

  const album: typeof albumTable.$inferInsert = {
    title: body.collectionName,
    artistName: body.artistName,
    coverUrl: body.artworkUrl100,
    itunesArtistId: body.artistId,
    artworkPreviewUrl: body.artworkUrl100.replace("100x100bb", "1000x1000bb"),
    artworkFullUrl: body.artworkUrl100.replace("100x100bb", "3000x3000bb"),
    explicitness: body.collectionExplicitness,
    trackCount: body.trackCount,
    copyright: body.copyright,
    country: body.country,
    releasedDate: body.releaseDate,
    genre: body.primaryGenreName,
  };

  try {
    const [insertedAlbum] = await db
      .insert(albumTable)
      .values(album)
      .returning();

    setResponseStatus(event, 201);
    return {
      album: insertedAlbum,
      message: "Album created successfully",
    };
  } catch (error) {
    console.error("Error creating album:", error);
    setResponseStatus(event, 500);
    return {
      error: "Failed to create album",
    };
  }
});
