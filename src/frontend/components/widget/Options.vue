<script setup lang="ts">
  const emits = defineEmits(['selectedRowSize', 'selectedDisplay']);

  const rowSize = ref(10);

  const isShowingContent = ref(true);

  const selectRowSize = () => {
    emits('selectedRowSize', rowSize.value);
  };

  const selectDisplay = (display: string) => {
    emits('selectedDisplay', display);
  };
</script>

<template>
  <div class="border w-80">
    <div class="p-2 flex gap-4 justify-center">
      <div class="font-bold text-center">Options</div>
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
        <div class="grid grid-cols-3 gap-4 md:mb-4">
          <div class="content-center">Rows:</div>
          <div class="w-full col-span-2">
            <select
              @change="selectRowSize"
              v-model="rowSize"
              class="block appearance-none w-full bg-white border py-2 px-3 leading-tight focus:outline-none focus:ring-2 focus:ring-indigo-500 text-right"
            >
              <option class="hover:bg-gray-200" :value="10">10</option>
              <option class="hover:bg-gray-200" :value="15">15</option>
              <option class="hover:bg-gray-200" :value="25">25</option>
              <option class="hover:bg-gray-200" :value="50">50</option>
              <option class="hover:bg-gray-200" :value="100">100</option>
            </select>
          </div>
        </div>

        <div class="hidden md:block">
          <div class="grid grid-cols-3 gap-4">
            <div class="content-center">Display:</div>
            <div class="w-full">
              <button
                title="Grid"
                @click="selectDisplay('grid')"
                class="w-full bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 text-white p-2 flex text-center justify-center"
              >
                <i class="pi pi-th-large"></i>
              </button>
            </div>
            <div class="w-full">
              <button
                title="List"
                @click="selectDisplay('list')"
                class="w-full bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 text-white p-2 flex text-center justify-center"
              >
                <i class="pi pi-list"></i>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
