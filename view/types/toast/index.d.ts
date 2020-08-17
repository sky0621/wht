// @ts-ignore
import { Toast } from '@nuxtjs/toast'

declare module 'vue/types/vue' {
  interface Vue {
    readonly toast: Toast
  }
}
