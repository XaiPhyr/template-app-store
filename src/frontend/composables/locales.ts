let currency = ref('USD');

const formatCurrencies = (payload: number) => {
  const obj: any = {
    style: 'currency',
    currency: currency.value,
  };

  const formatter = new Intl.NumberFormat('en', obj);

  if (!payload) {
    return formatter.format(0);
  }

  return formatter.format(payload);
};

const setCurrency = (item: string) => {
  currency.value = item;
  localStorage.setItem('currency', item);
};

const formatDecimals = (payload: number | string | any) => {
  const formatter = new Intl.NumberFormat('en-US', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  });

  if (!payload) {
    return formatter.format(0);
  }

  return formatter.format(payload);
};

export { formatCurrencies, setCurrency, formatDecimals };
