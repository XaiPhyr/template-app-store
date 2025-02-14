<script setup lang="ts">
  const store = storeCart();

  const props = defineProps({
    item: {
      type: Object,
      default: () => null,
    },
    display: {
      type: String,
      default: 'grid',
    },
  });

  const popoverRef = ref();

  const item = computed(() => {
    return props.item;
  });

  const cardTitleRef: any = ref(null);

  const addToCart = (item: any) => {
    store.setCart(item);
  };

  const wishList = (item: any) => {
    console.log('WISHLIST: ', item);
  };

  const viewProduct = (event: any) => {
    popoverRef.value.show(event);
  };
</script>

<template>
  <div class="">
    <div class="border hover:border-slate-400" v-if="display === 'grid'">
      <div class="relative pt-1 px-1 hidden md:block h-64">
        <img
          class="object-fill h-full w-full"
          src="https://fastly.picsum.photos/id/13/2500/1667.jpg?hmac=SoX9UoHhN8HyklRA4A3vcCWJMVtiBXUg0W4ljWTor7s"
        />

        <div
          v-if="isBefore(item.created_at)"
          class="hover:cursor-pointer absolute top-1 right-1 bg-blue-500 text-white px-3 py-1 text-xs"
        >
          NEW
        </div>
      </div>
      <div class="p-5">
        <div
          ref="cardTitleRef"
          class="text-base font-semibold tracking-wide truncate hover:overflow-visible hover:whitespace-normal hover:text-ellipsis"
          :title="item.name"
        >
          {{ item.name }}
        </div>
        <div class="text-xs text-slate-500">{{ item.category }}</div>
        <div class="text-xs text-slate-500">
          Added: {{ readableDateTime(item.created_at) }}
        </div>
        <div class="text-sm mt-2 md:hidden block">{{ item.description }}</div>
        <div class="text-sm mt-2">
          {{ formatCurrencies(item.price) }}
        </div>

        <div class="flex justify-end mt-5 gap-1">
          <button
            @click="addToCart(item)"
            class="bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 py-1 rounded w-full text-white"
          >
            Add to Cart
          </button>

          <button
            title="wishlist"
            @click="wishList(item)"
            class="border border-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 hover:text-white rounded px-3 py-1 flex justify-center items-center"
          >
            <i class="pi pi-heart"></i>
          </button>

          <button
            title="view"
            @click="viewProduct($event)"
            class="border border-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 hover:text-white rounded px-3 py-1 flex justify-center items-center"
          >
            <i class="pi pi-eye"></i>
          </button>
        </div>
      </div>
    </div>

    <div class="border hover:border-slate-400" v-if="display === 'list'">
      <div class="xl:grid xl:grid-cols-5 gap-4">
        <div class="py-1 px-1 hidden md:block h-full w-full">
          <img
            class="object-cover h-full"
            src="https://fastly.picsum.photos/id/13/2500/1667.jpg?hmac=SoX9UoHhN8HyklRA4A3vcCWJMVtiBXUg0W4ljWTor7s"
          />
        </div>

        <div class="p-5 col-span-4">
          <div
            ref="cardTitleRef"
            class="text-base font-semibold tracking-wide truncate hover:overflow-visible hover:whitespace-normal hover:text-ellipsis"
            :title="item.name"
          >
            {{ item.name }}
          </div>
          <div class="text-xs text-slate-500">{{ item.category }}</div>
          <div class="text-xs text-slate-500">
            Added: {{ formatCurrencies(item.price) }}
          </div>
          <div class="text-sm mt-2">{{ item.description }}</div>
          <div class="text-sm mt-2">
            {{ formatCurrencies(item.price) }}
          </div>

          <div class="flex justify-end mt-5 gap-1">
            <button
              @click="addToCart(item)"
              class="bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 py-1 rounded w-full text-white"
            >
              Add to Cart
            </button>
            <button
              title="wishlist"
              @click="wishList(item)"
              class="border border-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 hover:text-white rounded px-3 py-1 flex justify-center items-center"
            >
              <i class="pi pi-heart"></i>
            </button>
            <button
              title="view"
              @click="viewProduct($event)"
              class="border border-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 active:scale-105 hover:text-white rounded px-3 py-1 flex justify-center items-center"
            >
              <i class="pi pi-eye"></i>
            </button>
          </div>
        </div>
      </div>
    </div>

    <Popover ref="popoverRef">
      <div class="w-64">
        {{ item.description }}
      </div>
    </Popover>
  </div>
</template>
