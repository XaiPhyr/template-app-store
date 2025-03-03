<script setup lang="ts">
  import { posNavs } from '~/interfaces/navigations';

  const currency: any = ref('USD');

  onMounted(() => {
    currency.value = localStorage.getItem('currency');
    setCurrency(currency.value);
  });

  const selectedCurrency = (event: any) => {
    setCurrency(event.target.value);
  };
</script>

<template>
  <div class="">
    <div class="flex">
      <div
        class="hidden xl:block bg-indigo-500 text-white p-4 flex flex-col justify-center items-center xl:w-[6.5rem] 2xl:w-[5.9rem]"
      >
        <div class="fixed top-0 left-0 h-full flex flex-col items-center px-8">
          <div class="mt-4">L</div>

          <ul class="flex flex-col justify-center h-full">
            <li
              v-for="({ icon, link }, index) in posNavs"
              :key="index"
              @click="link"
              class="mb-8 hover:cursor-pointer hover:text-blue-950 active:scale-105"
            >
              <span v-if="icon">
                <i :class="icon" class="text-2xl"></i>
              </span>
            </li>
            <li
              class="mb-8 hover:cursor-pointer hover:text-blue-950 active:scale-105"
            >
              <i class="pi pi-th-large text-2xl"></i>
            </li>
          </ul>

          <ul class="mt-auto mb-4">
            <li
              class="mb-4 hover:cursor-pointer hover:text-blue-950 active:scale-105"
            >
              <i class="pi pi-cog text-2xl"></i>
            </li>
          </ul>
        </div>
      </div>

      <div class="flex-auto">
        <div
          class="w-full bg-indigo-950 text-white py-2 px-4 flex justify-between text-xs"
        >
          <div class="flex gap-4">
            <div class="content-center">Currency:</div>
            <div class="content-center">
              <select
                @change="selectedCurrency($event)"
                name="name-currency"
                id="id-currency"
                v-model="currency"
                class="border px-3 py-1 text-black appearance-none focus:outline-none focus:ring-2 focus:ring-indigo-500"
              >
                <option value="USD">USD</option>
                <option value="PHP">PHP</option>
              </select>
            </div>
          </div>
        </div>

        <div class="p-8">
          <slot />
        </div>
      </div>
    </div>

    <div class="hidden md:block">
      <CoreScrollToTop />
    </div>
  </div>
</template>
