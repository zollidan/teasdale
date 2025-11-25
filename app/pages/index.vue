<script setup lang="ts">
interface Album {
  id: number;
  title: string;
  artistName: string;
  coverUrl: string;
  rating: number;
  itunesArtistId: number;
  artworkPreviewUrl: string;
  artworkFullUrl: string;
  explicitness: string;
  trackCount: number;
  copyright: string;
  country: string;
  releasedDate: string;
  genre: string;
}

interface AlbumResponse {
  error?: string;
  albums: Array<Album>;
}

const albumName = ref("");
const artistName = ref("");
const isLoading = ref(false);
const message = ref("");

const { data, refresh } = await useFetch<AlbumResponse>("/api/albums");

const addAlbum = async () => {
  if (!albumName.value || !artistName.value) {
    message.value = "Please provide both album and artist names.";
    return;
  }

  isLoading.value = true;
  message.value = "";

  try {
    const itunesAlbum = await getItunesAlbum(albumName.value, artistName.value);

    await useFetch("/api/albums", {
      method: "POST",
      body: JSON.stringify(itunesAlbum),
    });

    await refresh();
    message.value = "Альбом успешно добавлен!";
    albumName.value = "";
    artistName.value = "";
  } catch (error) {
    console.error("Error adding album:", error);
    message.value = "Ошибка при добавлении.";
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <div
    class="min-h-screen bg-neutral-900 text-gray-100 p-6 font-sans selection:bg-indigo-500 selection:text-white"
  >
    <header class="mb-8 flex justify-between items-center">
      <h1 class="text-3xl font-bold tracking-tight text-white">
        Моя Коллекция
      </h1>
      <div
        v-if="message"
        :class="{
          'text-green-400': !message.includes('Ошибка'),
          'text-red-400': message.includes('Ошибка'),
        }"
        class="text-sm font-medium animate-pulse"
      >
        {{ message }}
      </div>
    </header>

    <div
      v-if="data?.error"
      class="rounded-lg bg-red-900/50 p-4 text-red-200 border border-red-700"
    >
      {{ data.error }}
    </div>

    <div
      v-else
      class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-6 mb-12"
    >
      <a
        v-for="album in data?.albums"
        :key="album.id"
        href="#"
        class="group relative block aspect-square overflow-hidden rounded-xl bg-neutral-800 shadow-xl ring-1 ring-white/10 transition-all hover:ring-indigo-500/50 focus:outline-none focus:ring-2 focus:ring-indigo-500"
      >
        <img
          :src="album.artworkFullUrl || album.coverUrl"
          :alt="album.title"
          class="h-full w-full object-cover transition-transform duration-500 will-change-transform group-hover:scale-110 group-hover:blur-[2px]"
          loading="lazy"
        />

        <div
          class="absolute inset-0 flex flex-col justify-end bg-gradient-to-t from-black/90 via-black/60 to-transparent p-4 opacity-0 transition-opacity duration-300 group-hover:opacity-100"
        >
          <div class="mb-2 flex text-yellow-400">
            <span v-for="i in 5" :key="i" class="text-xs">
              {{ i <= Math.round(album.rating || 0) ? "★" : "☆" }}
            </span>
          </div>

          <h3 class="text-lg font-bold leading-tight text-white drop-shadow-md">
            {{ album.title }}
          </h3>
          <p class="text-sm font-medium text-indigo-300">
            {{ album.artistName }}
          </p>

          <div
            class="mt-2 flex flex-wrap gap-2 text-[10px] text-gray-400 uppercase tracking-wider"
          >
            <span class="rounded bg-white/10 px-1.5 py-0.5">{{
              album.genre
            }}</span>
            <span>{{ new Date(album.releasedDate).getFullYear() }}</span>
            <span v-if="album.country" class="hidden sm:inline"
              >• {{ album.country }}</span
            >
          </div>

          <div class="mt-1 text-[9px] text-gray-500 truncate">
            {{ album.copyright }}
          </div>
        </div>
      </a>
    </div>

    <div
      class="mx-auto w-full max-w-2xl rounded-2xl border border-white/10 bg-white/5 p-1 backdrop-blur-md sticky bottom-6 shadow-2xl"
    >
      <div class="flex flex-col sm:flex-row gap-2 p-2">
        <input
          type="text"
          placeholder="Название альбома"
          v-model="albumName"
          class="flex-1 rounded-xl bg-transparent px-4 py-3 text-white placeholder-gray-500 focus:bg-white/10 focus:outline-none transition-colors"
          @keyup.enter="addAlbum"
        />
        <div class="h-px w-full bg-white/10 sm:h-auto sm:w-px"></div>
        <input
          type="text"
          placeholder="Артист"
          v-model="artistName"
          class="flex-1 rounded-xl bg-transparent px-4 py-3 text-white placeholder-gray-500 focus:bg-white/10 focus:outline-none transition-colors"
          @keyup.enter="addAlbum"
        />
        <button
          @click="addAlbum"
          :disabled="isLoading"
          class="rounded-xl bg-indigo-600 px-6 py-3 font-semibold text-white transition-all hover:bg-indigo-500 active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed sm:w-auto w-full"
        >
          <span v-if="isLoading" class="animate-spin inline-block mr-2">⟳</span>
          {{ isLoading ? "Загрузка..." : "Добавить" }}
        </button>
      </div>
    </div>
  </div>
</template>
