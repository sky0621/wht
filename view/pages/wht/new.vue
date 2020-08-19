<template>
  <WhtForm @submit="save" @cancel="moveToIndex" />
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import 'vue-apollo'
import WhtForm from '~/components/wht/Form.vue'
import { WhtInput } from '~/types/gql-types'
import { CreateWht } from '~/apollo/mutations/wht'
import '@nuxtjs/toast/index'

@Component({
  components: { WhtForm },
})
export default class MovieNewPage extends Vue {
  async save(input: WhtInput) {
    try {
      const res = await this.$apollo.mutate({
        mutation: CreateWht,
        variables: {
          wht: {
            recordDate: input.recordDate,
            image: input.image,
          },
        },
      })
      if (res) {
        this.$toast.success('「今日こと」の登録に成功しました。')
        await this.$router.push('/wht')
      } else {
        this.$toast.error('「今日こと」の登録に失敗しました。')
      }
    } catch (err) {
      this.$toast.error(err)
    }
  }

  moveToIndex(): void {
    this.$router.push('/wht')
  }
}
</script>
