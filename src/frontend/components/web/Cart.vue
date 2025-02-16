<script setup lang="ts">
  const router = useRouter();
  const route: any = useRoute();

  const store = storeCart();
  const items: any = ref([]);
  const total: Ref<number> = ref(0);

  watch(
    () => store.stateCart,
    (value) => {
      if (value) {
        items.value = store.getCart;
        total.value = store.getTotal;
      }
    },
    {
      deep: true,
    }
  );

  onMounted(() => {
    items.value = store.getCart;
    total.value = store.getTotal;

    if (route.query.cart) {
      const decodeB64 = atob(route.query.cart || '');

      items.value = JSON.parse(decodeB64 || '');

      total.value = items.value.reduce((a: any, b: any) => {
        return a + b.total;
      }, 0);

      store.stateCart = items.value;
      store.stateTotal = total.value;
    }
  });

  const cartLength = computed(() => {
    return store.stateCart.length;
  });

  const onCheckout = () => {
    if (cartLength.value > 0) {
      const stringify = JSON.stringify(store.stateCart);
      const b64 = btoa(stringify);
      localStorage.setItem('cart', b64);

      router.push(`/checkout?cart=${b64}`);
    }
  };

  const quantityPlus = (item: any) => {
    item.quantity += 1;
    store.setTotal(item, 'plus');
  };

  const quantityMinus = (item: any) => {
    if (item.quantity > 1) {
      item.quantity -= 1;
      store.setTotal(item, 'minus');
    }
  };

  const quantityInput = (item: any) => {
    item.total = item.price * item.quantity;
  };
</script>

<template>
  <div class="">
    <div class="border mb-4">
      <div class="p-2">
        <div class="font-bold text-center">Total</div>
      </div>
      <div class="border-t">
        <div class="px-5 py-1">
          <div class="p-1">
            <div class="grid grid-cols-3">
              <div class="">Subtotal:</div>
              <div class="col-span-2 text-right">
                {{ formatCurrencies(total) }}
              </div>
            </div>

            <hr class="my-2" />

            <div class="grid grid-cols-3">
              <div class="text-xl">Total:</div>
              <div class="col-span-2 text-right font-bold text-xl">
                {{ formatCurrencies(total) }}
              </div>
            </div>

            <div class="my-5">
              <button
                @click="onCheckout"
                class="bg-green-500 hover:bg-green-600 py-3 w-full text-white active:bg-green-700 active:scale-105"
              >
                CHECKOUT
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="border">
      <div class="p-2">
        <div class="font-bold text-center">Cart</div>
      </div>

      <div class="border-t" v-if="items.length > 0">
        <div
          class="first:mt-4 last:mb-4"
          v-for="(item, index) in items"
          :key="index"
        >
          <div class="p-1">
            <div class="px-5 py-1">
              <div class="grid grid-cols-4 gap-1">
                <div class="col-span-2 content-center text-md">
                  {{ item.name }}
                </div>
                <div class="col-span-2 flex justify-end">
                  <div class="flex items-center border py-2 w-32">
                    <button
                      v-if="item.quantity === 1"
                      @click="store.removeItem(item)"
                      class="px-2 text-gray-500 hover:text-red-500 hover:scale-110"
                    >
                      <i class="pi pi-trash"></i>
                    </button>

                    <button
                      v-if="item.quantity > 1"
                      @click="quantityMinus(item)"
                      class="px-2 text-gray-500 hover:scale-110"
                    >
                      <i class="pi pi-minus"></i>
                    </button>

                    <input
                      @change="quantityInput(item)"
                      type="number"
                      v-model="item.quantity"
                      class="w-full appearance-none [&::-webkit-inner-spin-button]:appearance-none [&::-webkit-outer-spin-button]:appearance-none [&::-moz-appearance:textfield] text-center outline-none bg-transparent"
                    />

                    <button
                      @click="quantityPlus(item)"
                      class="px-2 text-gray-500 hover:scale-110"
                    >
                      <i class="pi pi-plus"></i>
                    </button>
                  </div>
                </div>
              </div>

              <hr class="my-2" />

              <div class="flex">
                <div class="text-left w-full">
                  <span class="font-bold">
                    {{ formatCurrencies(item.price) }}
                  </span>
                  <span class="text-sm"> x{{ item.quantity }} </span>
                </div>
                <div class="text-right font-bold w-full">
                  {{ formatCurrencies(item.price * item.quantity) }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
