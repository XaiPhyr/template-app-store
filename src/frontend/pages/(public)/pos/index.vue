<script setup lang="ts">
  import { readProducts } from '~/composables/api/products';

  useHead({
    title: 'POS',
  });

  definePageMeta({
    layout: 'pos',
  });

  const display = ref('grid');
  const isLoading = ref(false);
  const voucher = ref('');
  const appliedVoucher = ref('');

  const products: any = ref([]);
  const productsPage = ref(1);
  const productsTotalPage = ref(0);
  const productsSize = ref(50);
  const productsTotalSize = ref(0);

  const paymentMethod = ref('');
  const amount = ref('0.00');

  onMounted(() => {
    loadProducts();
  });

  const loadProducts = async () => {
    products.value = [];
    isLoading.value = true;

    const params: any = {
      page: productsPage.value,
      size: productsSize.value,
    };

    const data = await readProducts(params);

    if (data) {
      const { results } = data;
      if (!results) {
        productsPage.value = 1;
        productsTotalSize.value = 1;
        productsTotalPage.value = 1;
        isLoading.value = false;
        return;
      }

      if (results) {
        products.value = results;
      }
    }

    isLoading.value = false;
  };

  const onCheckout = (v: string) => {};

  const hasProducts = computed(() => {
    return products.value.length > 0;
  });

  const onChangeAmount = (event: any) => {
    if (event.target.value) {
      const { value } = event.target;
      amount.value = formatDecimals(+value.replaceAll(',', ''));
      return;
    }

    amount.value = '0.00';
  };

  watch(
    () => voucher.value,
    (value) => {
      if (value === '') {
        appliedVoucher.value = '';
      }
    }
  );

  watch(
    () => paymentMethod.value,
    (value) => {
      if (value !== 'cash') {
        amount.value = '0.00';
      }
    }
  );
</script>

<template>
  <div class="">
    <div class="" v-if="hasProducts">
      <div class="text-3xl font-semibold">Main Content Area</div>

      <div class="mt-4">
        <div class="flex gap-4 justify-between">
          <div id="pos-products" class="">
            <div
              class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-3 2xl:grid-cols-4 gap-4 pb-16"
              v-if="hasProducts"
            >
              <div class="" v-for="(item, index) in products" :key="index">
                <SharedCard
                  :display="display"
                  :item="item"
                  hide-wishlist
                  hide-view
                />
              </div>
            </div>
          </div>
          <div id="pos-cart" class="hidden xl:block">
            <div class="border mb-4">
              <div class="p-2">
                <div class="font-bold text-center">Customer</div>
              </div>
              <div class="border-t">
                <div class="px-5 py-1">
                  <div class="flex md:flex-row flex-col gap-4 mb-4">
                    <div class="flex flex-col w-full">
                      <label for="id-name" class="font-bold">Name:</label>
                      <input
                        id="id-name"
                        type="text"
                        class="p-2 border focus:outline-indigo-500"
                      />
                    </div>
                  </div>

                  <div class="mb4">
                    <label class="flex items-center space-x-2 border p-4 mb-4">
                      <input
                        type="radio"
                        name="option"
                        value="cash"
                        class="w-4 h-4 accent-indigo-500"
                        v-model="paymentMethod"
                      />
                      <span class="flex gap-2">
                        <i class="pi pi-money-bill content-center"></i>
                        <div class="">Cash</div>
                      </span>
                    </label>
                  </div>

                  <div class="flex gap-4 mb-4" v-if="paymentMethod === 'cash'">
                    <label
                      for="id-cash-amount"
                      class="font-bold content-center"
                    >
                      Amount:
                    </label>
                    <div class="relative w-full">
                      <span
                        class="absolute left-0 top-1/2 transform -translate-y-1/2 px-2 text-slate-500"
                      >
                        <i class="pi pi-money-bill"></i>
                      </span>
                      <input
                        id="id-cash-amount"
                        type="text"
                        class="p-2 border w-full focus:outline-indigo-500 text-right pl-8"
                        :onchange="onChangeAmount"
                        :value="amount"
                      />
                    </div>
                  </div>

                  <div class="mb-4">
                    <label class="flex items-center space-x-2 border p-4 mb-4">
                      <input
                        type="radio"
                        name="option"
                        value="qr"
                        class="w-4 h-4 accent-indigo-500"
                        v-model="paymentMethod"
                      />
                      <span class="flex gap-2">
                        <i class="pi pi-qrcode content-center"></i>
                        <div class="">Scan QR</div>
                      </span>
                    </label>
                  </div>
                </div>
              </div>
            </div>

            <div class="border mb-4 w-80">
              <div class="p-2">
                <div class="font-bold text-center">Discount</div>
              </div>
              <div class="border-t flex gap-4 px-5 py-4">
                <input
                  id="id-voucher"
                  type="text"
                  class="p-2 border w-full focus:outline-indigo-500"
                  v-model="voucher"
                />
                <button
                  @click="appliedVoucher = voucher"
                  class="bg-indigo-500 hover:bg-indigo-600 py-2 px-4 text-white active:bg-indigo-700 active:scale-105"
                >
                  Apply
                </button>
              </div>
            </div>

            <SharedCart
              class="w-80"
              hide-copy
              pos-checkout
              :voucher="appliedVoucher"
              @checkout="onCheckout"
            />
          </div>
        </div>
      </div>
    </div>

    <div class="h-screen" v-else></div>
  </div>
</template>
