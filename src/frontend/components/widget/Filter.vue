<script setup lang="ts">
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
    isCategoryReset: {
      type: Boolean,
      default: false,
    },
  });

  const modelSortByName = defineModel('sortByName');
  const modelSortByPrice = defineModel('sortByPrice');
  const modelSortByDate = defineModel('sortByDate');

  const modelDateFrom = defineModel('dateFrom');
  const modelDateTo = defineModel('dateTo');

  const emits = defineEmits([
    'selectedCategory',
    'selectedPriceRange',
    'selectedSortByName',
    'selectedSortByDate',
    'selectedSortByPrice',
    'seeMoreCategories',
  ]);
</script>

<template>
  <div class="border w-80">
    <WidgetFilterCategories
      v-if="props.filter === 'categories'"
      @selected-category="(v) => emits('selectedCategory', v)"
    />

    <WidgetFilterPrice
      v-if="props.filter === 'price'"
      :price-range="priceRange"
      @selected-price-range="(v) => emits('selectedPriceRange', v)"
    />

    <WidgetFilterSort
      v-if="props.filter === 'sort'"
      v-model:sort-by-name="modelSortByName"
      v-model:sort-by-price="modelSortByPrice"
      v-model:sort-by-date="modelSortByDate"
      @selected-sort-by-name="emits('selectedSortByName')"
      @selected-sort-by-price="emits('selectedSortByPrice')"
      @selected-sort-by-date="emits('selectedSortByDate')"
    />

    <WidgetFilterDate
      v-if="props.filter === 'date'"
      v-model:date-from="modelDateFrom"
      v-model:date-to="modelDateTo"
    />
  </div>
</template>
