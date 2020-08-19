<template>
  <v-container>
    <v-form>
      <v-row justify="center">
        <v-col md="auto">
          RecordDate: <v-text-field v-model="input.recordDate"></v-text-field>
          <v-file-input v-model="input.image"></v-file-input>
          <v-btn @click="save">登録</v-btn>
          <v-btn @click="cancel">キャンセル</v-btn>
        </v-col>
      </v-row>
    </v-form>
  </v-container>
</template>

<script lang="ts">
import { Vue, Component, Emit } from 'nuxt-property-decorator'
import 'vue-apollo'
import { WhtInput } from '~/types/gql-types'

class WhtInputImpl implements WhtInput {
  /** 記録日 */
  recordDate: any
  /** 画像 */
  image: any
}

@Component({})
export default class WhtForm extends Vue {
  // 入力フォームの初期化
  input: WhtInput = new WhtInputImpl()

  save() {
    // TODO: バリデーション実装
    this.$emit('submit', this.input)
  }

  @Emit('cancel')
  cancel(): void {}
}
</script>
