import initApi from '~/api/instance';

const params: any = {
  url: '/categories',
  method: 'read',
  params: { size: 10 },
};

const readCategories = async ({ size }: any) => {
  try {
    params.params.size = size;
    const { data } = await initApi(params);

    return data;
  } catch (error) {
    console.log(error);
  }
};

export { readCategories };
