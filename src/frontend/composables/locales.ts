const formatCurrencies = (payload: number) => {
  const obj: any = {
    style: 'currency',
    currency: 'USD',
  };

  const currency = new Intl.NumberFormat('en', obj);

  if (!payload) {
    return currency.format(0);
  }

  return currency.format(payload);
};

export { formatCurrencies };
