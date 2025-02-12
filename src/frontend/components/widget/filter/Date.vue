<script setup lang="ts">
  const isShowingContent = ref(true);
  const dateToRef: any = ref(null);

  const modelDateFrom = defineModel('dateFrom');
  const modelDateTo = defineModel('dateTo');

  watch(
    () => modelDateTo.value,
    (value) => {
      const unixFrom = unixDateTime(modelDateFrom.value);
      const unixTo = unixDateTime(value);

      const el = dateToRef.value.classList;
      if (dateToRef.value) {
        if (unixFrom <= unixTo) {
          el.remove('border-red-500');
          el.remove('text-red-500');
          return;
        }

        el.add('border-red-500');
        el.add('text-red-500');
      }
    }
  );
</script>

<template>
  <div class="">
    <div class="p-2 flex gap-4 justify-center">
      <div class="font-bold text-center">Date Range</div>
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
      <div class="p-5">
        <div class="mb-4">
          <div class="grid grid-cols-4 gap-4">
            <div class="content-center">From:</div>
            <div class="w-full col-span-3">
              <input
                type="date"
                class="px-3 py-2 border w-full focus:outline-none focus:ring-2 focus:ring-indigo-500"
                v-model="modelDateFrom"
              />
            </div>
          </div>
        </div>

        <div class="grid grid-cols-4 gap-4">
          <div class="content-center">To:</div>
          <div class="w-full col-span-3">
            <input
              ref="dateToRef"
              type="date"
              class="px-3 py-2 border w-full focus:outline-none focus:ring-2 focus:ring-indigo-500"
              v-model="modelDateTo"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
