let currency = ref('USD');

const formatCurrencies = (payload: number) => {
  const obj: any = {
    style: 'currency',
    currency: currency.value,
  };

  const output = new Intl.NumberFormat('en', obj);

  if (!payload) {
    return output.format(0);
  }

  return output.format(payload);
};

const setCurrency = (item: string) => {
  currency.value = item;
  localStorage.setItem('currency', item);
};

export { formatCurrencies, setCurrency };
