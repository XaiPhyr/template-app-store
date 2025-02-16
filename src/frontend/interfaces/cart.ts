interface CartInterface {
  name: string;
  price: number;
  quantity: number;
}

export interface CartsInterface extends Array<CartInterface> {}
