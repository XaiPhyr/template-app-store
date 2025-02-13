const formatCurrencies = (payload: number) => {
  if (!payload) {
    return '0.00';
  }

  const obj: any = {
    style: 'currency',
    currency: 'USD',
  };

  const currency = new Intl.NumberFormat('en', obj);

  return currency.format(payload);
};

export { formatCurrencies };
