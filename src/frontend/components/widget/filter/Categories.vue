<script setup lang="ts">
  import { readCategories } from '~/composables/api/categories';

  const route = useRoute();

  const isShowingContent = ref(true);

  const categories = ref([]);
  const categoriesSize = ref(5);
  const categoriesTotal = ref(0);
  const categoriesSelected: any = ref([]);

  onMounted(() => {
    loadCategories();
  });

  const emits = defineEmits(['selectedCategory']);

  const loadCategories = async () => {
    const res = await readCategories({
      size: categoriesSize.value,
      sort: 'id',
    });

    if (res) {
      const { total, results } = res;
      categories.value = results;
      categoriesTotal.value = total;

      const { category }: any = route.query;
      categoriesSelected.value = category ? category.split(',') : [];
    }
  };

  const selectCategory = (item: string | number) => {
    const { category }: any = route.query;

    const items = category ? category.split(',') : [];

    if (!items.includes(item.toString())) {
      categoriesSelected.value.push(item);
    }

    if (items.includes(item.toString())) {
      const deselectCategory = categoriesSelected.value.filter(
        (x: string | number) => {
          return x !== item;
        }
      );

      categoriesSelected.value = deselectCategory;
    }

    emits('selectedCategory', categoriesSelected.value.join(','));
  };

  const isSelectedCategory = (item: string | any) => {
    if (categoriesSelected.value.includes(item)) {
      return 'font-bold';
    }

    return '';
  };

  const onSeeMoreCategories = () => {
    if (categoriesSize.value > categoriesTotal.value) {
      return;
    }

    categoriesSize.value += 10;

    loadCategories();
  };

  watch(
    () => route.query.category,
    (value) => {
      if (!value) {
        categoriesSelected.value = [];
      }
    }
  );
</script>

<template>
  <div class="">
    <div class="p-2 flex gap-4 justify-center">
      <div class="font-bold text-center">Categories</div>
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

      <div class="p-1 text-xs" @click="onSeeMoreCategories">
        <div class="px-5 py-1 hover:cursor-pointer active:bg-gray-100">
          <div class="p-1 flex justify-end text-blue-500">See more</div>
        </div>
      </div>
    </div>
  </div>
</template>
