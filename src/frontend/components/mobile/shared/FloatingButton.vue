<script setup lang="ts">
  const sideNavRef: any = ref(null);

  const showNav = (action: string) => {
    if (sideNavRef.value) {
      const el = sideNavRef.value.classList;

      if (action === 'show') {
        el.remove('hidden');
        return;
      }

      el.add('hidden');
      return;
    }
  };
</script>

<template>
  <div class="">
    <button
      @click="showNav('show')"
      class="fixed left-0 top-1/2 -translate-y-1/2 px-4 py-2 bg-indigo-500 text-white"
    >
      <i class="pi pi-bars"></i>
    </button>

    <div
      ref="sideNavRef"
      class="fixed inset-0 bg-gray-800 bg-opacity-50 z-50 hidden"
    >
      <div class="absolute left-0 top-0 h-full w-64 bg-white shadow-lg p-4">
        <ul>
          <li><a href="#" class="block py-2">Home</a></li>
          <li><a href="#" class="block py-2">About</a></li>
          <li><a href="#" class="block py-2">Services</a></li>
          <li><a href="#" class="block py-2">Contact</a></li>
        </ul>
      </div>

      <button
        @click="showNav('hide')"
        class="fixed left-64 top-1/2 -translate-y-1/2 px-4 py-2 bg-indigo-500 shadow-lg text-white"
      >
        <i class="pi pi-bars"></i>
      </button>
    </div>
  </div>
</template>
