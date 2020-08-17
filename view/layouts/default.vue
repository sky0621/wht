<template>
  <v-app>
    <v-system-bar fixed app><v-icon>mdi-gmail</v-icon></v-system-bar>
    <v-app-bar fixed app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <v-window></v-window>
    </v-app-bar>
    <v-navigation-drawer v-model="drawer" fixed app>
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
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
      <v-container>
        <nuxt />
      </v-container>
    </v-main>
    <v-footer app>
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import { Component, Provide, Vue } from 'nuxt-property-decorator'

@Component({})
export default class DefaultLayout extends Vue {
  readonly title: string = 'What happened today?'

  @Provide()
  drawer: boolean = false

  items: object[] = [
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
  ]
}
</script>
