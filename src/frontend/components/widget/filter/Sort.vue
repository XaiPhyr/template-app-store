<script setup lang="ts">
  const emits = defineEmits([
    'selectedSortByName',
    'selectedSortByDate',
    'selectedSortByPrice',
  ]);

  const modelSortByName = defineModel('sortByName');
  const modelSortByPrice = defineModel('sortByPrice');
  const modelSortByDate = defineModel('sortByDate');
</script>

<template>
  <div class="">
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
                @change="emits('selectedSortByName')"
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
                @change="emits('selectedSortByDate')"
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
                @change="emits('selectedSortByPrice')"
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
</template>
