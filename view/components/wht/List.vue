<template>
  <v-container>
    <v-row justify="center">
      <v-col md="1">
        <v-btn fab @click="move">
          <v-icon>mdi-plus</v-icon>
        </v-btn>
      </v-col>
      <v-col md="auto">
        <v-data-table :items="whts" :headers="headers">
          <template v-slot:item.path="{ item }">
            <v-img :src="item.path" eager max-width="128px"></v-img>
          </template>
        </v-data-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Vue, Component, Prop, Emit } from 'nuxt-property-decorator'
import { DataTableHeader } from '~/types/vuetify'
import 'vue-apollo'
import { Wht } from '~/types/gql-types'

@Component({})
export default class WhtList extends Vue {
  @Prop({ default: () => {} })
  readonly whts!: Wht[]

  get headers(): DataTableHeader[] {
    return [
      { text: 'ID', value: 'id' },
      { text: '日づけ', value: 'recordDate' },
      { text: '画像', value: 'path' },
    ]
  }

  @Emit('move')
  move(): void {}
}
</script>
