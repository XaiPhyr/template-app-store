<script setup lang="ts">
  import type { SliderSlideEndEvent } from 'primevue';
  import { readCategories } from '~/composables/api/categories';
  import initApi from '~/api/instance';

  useHead({
    title: 'Shop',
  });

  const router = useRouter();

  const display = ref('grid');
  const keyword = ref('');
  const isLoading: Ref<boolean, boolean> = ref(false);

  const products: any = ref([]);
  const productsPage = ref(1);
  const productsTotalPage = ref(0);
  const productsSize = ref(10);
  const productsTotalSize = ref(0);

  const priceRange: any = ref([]);
  const adjustedPriceRange = ref('');
  const originalPriceRange: any = ref([]);

  const sortByName = ref('');
  const sortByPrice = ref('');
  const sortByDate = ref('');

  const selectedCategory: Ref<string, number> | any = ref('');

  onMounted(() => {
    loadProducts();
  });

  const loadProducts = async () => {
    products.value = [];
    isLoading.value = true;
    const params: any = {
      url: '/products',
      method: 'read',
      params: { page: productsPage.value, size: productsSize.value },
    };

    if (keyword.value !== '') {
      params.params.q = keyword.value;
    }

    if (adjustedPriceRange.value) {
      params.params.qExt = `price=${adjustedPriceRange.value}`;
    }

    if (selectedCategory.value.length > 0) {
      params.params.qExt = `category=${selectedCategory.value}`;
    }

    const sortItems = [];
    if (sortByName.value) {
      sortItems.push(sortByName.value);
    }

    if (sortByDate.value) {
      sortItems.push(sortByDate.value);
    }

    if (sortByPrice.value) {
      sortItems.push(sortByPrice.value);
    }

    if (sortItems.length > 0) {
      params.params.sort = sortItems.join(',');
    }

    try {
      const { data } = await initApi(params);

      if (!data) {
        productsTotalSize.value = 1;
        productsTotalPage.value = 1;
      }

      if (data) {
        const { results, min_price, max_price, total } = data;
        if (!results) {
          isLoading.value = false;
          return;
        }

        if (results) {
          products.value = results;
        }

        if (priceRange.value.length === 0) {
          priceRange.value = [min_price, max_price];
        }

        originalPriceRange.value = [min_price, max_price];
        productsTotalSize.value = total;
        productsTotalPage.value = Math.ceil(total / productsSize.value);
      }

      isLoading.value = false;
    } catch (error) {
      isLoading.value = false;
    }
  };

  const onSelectCategory = (item: any) => {
    if (item === '') {
      router.replace(`/shop`);
      selectedCategory.value = '';
      loadProducts();
      return;
    }

    selectedCategory.value = item;

    router.replace(`/shop?category=${encodeURIComponent(item)}`);
    productsPage.value = 1;

    loadProducts();
  };

  const onSelectproductsSize = (item: number) => {
    productsSize.value = item;
    productsPage.value = 1;

    loadProducts();
  };

  const onSelectPiceRange = (item: SliderSlideEndEvent) => {
    const { value }: any = item;
    adjustedPriceRange.value = value.join('-');

    loadProducts();
  };

  const onSelectPage = (page: string) => {
    if (page === 'next') {
      if (productsPage.value < productsTotalPage.value) {
        productsPage.value += 1;
        loadProducts();
        return;
      }
    }

    if (productsPage.value > 1) {
      productsPage.value -= 1;
      loadProducts();
      return;
    }
  };

  const onResetFilter = () => {
    router.replace(`/shop`);

    sortByName.value = '';
    sortByPrice.value = '';
    sortByDate.value = '';

    priceRange.value = originalPriceRange.value;
    adjustedPriceRange.value = '';

    selectedCategory.value = '';

    loadProducts();
  };
</script>

<template>
  <div class="xl:p-5">
    <ClientOnly>
      <div class="flex">
        <div class="py-4 pl-4 hidden xl:block">
          <WidgetOptions
            class="mb-4"
            @selected-row-size="(v) => onSelectproductsSize(v)"
            @selected-display="(v) => (display = v)"
          />

          <WidgetFilter
            class="mb-4"
            filter="sort"
            v-model:sort-by-name="sortByName"
            v-model:sort-by-price="sortByPrice"
            v-model:sort-by-date="sortByDate"
            @selected-sort-by-name="loadProducts"
            @selected-sort-by-date="loadProducts"
            @selected-sort-by-price="loadProducts"
          />

          <WidgetFilter
            class="mb-4"
            filter="categories"
            @selected-category="(v) => onSelectCategory(v)"
          />

          <WidgetFilter
            class="mb-4"
            filter="price"
            @selected-price-range="(v) => onSelectPiceRange(v)"
            :price-range="priceRange"
          />

          <button
            @click="onResetFilter"
            class="bg-slate-500 hover:bg-slate-600 py-1 w-full text-white active:bg-slate-700 active:scale-105"
          >
            Reset Filter
          </button>
        </div>

        <div class="p-4 w-full">
          <div class="mb-4">
            <div class="relative">
              <span
                class="absolute left-0 top-1/2 transform -translate-y-1/2 px-2 text-gray-500 hover:cursor-pointer"
              >
                <i class="pi pi-search"></i>
              </span>
              <input
                @change="loadProducts"
                v-model="keyword"
                type="text"
                class="pl-8 pr-3 py-2 border w-full focus:outline-none focus:ring-2 focus:ring-indigo-500"
                placeholder="Search"
              />
            </div>
          </div>

          <div class="mb-4">
            <div class="flex justify-center">
              <div class="w-64 flex gap-4 justify-center">
                <div class="">
                  <button
                    @click="onSelectPage('previous')"
                    title="Grid"
                    class="bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 text-white px-4 py-2 flex"
                  >
                    <div class="">
                      <i class="pi pi-chevron-left"></i>
                    </div>
                  </button>
                </div>
                <div class="text-center content-center w-32">
                  <div class="">
                    {{ productsPage }} / {{ productsTotalPage }}
                  </div>
                </div>
                <div class="">
                  <button
                    @click="onSelectPage('next')"
                    title="List"
                    class="bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 text-white px-4 py-2 flex"
                  >
                    <div class="">
                      <i class="pi pi-chevron-right"></i>
                    </div>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <div
            :class="`grid md:grid-cols-3 ${
              display === 'grid' ? 'xl:grid-cols-4' : 'xl:grid-cols-1'
            } gap-4`"
            v-if="products.length > 0"
          >
            <div class="" v-for="(item, index) in products" :key="index">
              <SharedCard :display="display" :item="item" />
            </div>
          </div>

          <div
            :class="`grid md:grid-cols-3 ${
              display === 'grid' ? 'xl:grid-cols-4' : 'xl:grid-cols-1'
            } gap-4`"
            v-else-if="products.length === 0 && isLoading"
          >
            <div class="" v-for="(_, index) in 10" :key="index">
              <SkeletonCard />
            </div>
          </div>

          <div class="text-center" v-else>No Item(s) available</div>
        </div>

        <div class="py-4 pr-4 hidden xl:block">
          <WebCart />
        </div>
      </div>
    </ClientOnly>
  </div>
</template>
