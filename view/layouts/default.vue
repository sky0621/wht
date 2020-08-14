<template>
  <v-app>
    <v-system-bar><v-icon>mdi-gmail</v-icon></v-system-bar>
    <v-app-bar fixed app>
      <v-app-bar-nav-icon
        @click.stop="attributes.drawer = !attributes.drawer"
      />
      <v-toolbar-title v-text="attributes.title" />
      <v-spacer />
      <v-window></v-window>
    </v-app-bar>
    <v-navigation-drawer v-model="attributes.drawer" fixed app>
      <v-list>
        <v-list-item
          v-for="(item, i) in sideMenu"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-main>
      <nuxt />
    </v-main>
    <v-footer app>
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import { defineComponent, reactive } from '@vue/composition-api'

export default defineComponent({
  setup() {
    const sideMenu = reactive([
      {
        icon: 'mdi-home',
        title: 'Top',
        to: '/',
      },
      {
        icon: 'mdi-book',
        title: 'Diary',
        to: '/wht',
      },
    ])
    const attributes = reactive({
      title: 'What happened today?',
      drawer: false,
    })
    return {
      sideMenu,
      attributes,
    }
  },
})
</script>
