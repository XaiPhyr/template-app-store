<script setup lang="ts">
  import type { SliderSlideEndEvent } from 'primevue';
  import { readProducts } from '~/composables/api/products';

  useHead({
    title: 'Shop',
  });

  const route = useRoute();
  const router = useRouter();

  const display = ref('grid');
  const keyword = ref('');
  const isLoading: Ref<boolean, boolean> = ref(false);
  const openDrawer = ref(false);

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

  const dateFrom = ref('');
  const dateTo = ref('');

  const selectedCategory: Ref<string, number> | any = ref('');

  onMounted(() => {
    const { category }: any = route.query;

    if (category) {
      selectedCategory.value = category;
    }

    loadProducts();
  });

  const loadProducts = async () => {
    products.value = [];
    isLoading.value = true;

    const params: any = {
      page: productsPage.value,
      size: productsSize.value,
    };

    if (keyword.value !== '') {
      params.q = keyword.value;
    }

    const qExtItems = [];
    if (adjustedPriceRange.value) {
      qExtItems.push(`price=${adjustedPriceRange.value}`);
    }

    if (selectedCategory.value) {
      qExtItems.push(`category=${selectedCategory.value}`);
    }

    if (dateFrom.value && dateTo.value) {
      qExtItems.push(`created_at=${dateFrom.value}|${dateTo.value}`);
    }

    if (qExtItems.length > 0) {
      params.qExt = qExtItems.join('&');
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
      params.sort = sortItems.join(',');
    }

    const data = await readProducts(params);

    if (data) {
      const { results, min_price, max_price, total } = data;
      if (!results) {
        productsPage.value = 1;
        productsTotalSize.value = 1;
        productsTotalPage.value = 1;
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

  const onSortBy = () => {
    productsPage.value = 1;
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

    dateFrom.value = '';
    dateTo.value = '';

    loadProducts();
  };

  watch(
    () => dateFrom.value,
    (value) => {
      const unixFrom = unixDateTime(value);
      const unixTo = unixDateTime(dateTo.value);

      if (unixFrom <= unixTo) {
        loadProducts();
      }
    }
  );

  watch(
    () => dateTo.value,
    (value) => {
      const unixFrom = unixDateTime(dateFrom.value);
      const unixTo = unixDateTime(value);

      if (unixFrom <= unixTo) {
        loadProducts();
      }
    }
  );
</script>

<template>
  <div class="xl:p-5">
    <ClientOnly>
      <div class="flex">
        <div class="py-4 pl-4 hidden 2xl:block">
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
            @selected-sort-by-name="onSortBy"
            @selected-sort-by-date="onSortBy"
            @selected-sort-by-price="onSortBy"
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

          <WidgetFilter
            class="mb-4"
            filter="date"
            v-model:date-from="dateFrom"
            v-model:date-to="dateTo"
          />

          <button
            @click="onResetFilter"
            class="bg-slate-500 hover:bg-slate-600 py-2 w-full text-white active:bg-slate-700 active:scale-105"
          >
            Reset Filter
          </button>
        </div>

        <div class="p-4 w-full">
          <div class="mb-4">
            <div class="flex gap-4 justify-between">
              <div class="relative w-full">
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

              <button
                @click="openDrawer = true"
                class="border border-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 hover:text-white flex justify-center items-center px-3 block lg:hidden"
              >
                <i class="pi pi-shopping-cart"></i>
              </button>
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
              display === 'grid' ? 'xl:grid-cols-3' : 'xl:grid-cols-1'
            } gap-4`"
            v-if="products.length > 0"
          >
            <div class="" v-for="(item, index) in products" :key="index">
              <SharedCard :display="display" :item="item" />
            </div>
          </div>

          <div
            :class="`grid md:grid-cols-3 ${
              display === 'grid' ? 'xl:grid-cols-3' : 'xl:grid-cols-1'
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
          <WebCart class="w-80" />
        </div>
      </div>

      <transition name="fade">
        <div
          v-if="openDrawer"
          class="fixed inset-0 bg-black bg-opacity-10 backdrop-blur-sm"
          @click="openDrawer = false"
        ></div>
      </transition>

      <div
        class="fixed bottom-0 h-1/2 w-full bg-white p-5 transform transition-transform duration-500 overflow-scroll"
        :class="openDrawer ? 'translate-y-0' : 'translate-y-full'"
      >
        <div class="flex justify-end">
          <button @click="openDrawer = false" class="text-gray-500">
            <i class="pi pi-times" />
          </button>
        </div>
        <div class="mt-3">
          <WebCart class="w-full" />
        </div>
      </div>
    </ClientOnly>
  </div>
</template>
