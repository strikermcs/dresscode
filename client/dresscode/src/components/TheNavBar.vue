<template>
    <v-app-bar theme="dark" color="#323B50" flat>
        <ProductTabs :tabs="links" />
        <v-btn size="small" color="#F34F64" variant="contained" class="btn">
            <v-icon icon="mdi-chevron-down"></v-icon>
        </v-btn>
        <v-spacer></v-spacer>
        <v-divider vertical inset class="mr-4" color="#778290"></v-divider>
        <v-text-field label="Search" 
        prepend-inner-icon="mdi-magnify" 
        bg-color="#323B50"
        color="#778290"
        variant="underlined" 
        ></v-text-field>
        <v-btn class="user">
            Hi, strikermcs
            <v-icon icon="mdi-chevron-down"></v-icon>
        </v-btn>
    </v-app-bar>
</template>


<script>
import ProductTabs from './products/ProductTabs.vue'

export default {
    data: () => ({
        links: [
            {value: 'Dashboard', name: 'Dashboard'},
            {value: 'Messages', name: 'Messages'},
            {value: 'Profile', name: 'Profile'},
            {value: 'Updates', name: 'Updates'}  
        ],
        items : [
            {icon: 'mdi-laptop'},
            {icon: 'mdi-tag-outline'},
            {icon: 'mdi-shopping-outline'},
            {icon: 'mdi-wallet-travel'},
            {icon: 'mdi-apps'},
            {icon: 'mdi-text-box-multiple-outline'},
        ]
    }),

    components: {
        ProductTabs
    }
}
</script>

<style>

.user {
   text-transform: none !important;
   color: #778290 !important;
}


.v-btn:not(.v-btn--round).v-size--small {
    min-width: 10px !important;
    padding: 0 5px !important;
}

.v-input__control {
    color: #778290;
}

.v-input__details{
    display: none !important;
}

.btn {
    color: white !important;
}
</style>