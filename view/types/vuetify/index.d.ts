export interface VCalendarEventInfo {
  date: Date
  time: string
  year: number
  month: number
  day: number
  hour: number
  minute: number
  weekday: number
  hasDay: boolean
  hasTime: boolean
  past: boolean
  present: boolean
  future: boolean
}

export interface VCalendarEvent {
  start: VCalendarEventInfo
  end: VCalendarEventInfo
}
