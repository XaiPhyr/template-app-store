<script setup lang="ts">
  import { readProducts } from '~/composables/api/products';

  useHead({
    title: 'Home',
  });

  const products: any = ref([]);
  const responsiveOptions = ref([
    {
      breakpoint: '1400px',
      numVisible: 2,
      numScroll: 1,
    },
    {
      breakpoint: '1199px',
      numVisible: 3,
      numScroll: 1,
    },
    {
      breakpoint: '767px',
      numVisible: 2,
      numScroll: 1,
    },
    {
      breakpoint: '575px',
      numVisible: 1,
      numScroll: 1,
    },
  ]);

  onMounted(() => {
    loadProducts();
  });

  const loadProducts = async () => {
    const params: any = {
      page: 1,
      size: 15,
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
  <div class="p-4">
    <Carousel
      :value="products"
      :num-visible="4"
      :num-scroll="1"
      :autoplay-interval="3000"
      :responsive-options="responsiveOptions"
      circular
    >
      <template #item="slotProps">
        <SharedCard class="mx-4" display="grid" :item="slotProps.data" />
      </template>
    </Carousel>
  </div>
</template>
