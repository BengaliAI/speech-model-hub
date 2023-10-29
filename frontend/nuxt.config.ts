// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
  srcDir: "src",
  css: ['~/assets/css/main.css'],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  modules: [
    '@element-plus/nuxt',
    '@pinia/nuxt',
  ],
  devtools: {enable: true},
  pinia: {
    disableVuex: true,
  },
  ssr: false,
})