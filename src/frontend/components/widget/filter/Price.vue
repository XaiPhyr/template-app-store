<script setup lang="ts">
  const emits = defineEmits(['selectedPriceRange']);

  const props = defineProps({
    priceRange: {
      type: Array,
      default: [0, 0],
    },
  });

  const isShowingContent = ref(true);

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
</script>

<template>
  <div class="">
    <div class="p-2 flex gap-4 justify-center">
      <div class="font-bold text-center">Price Range</div>
      <i
        v-if="!isShowingContent"
        @click="isShowingContent = !isShowingContent"
        class="pi pi-angle-down content-center hover:cursor-pointer active:scale-110"
      />
      <i
        v-else
        @click="isShowingContent = !isShowingContent"
        class="pi pi-angle-up content-center hover:cursor-pointer active:scale-110"
      />
    </div>

    <div :class="`border-t ${isShowingContent ? '' : 'hidden'}`">
      <div class="pt-5 px-5">
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
        <div class="p-1 flex flex-col">
          <div class="">Amount:</div>
          <div class="text-sm text-slate-500">
            {{ displayPriceRange }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
