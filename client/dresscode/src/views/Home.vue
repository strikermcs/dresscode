<template>
    <div class="home_page">
        <TheSideBar />
        <TheNavBar />
        <v-container fluid>

        </v-container>
    </div>
</template>

<script>
import TheSideBar from '../components/TheSidebar.vue'
import TheNavBar from '../components/TheNavBar.vue'

export default {
    name: 'Home',
    components: {
        TheSideBar,
        TheNavBar
    }
}
</script>