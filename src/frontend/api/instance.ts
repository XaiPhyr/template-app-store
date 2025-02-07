import axios, { type AxiosInstance } from 'axios';
import type { ApiInterface } from '~/interfaces/apiInterface';

const instance = () => {
  const runtime = useRuntimeConfig();
  const baseURL = runtime.public.apiUrl;

  const config = {
    baseURL,
    timeout: 1000,
    headers: { Accept: 'application/json', 'Content-Type': 'application/json' },
  };

  return axios.create(config);
};

const initApi = (item: ApiInterface) => {
  const init = instance();
  let action: any = null;

  const { url, method } = item;

  createInteceptorRequest(init);
  createInteceptorResponse(init);

  switch (method) {
    case 'create':
      break;
    case 'read':
      action = init.get(url, { params: item?.params });
      break;
    case 'update':
      break;
    case 'delete':
      break;

    default:
      break;
  }

  return action;
};

const createInteceptorRequest = (init: AxiosInstance) => {
  init.interceptors.request.use(
    (res) => {
      // console.log('INTERCEPTOR REQUEST: ', res);
      return res;
    },
    (err) => {
      return Promise.reject(err);
    }
  );
};

const createInteceptorResponse = (init: AxiosInstance) => {
  init.interceptors.response.use(
    (res) => {
      // console.log('INTERCEPTOR RESPONSE: ', res);
      return res;
    },
    (err) => {
      return Promise.reject(err);
    }
  );
};

export default initApi;
