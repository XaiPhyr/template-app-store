<script setup lang="ts">
  const router = useRouter();
  const currency: any = ref('USD');

  onMounted(() => {
    currency.value = localStorage.getItem('currency');
    setCurrency(currency.value);
  });

  const goTo = (link: string) => {
    router.push(link);
  };

  const selectedCurrency = (event: any) => {
    setCurrency(event.target.value);
  };
</script>

<template>
  <div class="">
    <div
      class="w-full bg-indigo-950 text-white p-2 flex justify-between text-xs"
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
            @click="goTo('/')"
            class="hover:cursor-pointer hover:text-blue-950 active:scale-105"
          >
            Home
          </li>
          <li
            @click="goTo('/shop')"
            class="hover:cursor-pointer hover:text-blue-950 active:scale-105"
          >
            Shop
          </li>
          <li class="hover:cursor-pointer hover:text-blue-950 active:scale-105">
            Profile
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>
