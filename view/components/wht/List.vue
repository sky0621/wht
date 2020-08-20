<template>
  <v-container>
    <v-row justify="center">
      <v-col md="12">
        <v-toolbar flat>
          <v-btn outlined class="mr-4" color="grey darken-2" @click="setToday">
            Today
          </v-btn>
          <v-btn text small color="grey darken-2" @click="prev">
            <v-icon small>mdi-chevron-left</v-icon>Prev
          </v-btn>
          <v-btn text small color="grey darken-2" @click="next">
            Next<v-icon small>mdi-chevron-right</v-icon>
          </v-btn>
          <v-toolbar-title>{{ title }}</v-toolbar-title>
        </v-toolbar>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col md="12">
        <v-calendar
          :weekdays="weekdays"
          :short-months="isShort"
          :short-weekdays="isShort"
        ></v-calendar>
      </v-col>
    </v-row>
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

  readonly weekdays = [1, 2, 3, 4, 5, 6, 0]
  readonly isShort = false

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
