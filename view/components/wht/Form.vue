<template>
  <v-container>
    <v-form>
      <v-row justify="center">
        <v-col md="auto">
          RecordDate: <v-text-field v-model="input.recordDate"></v-text-field>
          <v-file-input
            v-model="input.image"
            accept="image/*"
            show-size
            label="「今日こと」画像ファイルをアップロードしてください"
            prepend-icon="mdi-image"
          ></v-file-input>
          <v-btn @click="save">登録</v-btn>
          <v-btn @click="cancel">キャンセル</v-btn>
        </v-col>
      </v-row>
    </v-form>
  </v-container>
</template>

<script lang="ts">
import { Vue, Component, Emit, Prop } from 'nuxt-property-decorator'
import { WhtInput } from '~/types/gql-types'

class WhtInputImpl implements WhtInput {
  /** 記録日 */
  recordDate: any
  /** 画像 */
  image: any

  constructor(recordDate: any) {
    this.recordDate = recordDate
  }
}

@Component({})
export default class WhtForm extends Vue {
  @Prop({ default: () => {} })
  readonly recordDate!: string

  // 入力フォームの初期化
  input: WhtInput = new WhtInputImpl(this.recordDate)

  save() {
    // TODO: バリデーション実装
    this.$emit('submit', this.input)
  }

  @Emit('cancel')
  cancel(): void {}
}
</script>
