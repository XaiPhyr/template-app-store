<script setup lang="ts">
  import { computed } from 'vue';

  const route = useRoute();

  const props = defineProps({
    filter: {
      type: String,
      default: '',
    },
    categories: {
      type: Array,
      default: () => [],
    },
    priceRange: {
      type: Array,
      default: [0, 0],
    },
  });

  const modelSortByName = defineModel('sortByName');
  const modelSortByPrice = defineModel('sortByPrice');
  const modelSortByDate = defineModel('sortByDate');

  const emits = defineEmits([
    'selectedCategory',
    'selectedPriceRange',
    'selectedSortByName',
    'selectedSortByDate',
    'selectedSortByPrice',
    'seeMoreCategories',
  ]);

  const multiSelectedCategory: any = ref([]);

  const selectCategory = (item: string | number) => {
    const { category }: any = route.query;

    const items = category ? category.split(',') : [];

    if (!items.includes(item.toString())) {
      multiSelectedCategory.value.push(item);
    }

    if (items.includes(item.toString())) {
      const deselectCategory = multiSelectedCategory.value.filter(
        (x: string | number) => {
          return x !== item;
        }
      );

      multiSelectedCategory.value = deselectCategory;
    }

    emits('selectedCategory', multiSelectedCategory.value.join(','));
  };

  const adjustedPriceRange: any = ref([]);

  const priceRange: any = computed({
    get() {
      adjustedPriceRange.value = props.priceRange;
      return props.priceRange;
    },
    set(value) {
      adjustedPriceRange.value = value;
    },
  });

  const minPrice: any = computed(() => {
    return props.priceRange[0];
  });

  const maxPrice: any = computed(() => {
    return props.priceRange[1];
  });

  const displayPriceRange: any = computed(() => {
    const [min, max] = adjustedPriceRange.value;
    return `${formatCurrencies(min)} - ${formatCurrencies(max)}`;
  });

  const categories: any = computed({
    get() {
      return props.categories;
    },
    set(value) {
      //
    },
  });

  const isSelectedCategory = (item: string | any) => {
    if (multiSelectedCategory.value.includes(item)) {
      return 'font-bold';
    }

    return '';
  };
</script>

<template>
  <div class="border w-64">
    <div class="" v-if="props.filter === 'categories'">
      <div class="p-2">
        <div class="font-bold text-center">Categories</div>
      </div>

      <div class="border-t">
        <div class="" v-for="({ name }, index) in categories" :key="index">
          <div class="p-1">
            <div
              class="px-5 py-1 hover:cursor-pointer active:bg-gray-100"
              @click="selectCategory(name)"
            >
              <div :class="`p-1 ${isSelectedCategory(name)}`">
                <i
                  class="pi pi-check text-xs mr-2"
                  v-if="isSelectedCategory(name)"
                ></i>
                {{ name }}
              </div>
            </div>
          </div>
        </div>

        <div class="p-1 text-xs" @click="emits('seeMoreCategories', true)">
          <div class="px-5 py-1 hover:cursor-pointer active:bg-gray-100">
            <div class="p-1 flex justify-end text-blue-500">See more</div>
          </div>
        </div>
      </div>
    </div>

    <div class="" v-if="props.filter === 'price'">
      <div class="p-2">
        <div class="font-bold text-center">Price</div>
      </div>

      <div class="border-t">
        <div class="p-5">
          <Slider
            @slideend="emits('selectedPriceRange', $event)"
            v-model="priceRange"
            :min="minPrice"
            :max="maxPrice"
            range
            class="w-full"
          />
        </div>
        <div class="px-5 py-1">
          <div class="p-1">
            Range:
            <span class="">
              {{ displayPriceRange }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <div class="" v-if="props.filter === 'sort'">
      <div class="p-2">
        <div class="font-bold text-center">Sort By</div>
      </div>

      <div class="border-t">
        <div class="p-5">
          <div class="mb-4">
            <div class="grid grid-cols-4 gap-4">
              <div class="content-center">Name</div>
              <div class="w-full col-span-3">
                <select
                  @change="emits('selectedSortByName', true)"
                  v-model="modelSortByName"
                  placeholder="Sort by Name"
                  class="block appearance-none w-full bg-white border p-2 px-3 leading-tight focus:outline-none focus:ring-2 focus:ring-indigo-500"
                >
                  <option class="" value="name">Name (A-Z)</option>
                  <option class="" value="-name">Name (Z-A)</option>
                </select>
              </div>
            </div>
          </div>
          <div class="mb-4">
            <div class="grid grid-cols-4 gap-4">
              <div class="content-center">Date</div>
              <div class="w-full col-span-3">
                <select
                  @change="emits('selectedSortByDate', true)"
                  v-model="modelSortByDate"
                  placeholder="Sort by Price"
                  class="block appearance-none w-full bg-white border p-2 px-3 leading-tight focus:outline-none focus:ring-2 focus:ring-indigo-500"
                >
                  <option class="" value="created_at">Earliest</option>
                  <option class="" value="-created_at">Recently</option>
                </select>
              </div>
            </div>
          </div>
          <div class="">
            <div class="grid grid-cols-4 gap-4">
              <div class="content-center">Price</div>
              <div class="w-full col-span-3">
                <select
                  @change="emits('selectedSortByPrice', true)"
                  v-model="modelSortByPrice"
                  placeholder="Sort by Price"
                  class="block appearance-none w-full bg-white border p-2 px-3 leading-tight focus:outline-none focus:ring-2 focus:ring-indigo-500"
                >
                  <option class="" value="price">Low to High</option>
                  <option class="" value="-price">High to Low</option>
                </select>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
