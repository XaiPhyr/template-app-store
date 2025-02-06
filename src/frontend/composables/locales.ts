const formatCurrencies = (payload: number) => {
  if (!payload) {
    return '';
  }

  const obj: any = {
    style: 'currency',
    currency: 'USD',
  };

  const currency = new Intl.NumberFormat('en', obj);

  return currency.format(payload);
};

export { formatCurrencies };
