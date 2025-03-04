<script setup lang="ts">
  import { extendedNavs } from '~/interfaces/navigations';

  const currency: any = ref('USD');

  const navigations = extendedNavs;

  onMounted(() => {
    currency.value = localStorage.getItem('currency');
    setCurrency(currency.value);
  });

  const selectedCurrency = (event: any) => {
    setCurrency(event.target.value);
  };

  const isOpen = ref(false);

  const toggleDropdown = () => {
    isOpen.value = !isOpen.value;
  };
</script>

<template>
  <div class="">
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
        <div class="content-center">Contact Us: +123-456-7890</div>
      </div>

      <div class="flex gap-8">
        <div class="content-center">
          Unlock exclusive benefits and savings as a first-time customer with
          <span class="font-bold hover:text-indigo-300 hover:cursor-pointer">
            FIRSTCOMER
          </span>
          – your journey starts here!
        </div>
      </div>

      <div class="flex gap-8">
        <div
          class="content-center hover:cursor-pointer hover:text-blue-500 active:scale-105"
        >
          Privacy Policy
        </div>
        <div
          class="content-center hover:cursor-pointer hover:text-blue-500 active:scale-105"
        >
          Terms & Conditions
        </div>
        <div
          class="content-center hover:cursor-pointer hover:text-blue-500 active:scale-105"
        >
          About Us
        </div>
      </div>
    </div>

    <div
      class="w-full bg-indigo-500 text-white px-96 py-8 flex gap-4 justify-between"
    >
      <div class="">LOGO</div>
      <div class="">
        <ul class="flex gap-16">
          <li
            v-for="({ title, icon, link }, index) in navigations"
            :key="index"
            @click="link"
            class="flex hover:cursor-pointer hover:text-blue-950 active:scale-105"
          >
            <div class="flex gap-1">
              <span v-if="icon">
                <i :class="icon"></i>
              </span>
              <span v-if="title">
                {{ title }}
              </span>
            </div>
          </li>

          <transition name="fade">
            <div
              v-if="isOpen"
              class="fixed inset-0"
              @click="isOpen = false"
            ></div>
          </transition>

          <span class="relative inline-block text-left">
            <li
              @click="toggleDropdown"
              class="hover:cursor-pointer hover:text-blue-950 active:scale-105"
            >
              <i class="pi pi-user"></i>
            </li>

            <div
              v-show="isOpen"
              class="absolute right-0 mt-2 w-48 shadow-lg bg-white ring-1 ring-black ring-opacity-5 focus:outline-none z-50"
            >
              <div
                class="absolute right-0 top-[-8px] w-0 h-0 border-l-8 border-r-8 border-b-8 border-transparent border-b-white"
              ></div>

              <div class="">
                <ul class="first:pt-1 last:pb-1">
                  <li
                    class="text-slate-700 block px-4 py-2 text-sm hover:bg-indigo-500 hover:text-white hover:cursor-pointer"
                  >
                    Action I
                  </li>
                  <li
                    class="text-slate-700 block px-4 py-2 text-sm hover:bg-indigo-500 hover:text-white hover:cursor-pointer"
                  >
                    Action II
                  </li>
                </ul>
              </div>
            </div>
          </span>
        </ul>
      </div>
    </div>
  </div>
</template>
