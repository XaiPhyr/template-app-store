<script setup lang="ts">
  import type { SliderSlideEndEvent } from 'primevue';
  import { readCategories } from '~/composables/api/categories';
  import initApi from '~/api/instance';

  useHead({
    title: 'Shop',
  });

  const router = useRouter();

  const items: any = ref([]);
  const itemsCategories: any = ref([]);
  const rowSizeCategories: any = ref(5);
  const totalPageCategories: any = ref(0);
  const priceRange: any = ref([]);
  const originalPriceRange: any = ref([]);
  const searchIcon = resolveComponent('SearchIcon');
  const keyword = ref('');
  const isLoading: Ref<boolean, boolean> = ref(false);
  const selectedCategory: Ref<string, number> | any = ref('');
  const rowPage = ref(1);
  const totalPage = ref(0);
  const rowSize = ref(10);
  const totalSize = ref(0);
  const sortByName = ref('');
  const sortByPrice = ref('');
  const sortByDate = ref('');
  const display = ref('grid');
  const adjustedPriceRange = ref('');

  onMounted(() => {
    loadCategories();
    loadProducts();
  });

  const loadProducts = async () => {
    items.value = [];
    isLoading.value = true;
    const params: any = {
      url: '/products',
      method: 'read',
      params: { page: rowPage.value },
    };

    if (rowSize.value) {
      params.params.size = rowSize.value;
    }

    if (keyword.value !== '') {
      params.params.q = keyword.value;
    }

    if (adjustedPriceRange.value) {
      params.params.qExt = `price=${adjustedPriceRange.value}`;
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

    if (selectedCategory.value.length > 0) {
      params.params.qExt = `category=${selectedCategory.value}`;
    }

    try {
      const { data } = await initApi(params);

      if (data) {
        const { results, min_price, max_price, total } = data;
        if (!results) {
          isLoading.value = false;
          return;
        }

        if (results) {
          items.value = results;
        }

        if (priceRange.value.length === 0) {
          priceRange.value = [min_price, max_price];
        }

        originalPriceRange.value = [min_price, max_price];
        totalSize.value = total;
        totalPage.value = Math.ceil(total / rowSize.value);
      }

      isLoading.value = false;
    } catch (error) {
      isLoading.value = false;
    }
  };

  const loadCategories = async () => {
    const { total, results } = await readCategories({
      size: rowSizeCategories.value,
    });

    itemsCategories.value = results;
    totalPageCategories.value = total;
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

    loadProducts();
  };

  const onSelectRowSize = (item: number) => {
    rowSize.value = item;

    loadProducts();
  };

  const onSelectPiceRange = (item: SliderSlideEndEvent) => {
    const { value }: any = item;
    adjustedPriceRange.value = value.join('-');

    loadProducts();
  };

  const onSeeMoreCategories = () => {
    if (rowSizeCategories.value > totalPageCategories.value) {
      return;
    }

    rowSizeCategories.value += 10;

    loadCategories();
  };

  const pagination = (page: string) => {
    if (page === 'next') {
      if (rowPage.value < totalPage.value) {
        rowPage.value += 1;
        loadProducts();
        return;
      }

      return;
    }

    if (rowPage.value > 1) {
      rowPage.value -= 1;
      loadProducts();
      return;
    }

    return;
  };

  const onResetFilter = () => {
    sortByName.value = '';
    sortByPrice.value = '';
    sortByDate.value = '';
    adjustedPriceRange.value = '';
    priceRange.value = originalPriceRange.value;

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
            @selected-row-size="(v) => onSelectRowSize(v)"
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
            @see-more-categories="(v) => onSeeMoreCategories()"
            :categories="itemsCategories"
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
                <component :is="searchIcon" />
              </span>
              <input
                @change="loadProducts"
                v-model="keyword"
                type="text"
                class="pl-10 pr-3 py-2 border w-full focus:outline-none focus:ring-2 focus:ring-indigo-500"
                placeholder="Search"
              />
            </div>
          </div>

          <div class="mb-4" v-if="items.length !== 0 && !isLoading">
            <div class="flex justify-center">
              <div class="w-64 flex gap-4 justify-center">
                <div class="">
                  <button
                    @click="pagination('previous')"
                    title="Grid"
                    class="bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 text-white px-2 py-1 flex text-center justify-center"
                  >
                    <ChevronLeftIcon /> <span class="pr-2">Back</span>
                  </button>
                </div>
                <div class="text-center content-center w-32">
                  <div class="">{{ rowPage }} / {{ totalPage }}</div>
                </div>
                <div class="">
                  <button
                    @click="pagination('next')"
                    title="List"
                    class="bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 text-white px-2 py-1 flex text-center justify-center"
                  >
                    <span class="pl-2">Next</span> <ChevronRightIcon />
                  </button>
                </div>
              </div>
            </div>
          </div>

          <div
            :class="`grid md:grid-cols-3 ${
              display === 'grid' ? 'xl:grid-cols-4' : 'xl:grid-cols-1'
            } gap-4`"
            v-if="items.length > 0"
          >
            <div class="" v-for="(item, index) in items" :key="index">
              <SharedCard :display="display" :item="item" />
            </div>
          </div>

          <div
            :class="`grid md:grid-cols-3 ${
              display === 'grid' ? 'xl:grid-cols-4' : 'xl:grid-cols-1'
            } gap-4`"
            v-else-if="items.length === 0 && isLoading"
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
