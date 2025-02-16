<script setup lang="ts">
  const sideNavOpen = ref(false);
</script>

<template>
  <div class="">
    <div
      class="fixed bottom-0 left-0 w-full bg-indigo-500 text-white p-4 flex justify-around"
    >
      <a title="Shop" href="#" class="px-4 py-1">
        <i class="pi pi-shop text-xl" />
      </a>
      <a title="Profile" href="#" class="px-4 py-1">
        <i class="pi pi-user text-xl" />
      </a>
      <button
        @click="sideNavOpen = !sideNavOpen"
        title="Menu"
        class="px-4 py-1"
      >
        <i class="pi pi-ellipsis-v text-xl" />
      </button>
    </div>

    <transition name="fade">
      <div
        v-if="sideNavOpen"
        class="fixed inset-0 bg-black bg-opacity-10 backdrop-blur-sm z-40"
        @click="sideNavOpen = false"
      ></div>
    </transition>

    <div
      class="fixed left-0 top-0 h-full w-64 bg-white p-5 transform transition-transform duration-500 overflow-y-auto shadow-lg z-50"
      :class="sideNavOpen ? 'translate-x-0' : '-translate-x-full'"
    >
      <div class="absolute left-0 top-0 h-full w-64 bg-white shadow-lg p-4">
        <SharedNavigations />
      </div>
    </div>
  </div>
</template>
