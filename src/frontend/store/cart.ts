import { defineStore } from 'pinia';

const storeCart = defineStore('cart', () => {
  const stateCart: any = ref([]);
  const stateTotal: Ref<number> = ref(0);

  const getCart = computed(() => {
    return stateCart.value;
  });

  const getTotal = computed(() => {
    const total = stateCart.value.reduce((a: any, b: any) => {
      return a + b.total;
    }, 0);

    stateTotal.value = total;

    return stateTotal.value;
  });

  const setCart = (item: any) => {
    const existingItem = stateCart.value.find((x: any) => {
      return x.id === item.id;
    });

    if (!existingItem) {
      item.quantity = 1;
      item.total = item.price;
      stateCart.value.push(item);
      return;
    }

    const index = stateCart.value.indexOf(existingItem);

    stateCart.value[index].quantity += 1;

    item.total += item.price;
  };

  const setTotal = (item: any, method: string) => {
    if (method === 'plus') {
      item.total += item.price;
      return;
    }

    item.total -= item.price;
  };

  const getQuantity = (item: any) => {
    const existingItem = stateCart.value.find((x: any) => {
      return x.id === item.id;
    });

    if (!existingItem) {
      return '';
    }

    return existingItem.quantity;
  };

  const removeItem = (item: any) => {
    stateTotal.value -= item.total;

    stateCart.value = stateCart.value.filter((x: any) => {
      return x !== item;
    });

    if (stateTotal.value < 0) {
      stateTotal.value = 0.0;
    }
  };

  return {
    stateCart,
    getCart,
    setCart,
    stateTotal,
    getTotal,
    setTotal,
    getQuantity,
    removeItem,
  };
});

export { storeCart };
