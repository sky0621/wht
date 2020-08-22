export default {
  /*
   ** Nuxt rendering mode
   ** See https://nuxtjs.org/api/configuration-mode
   */
  mode: 'spa',
  /*
   ** Nuxt target
   ** See https://nuxtjs.org/api/configuration-target
   */
  target: 'static',
  /*
   ** Headers of the page
   ** See https://nuxtjs.org/api/configuration-head
   */
  head: {
    titleTemplate: '%s - ' + process.env.npm_package_name,
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || '',
      },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },
  /*
   ** Global CSS
   */
  css: [],
  /*
   ** Plugins to load before mounting the App
   ** https://nuxtjs.org/guide/plugins
   */
  plugins: [],
  /*
   ** Auto import components
   ** See https://nuxtjs.org/api/configuration-components
   */
  components: true,
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [
    '@nuxtjs/eslint-module',
    '@nuxt/typescript-build',
    '@nuxtjs/vuetify',
  ],
  /*
   ** Nuxt.js modules
   */
  modules: [
    'nuxt-i18n',
    '@nuxtjs/apollo',
    '@nuxtjs/proxy',
    '@nuxtjs/toast',
    '@nuxtjs/auth',
  ],
  i18n: {
    locales: [
      {
        code: 'ja',
        file: 'ja.json',
      },
    ],
    defaultLocale: 'ja',
    lazy: true,
    langDir: 'lang/',
  },
  apollo: {
    clientConfigs: {
      default: {
        httpEndpoint:
          process.env.WHT_API_ENDPOINT || 'http://localhost:8080/query',
      },
    },
    defaultOptions: {
      $query: {
        fetchPolicy: 'network-only',
      },
    },
    errorHandler: '~/plugins/apollo-error-handler.ts',
    // tokenName: 'auth._token.auth0',
    // authenticationType: '', // default の Bearer だと「Bearer: Bearer」というように重複が起きるため
  },
  /*
   ** vuetify module configuration
   ** https://github.com/nuxt-community/vuetify-module
   */
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    optionsPath: '~/settings/vuetify.ts',
  },
  // auth: {
  //   redirect: {
  //     login: '/login',
  //     logout: '/login',
  //     callback: '/callback',
  //     home: '/',
  //   },
  //   strategies: {
  //     auth0: {
  //       domain: process.env.AUTH0_DOMAIN,
  //       client_id: process.env.AUTH0_CLIENT_ID,
  //       audience: process.env.AUTH0_AUDIENCE,
  //     },
  //   },
  //   plugins: ['~/plugins/auth.ts'],
  // },
  // router: {
  //   middleware: ['auth'],
  // },
  toast: {
    position: 'top-center',
    duration: 5000,
  },
  /*
   ** Build configuration
   ** See https://nuxtjs.org/api/configuration-build/
   */
  build: {},
  generate: {
    dir: './frontendgen',
  },
}
