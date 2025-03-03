<script setup lang="ts">
  import { readProducts } from '~/composables/api/products';

  useHead({
    title: 'Home',
  });

  const products: any = ref([]);

  onMounted(() => {
    loadProducts();
  });

  const loadProducts = async () => {
    const params: any = {
      page: 1,
      size: 10,
    };

    const data = await readProducts(params);

    if (data) {
      const { results } = data;

      if (results) {
        products.value = results;
      }
    }
  };
</script>

<template>
  <div class="">
    <div class="p-4 hidden md:block">
      <div class="">
        <Carousel
          :value="products"
          :num-visible="5"
          :num-scroll="1"
          :autoplay-interval="3000"
          circular
        >
          <template #item="slotProps">
            <SharedCard class="mx-4" display="grid" :item="slotProps.data" />
          </template>
        </Carousel>
      </div>
    </div>
  </div>
</template>
