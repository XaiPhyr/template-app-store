<script setup lang="ts">
  const sideNavOpen = ref(false);
</script>

<template>
  <div class="relative">
    <button
      @click="sideNavOpen = !sideNavOpen"
      class="fixed top-1/2 -translate-y-1/2 px-4 py-2 bg-indigo-500 text-white transition-all duration-500 z-50"
      :class="sideNavOpen ? 'left-64' : 'left-0'"
    >
      <i class="pi pi-bars"></i>
    </button>

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
        <ul>
          <li><a href="#" class="block p-2">Home</a></li>
          <li><a href="#" class="block p-2">About</a></li>
          <li><a href="#" class="block p-2">Services</a></li>
          <li><a href="#" class="block p-2">Contact</a></li>
        </ul>
      </div>
    </div>
  </div>
</template>
