<template>
  <v-row class="fill-height">
    <v-col>
      <v-sheet height="64">
        <v-toolbar flat color="white">
          <v-btn outlined class="mr-4" color="grey darken-2" @click="setToday">
            Today
          </v-btn>
          <v-btn fab text small color="grey darken-2" @click="prev">
            <v-icon small>mdi-chevron-left</v-icon>
          </v-btn>
          <v-btn fab text small color="grey darken-2" @click="next">
            <v-icon small>mdi-chevron-right</v-icon>
          </v-btn>
          <v-toolbar-title v-if="$refs.calendar">
            {{ $refs.calendar.title }}
          </v-toolbar-title>
        </v-toolbar>
      </v-sheet>
      <v-sheet height="500">
        <v-calendar
          ref="calendar"
          v-model="focus"
          color="primary"
          event-color="accent"
          :events="events"
          :weekdays="weekdays"
          @click:date="editWht"
          @click:event="showEvent"
          @change="updateRange"
        ></v-calendar>
        <v-menu
          v-model="selectedOpen"
          :close-on-content-click="false"
          :activator="selectedElement"
          offset-x
        >
          <v-card color="grey lighten-4" min-width="200px" flat>
            <v-toolbar :color="selectedEvent.color" dark>
              <v-btn icon>
                <v-icon>mdi-pencil</v-icon>
              </v-btn>
              <v-toolbar-title v-html="selectedEvent.name"></v-toolbar-title>
              <v-spacer></v-spacer>
              <v-btn icon>
                <v-icon>mdi-heart</v-icon>
              </v-btn>
              <v-btn icon>
                <v-icon>mdi-dots-vertical</v-icon>
              </v-btn>
            </v-toolbar>
            <v-card-text>
              <span v-html="selectedEvent.details"></span>
            </v-card-text>
            <v-card-actions>
              <v-btn text color="secondary" @click="selectedOpen = false">
                Cancel
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-menu>
      </v-sheet>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { Vue, Component, Prop, Watch } from 'nuxt-property-decorator'
import { Wht } from '~/types/gql-types'

@Component({})
export default class WhtList extends Vue {
  $refs!: {
    calendar: any
  }

  @Prop({ default: () => {} })
  readonly whts!: Wht[]

  focus: string = ''
  selectedEvent = {}
  selectedElement = null
  selectedOpen = false
  events: Array<any> = []

  readonly weekdays = [1, 2, 3, 4, 5, 6, 0]

  mounted() {
    this.$refs.calendar.checkChange()
  }

  async editWht(ctx: any) {
    this.focus = ctx.date
    await this.$router.push({
      path: '/wht/new?date=',
      query: { date: ctx.date },
    })
  }

  getEventColor(event: any) {
    return event.color
  }

  setToday() {
    this.focus = ''
  }

  prev() {
    this.$refs.calendar.prev()
  }

  next() {
    this.$refs.calendar.next()
  }

  showEvent(ctx: any) {
    console.log('showEvent')
    const open = () => {
      this.selectedEvent = ctx.event
      this.selectedElement = ctx.nativeEvent.target
      setTimeout(() => (this.selectedOpen = true), 10)
    }

    if (this.selectedOpen) {
      this.selectedOpen = false
      setTimeout(open, 10)
    } else {
      open()
    }

    ctx.nativeEvent.stopPropagation()
  }

  @Watch('whts')
  updateRange() {
    if (!this.whts) {
      return
    }
    this.whts.forEach((wht) => {
      this.events.push({
        name: 'done',
        start: new Date(`${wht.recordDate}T00:00:00`),
      })
    })
  }
}
</script>
