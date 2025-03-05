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
  const voucherAmount = ref(0);
  const appliedVoucher = ref('');
  const showDialog = ref(false);
  const customer = ref('');
  const hasNoCustomer = ref(false);
  const change = ref('');
  const insufficientCash = ref(false);

  const products: any = ref([]);
  const productsPage = ref(1);
  const productsTotalPage = ref(0);
  const productsSize = ref(10);
  const productsTotalSize = ref(0);

  const paymentMethod = ref('cash');
  const amount = ref('0.00');

  const selectedDiscount: any = ref(null);
  const discounts = ref([
    { text: 'Senior', value: 20 },
    { text: 'Student', value: 15 },
  ]);

  onMounted(() => {
    loadProducts();

    window.addEventListener('scroll', scrollHandler);
    window.addEventListener('keydown', keydownHandler);
  });

  onBeforeUnmount(() => {
    window.removeEventListener('scroll', scrollHandler);
    window.removeEventListener('keydown', keydownHandler);
  });

  const scrollHandler = () => {
    const { scrollY, innerHeight } = window;
    const { scrollHeight } = document.documentElement;

    const totalHeight = scrollY + innerHeight;

    if (totalHeight >= scrollHeight - 10) {
      if (!isLoading.value) {
        if (productsSize.value < productsTotalSize.value) {
          loadProducts();
        }
      }
    }
  };

  const keydownHandler = (event: any) => {
    if (event.key === 'Escape') {
      if (showDialog.value) {
        showDialog.value = false;
      }
    }
  };

  const loadProducts = async () => {
    if (isLoading.value) return;

    isLoading.value = true;

    const params: any = {
      page: productsPage.value,
      size: productsSize.value,
    };

    const data = await readProducts(params);

    if (data) {
      const { results, total } = data;
      if (!results) {
        productsPage.value = 1;
        productsTotalSize.value = 1;
        productsTotalPage.value = 1;
        isLoading.value = false;
        return;
      }

      if (results) {
        productsSize.value += 10;
        products.value = results;
        productsTotalSize.value = total;
      }
    }

    isLoading.value = false;
  };

  const onCheckout = (v: string) => {
    const parsedAmount = parseFloat(amount.value);
    const parsedTotal = parseFloat(v);

    if (customer.value === '') {
      hasNoCustomer.value = true;
      return;
    }

    if (parsedAmount >= parsedTotal) {
      showDialog.value = true;
      hasNoCustomer.value = false;
      insufficientCash.value = false;
      change.value = formatCurrencies(parsedAmount - parsedTotal);
      return;
    }

    insufficientCash.value = true;
    hasNoCustomer.value = false;
  };

  const onChangeAmount = (event: any) => {
    if (event.target.value) {
      const { value } = event.target;

      const sanitizedValue = value.replace(/,/g, '');
      const parsedValue = parseFloat(sanitizedValue);

      if (isNaN(parsedValue)) {
        amount.value = '0.00';
        return;
      }

      amount.value = formatDecimals(parseFloat(sanitizedValue));
      return;
    }

    amount.value = '0.00';
  };

  const onApplyVoucher = () => {
    appliedVoucher.value = voucher.value;
    selectedDiscount.value = null;
  };

  const onRemoveVouchers = (event: any) => {
    event.stopPropagation();

    voucher.value = '';
    appliedVoucher.value = '';
    selectedDiscount.value = null;
    voucherAmount.value = 0;
  };

  const hasProducts = computed(() => {
    return products.value.length > 0;
  });

  const hasVouhcers = computed(() => {
    return (
      (selectedDiscount.value && selectedDiscount.value.text) || voucher.value
    );
  });

  const customerNumber = computed(() => {
    return `C-${Math.ceil(Math.random() * 10000)}`;
  });

  watch(
    () => voucher.value,
    (value) => {
      if (value === '' && !selectedDiscount.value) {
        appliedVoucher.value = '';
        return;
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

  watch(
    () => showDialog.value,
    (value) => {
      if (value) {
        document.body.style.overflow = 'hidden';
        return;
      }

      document.body.style.overflow = '';
    }
  );

  watch(
    () => selectedDiscount.value,
    (value: any) => {
      if (value) {
        const { text, value: amount } = value;
        appliedVoucher.value = text;
        voucherAmount.value = amount;
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
              class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-3 2xl:grid-cols-4 gap-4 last:pb-16 pb-4"
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

            <div
              class="w-full bg-gray-200 h-2 rounded-full overflow-hidden"
              v-if="isLoading"
            >
              <div class="w-full h-full bg-blue-500 animate-progress"></div>
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
                        :class="hasNoCustomer ? 'border-red-500' : ''"
                        v-model="customer"
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
                        class="absolute left-0 top-1/2 transform -translate-y-1/2 px-2"
                        :class="
                          insufficientCash ? 'text-red-500' : 'text-slate-500'
                        "
                      >
                        <i class="pi pi-money-bill"></i>
                      </span>
                      <input
                        id="id-cash-amount"
                        type="text"
                        class="p-2 border w-full focus:outline-indigo-500 text-right pl-8"
                        :class="
                          insufficientCash ? 'border-red-500 text-red-500' : ''
                        "
                        :onchange="onChangeAmount"
                        v-model="amount"
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
              <div class="border-t px-5 py-4">
                <div class="flex gap-4 mb-4">
                  <input
                    id="id-voucher"
                    type="text"
                    class="p-2 border w-full focus:outline-indigo-500"
                    v-model="voucher"
                  />
                  <button
                    @click="onApplyVoucher"
                    class="bg-indigo-500 hover:bg-indigo-600 py-2 px-4 text-white active:bg-indigo-700 active:scale-105"
                  >
                    Apply
                  </button>
                </div>

                <div
                  class="mb-4 last:mb-0"
                  v-for="({ text, value }, index) in discounts"
                  :key="index"
                >
                  <label class="flex items-center space-x-2 border p-4">
                    <input
                      type="radio"
                      name="option"
                      :value="{ text, value }"
                      class="w-4 h-4 accent-indigo-500"
                      v-model="selectedDiscount"
                    />
                    <span class="flex gap-2 w-full justify-between">
                      <div class="">{{ text }}</div>
                      <div class="text-sm content-center text-red-500">
                        ({{ value }}%)
                      </div>
                    </span>
                  </label>
                </div>

                <div class="" v-if="hasVouhcers">
                  <button
                    @click="onRemoveVouchers"
                    class="w-full bg-indigo-500 hover:bg-indigo-600 py-2 px-4 text-white active:bg-indigo-700 active:scale-105"
                  >
                    Remove
                  </button>
                </div>
              </div>
            </div>

            <SharedCart
              class="w-80"
              hide-copy
              pos-checkout
              :voucher="appliedVoucher"
              :voucher-amount="voucherAmount"
              @checkout="onCheckout"
            />
          </div>
        </div>
      </div>
    </div>

    <div class="h-screen" v-else></div>

    <div
      v-if="showDialog"
      id="id-dialog"
      class="fixed inset-0 bg-black bg-opacity-10 backdrop-blur-sm flex items-center justify-center"
      @click="showDialog = false"
    >
      <div class="bg-white p-4 w-96 shadow-lg text-center" @click.stop>
        <div class="flex justify-end">
          <div
            @click="showDialog = false"
            class="font-bold text-slate-500 hover:cursor-pointer"
          >
            <i class="pi pi-times"></i>
          </div>
        </div>

        <div class="p-4">
          <div class="mb-4">
            Customer
            <div class="font-bold text-2xl">{{ customer }}</div>
          </div>
          <div class="mb-4">
            Change
            <div class="font-bold text-2xl">
              {{ change }}
            </div>
          </div>
          <div class="mb-4">
            Number
            <div class="font-bold text-2xl" v-if="showDialog">
              {{ customerNumber }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
  @keyframes progress {
    0% {
      transform: translateX(-100%);
    }
    100% {
      transform: translateX(100%);
    }
  }

  .animate-progress {
    animation: progress 2s infinite linear;
  }
</style>
