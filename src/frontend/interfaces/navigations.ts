interface NavigationInterface {
  title: string;
  icon?: string;
  link?(): void;
}

interface NavigationsInterface extends Array<NavigationInterface> {}

export const defaultNavs: NavigationsInterface = [
  {
    title: 'Home',
    link: () => {
      const router = useRouter();
      router.push('/');
    },
  },
  {
    title: 'Shop',
    link: () => {
      const router = useRouter();
      router.push('/shop');
    },
  },
  {
    title: 'POS',
    link: () => {
      const router = useRouter();
      router.push('/pos');
    },
  },
];

export const extendedNavs: NavigationsInterface = [
  ...defaultNavs,

  {
    title: '',
    icon: 'pi pi-heart',
  },
  {
    title: '',
    icon: 'pi pi-shopping-cart',
  },
];

export const posNavs: NavigationsInterface = [
  {
    title: 'Home',
    icon: 'pi pi-home',
    link: () => {
      const router = useRouter();
      router.push('/');
    },
  },
  {
    title: '',
    icon: 'pi pi-shopping-cart',
  },
];
