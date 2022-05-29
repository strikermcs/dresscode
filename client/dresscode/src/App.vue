<template>
  <v-app :style="{background: $vuetify.theme.themes.dark.background}">
    <Home />
  </v-app>
</template>

<script>
import Home from './views/Home.vue'

export default {
  name: 'App',

  components: {
    Home
  }
}
</script>
