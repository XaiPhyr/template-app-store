<script setup lang="ts">
  const showButton = ref(false);

  onMounted(() => {
    window.addEventListener('scroll', (event) => {
      const { scrollY } = window;

      if (scrollY >= 500) {
        showButton.value = true;
        return;
      }

      showButton.value = false;
    });
  });

  const scrollToTop = () => {
    window.scrollTo({
      top: 0,
      behavior: 'smooth',
    });
  };
</script>

<template>
  <div class="" v-if="showButton">
    <div class="fixed bottom-4 right-4 text-black">
      <button @click="scrollToTop" class="bg-indigo-500 text-white py-2 px-4">
        <i class="pi pi-arrow-up"></i>
      </button>
    </div>
  </div>
</template>
