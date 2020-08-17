<template>
  <v-layout column justify-center align-center>
    <v-flex>
      <div>Wht-component</div>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { defineComponent } from '@vue/composition-api'
import { useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { watch } from '@vue/composition-api'

export default defineComponent({
  setup() {
    const { result } = useQuery(gql`
      query FindWht {
        findWht {
          id
          recordDate
          title
          textContents {
            id
            name
            text
          }
          imageContents {
            id
            name
            path
          }
        }
      }
    `)
    watch(result, (value) => {
      console.log(result)
      console.log(value)
    })
  },
})
</script>
