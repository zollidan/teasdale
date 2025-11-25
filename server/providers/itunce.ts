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

export const getAlbum = async (albumName: string, artistName: string) => {
  const params = new URLSearchParams({
    term: `${albumName} ${artistName}`,
    media: "music",
    entity: "album",
    country: "US",
    limit: "1",
  });

  const url = `${ITUNES_API_URL}?${params.toString()}`;

  const res = await fetch(url);
  const data: ItunesAlbumResponse = await res.json();
  return data.results;
};
