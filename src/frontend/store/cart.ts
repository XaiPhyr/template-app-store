import { defineStore } from 'pinia';

const storeCart = defineStore('cart', () => {
  const stateCart: any = ref([]);
  const stateTotal: Ref<number> = ref(0);

  const getCart = computed(() => {
    return stateCart.value;
  });

  const getTotal = computed(() => {
    return stateTotal.value;
  });

  const setCart = (item: any) => {
    const existingItem = stateCart.value.find((x: any) => {
      return x.id === item.id;
    });

    if (!existingItem) {
      item.quantity = 1;
      stateCart.value.push(item);
      stateTotal.value += item.price * item.quantity;
      return;
    }

    const index = stateCart.value.indexOf(existingItem);

    stateCart.value[index].quantity += 1;

    stateTotal.value += item.price;
  };

  const setTotal = (item: any, method: string) => {
    if (method === 'plus') {
      stateTotal.value += item.price;
      return;
    }

    stateTotal.value -= item.price;
  };

  const removeItem = (item: any) => {
    stateTotal.value -= item.price * item.quantity;

    stateCart.value = stateCart.value.filter((x: any) => {
      return x !== item;
    });
  };

  return {
    stateCart,
    getCart,
    setCart,
    getTotal,
    setTotal,
    removeItem,
  };
});

export { storeCart };
