<template>
  <v-app :style="{background: $vuetify.theme.themes.dark.background}">
    <TheSideBar />
    <TheNavBar />
     <v-container fluid>
          <router-view />
     </v-container>
  </v-app>
</template>

<script>
import TheSideBar from './components/TheSidebar.vue'
import TheNavBar from './components/TheNavBar.vue'

export default {
  name: 'App',

  components: {
    TheSideBar,
    TheNavBar
  }
}
</script>

<style>
::-webkit-scrollbar {
  display: none;
}

* {
  -ms-overflow-style: none;  
  scrollbar-width: none;
}
</style>
