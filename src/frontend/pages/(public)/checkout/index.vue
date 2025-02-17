<script setup lang="ts">
  import type { CartsInterface } from '~/interfaces/cart';

  useHead({
    title: 'Checkout',
  });

  definePageMeta({
    middleware: [
      (to, from) => {
        if (!to.query.cart) {
          const router = useRouter();
          router.push('/shop');
        }
      },
    ],
  });

  const route: any = useRoute();
  const router = useRouter();
  const items: Ref<CartsInterface> = ref([]);
  const total = ref(0);
  const subtotal = ref(0);
  const voucher = ref('');
  const hasPromo = ref(false);
  const promoAmount = ref(0);
  const shippingMethod = ref('standard');
  const standardAmount = ref(29.99);
  const expressAmount = ref(79.99);

  if (route.query.cart) {
    const decodeB64 = atob(route.query.cart || '');

    items.value = JSON.parse(decodeB64 || '');

    subtotal.value = items.value.reduce((a: any, b: any) => {
      return a + b.total;
    }, 0);

    total.value = subtotal.value;
  }

  const goBack = () => {
    router.push(`/shop?cart=${route.query.cart}`);
  };

  const proceed = () => {
    localStorage.removeItem('cart');
  };

  const computedTotal = () => {
    if (hasPromo.value) {
      const discounted = (subtotal.value * promoAmount.value) / 100;

      if (shippingMethod.value === 'express') {
        const withShipping = subtotal.value + expressAmount.value;
        return (total.value = withShipping - discounted + expressAmount.value);
      }

      return (total.value = subtotal.value - discounted + standardAmount.value);
    }

    if (shippingMethod.value === 'express') {
      return (total.value = subtotal.value + expressAmount.value);
    }

    return (total.value = subtotal.value + standardAmount.value);
  };

  const onApplyVoucher = () => {
    if (voucher.value.toLocaleLowerCase() === 'firstcomer') {
      hasPromo.value = true;
      promoAmount.value = 25;
      return;
    }

    hasPromo.value = false;
    promoAmount.value = 0;
  };

  const getShippingAmount = computed({
    get() {
      if (shippingMethod.value === 'standard') {
        return formatCurrencies(standardAmount.value);
      }

      return formatCurrencies(expressAmount.value);
    },
    set(v) {
      //
    },
  });
</script>

