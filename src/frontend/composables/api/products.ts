import initApi from '~/api/instance';

const params: any = {
  url: '/products',
  method: 'read',
  params: { page: 1, size: 10 },
};

const readProducts = async ({ page, size, sort, q, qExt }: any) => {
  try {
    params.params.page = page;
    params.params.size = size;
    params.params.sort = sort;
    params.params.q = q;
    params.params.qExt = qExt;

    const { data } = await initApi(params);

    return data;
  } catch (error) {
    console.log(error);
  }
};

export { readProducts };
