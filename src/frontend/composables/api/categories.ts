import initApi from '~/api/instance';

const params: any = {
  url: '/categories',
  method: 'read',
  params: { size: 10 },
};

const readCategories = async ({ size, sort }: any) => {
  try {
    params.params.size = size;
    params.params.sort = sort;
    const { data } = await initApi(params);

    return data;
  } catch (error) {
    console.log(error);
  }
};

export { readCategories };
