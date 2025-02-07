// https://nuxt.com/docs/api/configuration/nuxt-config
import indigoPreset from './presets/indigo';

const config = defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  app: {
    head: {
      titleTemplate: '%s - APP STORE',
    },
  },
  css: ['./assets/css/global.css'],
  modules: ['@primevue/nuxt-module', '@nuxtjs/tailwindcss'],
  primevue: {
    options: {
      theme: {
        preset: indigoPreset,
        options: {
          darkModeSelector: 'none',
        },
      },
    },
  },
  runtimeConfig: {
    apiKey: '',
    public: {
      apiUrl: process.env.API_URL,
      staticUrl: process.env.STATIC_URL,
    },
  },
});

console.log(JSON.stringify(config, null, 2));

export default config;
