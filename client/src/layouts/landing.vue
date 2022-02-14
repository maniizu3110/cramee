<template>
  <v-app>
    <v-app-bar flat height="80">
      <v-container class="py-0 px-0 px-sm-2 fill-height">
        <v-spacer></v-spacer>
      </v-container>
    </v-app-bar>

    <v-main>
      <nuxt />

      <v-footer color="transparent">
      </v-footer>
    </v-main>
  </v-app>
</template>

<script>
import config from '@/configs'

export default {
  data() {
    return {
      config,
      links: [{
        label: 'Overview',
        to: '#'
      }, {
        label: 'Features',
        to: '#'
      }, {
        label: 'Pricing',
        to: '#'
      }, {
        label: 'Documentation',
        to: '#'
      }, {
        label: 'News',
        to: '#'
      }, {
        label: 'FAQ',
        to: '#'
      }, {
        label: 'About us',
        to: '#'
      }, {
        label: 'Carrers',
        to: '#'
      }, {
        label: 'Press',
        to: '#'
      }]
    }
  }
}
</script>