<template>
  <div class="p-4 xl:px-64 pb-[5rem]">
    <div
      @click="goBack"
      class="flex mb-4 text-blue-500 hover:cursor-pointer active:text-blue-600"
    >
      <div class="content-center mr-2">
        <i class="pi pi-chevron-left text-xs"></i>
      </div>
      <div class="content-center">Back</div>
    </div>
    <div class="flex md:flex-row flex-col gap-4">
      <div class="w-full md:w-4/6">
        <div class="border h-full w-auto p-4 md:p-32">
          <div class="mb-4">
            <div class="font-bold text-2xl mb-4">
              <i class="pi pi-credit-card text-xl"></i>
              Checkout
            </div>
            <div class="text-sm text-slate-500 mb-4">
              <span class="font-bold">Notice: </span>We value your privacy and
              do not store any payment or personal details after your
              transaction is completed. All payment processing is securely
              handled by our trusted third-party providers, ensuring that your
              sensitive information, including credit card details, is not saved
              on our servers. We implement strict security measures to protect
              your data during the checkout process. By proceeding with your
              purchase, you acknowledge that our company does not retain or
              access your payment information.
            </div>
            <div class="text-xl font-bold mb-4">Billing Details</div>
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
            <div class="flex md:flex-row flex-col gap-4 mb-4">
              <div class="flex flex-col w-full">
                <label for="id-address" class="font-bold">Address:</label>
                <input
                  id="id-address"
                  type="text"
                  class="p-2 border focus:outline-indigo-500"
                />
              </div>
            </div>
            <div class="flex md:flex-row flex-col gap-4 mb-4">
              <div class="flex flex-col w-full">
                <label for="id-country" class="font-bold">Country:</label>
                <input
                  id="id-country"
                  type="text"
                  class="p-2 border focus:outline-indigo-500"
                />
              </div>
              <div class="flex flex-col w-full">
                <label for="id-postal" class="font-bold">Postal:</label>
                <input
                  id="id-postal"
                  type="text"
                  class="p-2 border focus:outline-indigo-500"
                />
              </div>
            </div>
            <div class="flex md:flex-row flex-col gap-4 mb-4">
              <div class="flex flex-col w-full">
                <label for="id-email" class="font-bold">Email:</label>
                <input
                  id="id-email"
                  type="text"
                  class="p-2 border focus:outline-indigo-500"
                />
              </div>
              <div class="flex flex-col w-full">
                <label for="id-mobile" class="font-bold">Mobile:</label>
                <input
                  id="id-mobile"
                  type="text"
                  class="p-2 border focus:outline-indigo-500"
                />
              </div>
            </div>
            <div class="flex md:flex-row flex-col gap-4">
              <div class="">
                <label class="flex items-center space-x-2 mb-1">
                  <input
                    type="checkbox"
                    name="option"
                    value="newsletter"
                    class="w-4 h-4"
                  />
                  <span class="text-sm text-slate-500">
                    Sign Up for Weekly Newsletters
                  </span>
                </label>
                <div class="text-xs text-slate-500 italic">
                  By signing up for our weekly newsletter, you'll receive
                  exclusive updates, special offers, and the latest news
                  straight to your inbox.
                </div>
              </div>
            </div>
          </div>

          <div class="">
            <div class="text-xl font-bold mb-4">Shipping Method</div>

            <div class="mb-8">
              <div class="">
                <div class="flex justify-between">
                  <label class="flex items-center space-x-2">
                    <input
                      type="radio"
                      name="option"
                      value="standard"
                      class="w-4 h-4"
                      v-model="shippingMethod"
                    />
                    <span>Standard</span>
                  </label>
                  <div class="">
                    {{ formatCurrencies(standardAmount) }}
                  </div>
                </div>
              </div>
              <div class="text-sm text-slate-500">
                Standard shipping is a cost-effective delivery option that
                typically takes longer than express or expedited shipping. It
                usually follows a predetermined transit time, often ranging from
                3 to 7 business days, depending on the destination. This method
                is commonly used for non-urgent orders and may not include
                tracking or guaranteed delivery dates.
              </div>
            </div>
            <div class="">
              <div class="">
                <div class="flex justify-between">
                  <label class="flex items-center space-x-2">
                    <input
                      type="radio"
                      name="option"
                      value="express"
                      class="w-4 h-4"
                      v-model="shippingMethod"
                    />
                    <span>Express</span>
                  </label>
                  <div class="">
                    {{ formatCurrencies(expressAmount) }}
                  </div>
                </div>
              </div>
              <div class="text-sm text-slate-500">
                Express shipping is a faster delivery option that ensures orders
                arrive quicker, usually within 1 to 3 business days. It
                typically costs more than standard shipping but includes
                tracking and prioritized handling. This method is ideal for
                urgent orders that require timely delivery.
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="w-full md:w-2/6 h-full">
        <div class="border h-full w-auto p-4 mb-4">
          <div class="flex gap-4">
            <label for="id-voucher" class="font-bold content-center">
              Voucher:
            </label>
            <input
              id="id-voucher"
              type="text"
              class="p-2 border w-full focus:outline-indigo-500"
              v-model="voucher"
            />
            <button
              @click="onApplyVoucher"
              class="bg-green-500 hover:bg-green-600 py-2 px-4 text-white active:bg-green-700 active:scale-105"
            >
              Apply
            </button>
          </div>
        </div>
        <div class="border h-full w-auto p-4 content-center mb-4">
          <div class="">
            <div
              class="first:mt-4 last:mb-4"
              v-for="(item, index) in items"
              :key="index"
            >
              <div class="p-1">
                <div class="py-1">
                  <div class="grid grid-cols-4 gap-4">
                    <div class="">
                      <img
                        class="object-fill h-full w-full"
                        src="https://fastly.picsum.photos/id/13/2500/1667.jpg?hmac=SoX9UoHhN8HyklRA4A3vcCWJMVtiBXUg0W4ljWTor7s"
                      />
                    </div>
                    <div class="col-span-3 content-center text-md">
                      {{ item.name }}
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

          <div class="p-1">
            <div class="grid grid-cols-3">
              <div class="">Subtotal:</div>
              <div class="col-span-2 text-right">
                {{ formatCurrencies(subtotal) }}
              </div>
            </div>

            <div class="grid grid-cols-3" v-if="hasPromo && voucher !== ''">
              <div class="col-span-2">
                Promo <span class="font-bold">({{ voucher }})</span>:
              </div>
              <div class="text-right">{{ promoAmount }}%</div>
            </div>

            <hr class="my-2" />

            <div class="text-slate-500 text-sm">
              <div class="grid grid-cols-3">
                <div class="col-span-2">Shipping:</div>
                <div class="text-right">
                  {{ getShippingAmount }}
                </div>
              </div>
              <div class="text-xs">
                {{
                  shippingMethod === 'standard'
                    ? '(1 to 3 business days)'
                    : '(3 to 7 business days)'
                }}
              </div>
            </div>

            <hr class="my-2" />

            <div class="grid grid-cols-3 text-xl">
              <div class="">Total:</div>
              <div class="col-span-2 text-right font-bold">
                {{ formatCurrencies(computedTotal()) }}
              </div>
            </div>
          </div>
        </div>

        <div class="flex justify-end gap-4">
          <div class="w-full">
            <button
              @click="proceed"
              class="bg-green-500 hover:bg-green-600 py-3 w-full text-white active:bg-green-700 active:scale-105"
            >
              Proceed to payment
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
