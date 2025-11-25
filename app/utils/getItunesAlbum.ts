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

type ItunesAlbumResponse = {
  results: ItunesAlbum[];
};

const ITUNES_API_URL = "https://itunes.apple.com/search";

export default async function (albumName: string, artistName: string) {
  const params = new URLSearchParams({
    term: `${albumName} ${artistName}`,
    media: "music",
    entity: "album",
    country: "US",
    limit: "1",
  });

  const url = `${ITUNES_API_URL}?${params}`;

  const res = await fetch(url);
  const data = (await res.json()) as ItunesAlbumResponse;
  return data.results[0];
}
