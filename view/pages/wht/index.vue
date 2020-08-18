<template>
  <WhtList :whts="findWht" @move="moveToNew" />
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import 'vue-apollo'
import WhtList from '@/components/wht/List.vue'
import { FindWht } from '@/apollo/queries/wht'

@Component({
  components: { WhtList },
  apollo: {
    findWht: {
      prefetch: true,
      query: FindWht,
    },
  },
})
export default class WhtIndexPage extends Vue {
  moveToNew(): void {
    this.$router.push('/wht/new')
  }
}
</script>
