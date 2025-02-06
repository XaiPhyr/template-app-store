interface ParamsInterface {
  q?: string;
  size?: number;
  page: number;
  sort?: string;
}

export interface ApiInterface {
  url: string;
  method: string;
  body?: any | null;
  params?: ParamsInterface | null;
  id?: string | null;
}
